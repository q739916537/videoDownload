package http_repo

import (
	"encoding/json"
	"fmt"
	"github.com/lbbniu/aliyun-m3u8-downloader/pkg/download"
	"github.com/lbbniu/aliyun-m3u8-downloader/pkg/tool"
	"go.uber.org/zap"
	"strings"
	"sync"
	"time"
	"videoDownload/api/model"
	"videoDownload/middleware"
	"videoDownload/pkg"
	"videoDownload/pkg/http"
)

type URLProcessor func(url string) string

type TargetList struct {
	rwMutex  *sync.RWMutex
	wg       *sync.WaitGroup
	result   []VidePageList
	dbResult []model.VideInfo
	ids      []int
	ch       chan string
}

var tarGetRes = &TargetList{
	rwMutex: &sync.RWMutex{},
	wg:      &sync.WaitGroup{},
	ch:      make(chan string, 1),
}

// 获取总页数
func GetAllInfoIds() (string, error) {
	url := pkg.BaseUrl + pkg.PageListSuffix + "1"
	method := "GET"
	body, err := internal_http.HTTPRequest(url, method, []byte{})
	if err != nil {
		middleware.DefaultLog().Error("GetAllInfoIds is fail", zap.Error(err))
		return "", err
	}
	videIndex := &VidePageList{}
	err = json.Unmarshal(body, videIndex)
	if err != nil {
		middleware.DefaultLog().Error("GetAllInfoIds Unmarshal is fail", zap.Error(err))
		return "", err
	}
	pagecount := videIndex.Pagecount
	urlChan := make(chan string, 1)

	tarGetRes.wg.Add(1)
	go func() {
		defer func() {
			tarGetRes.wg.Done()
			close(urlChan)
		}()
		for i := 0; i < pagecount; i++ {
			var tagetUrl = pkg.BaseUrl + pkg.PageListSuffix + fmt.Sprint(i+1)
			urlChan <- tagetUrl
		}
	}()

	for i := 0; i < pkg.MaxGetPageList; i++ {
		tarGetRes.wg.Add(1)
		go func() {
			ConsumeUrl(urlChan)
			defer tarGetRes.wg.Done()
		}()
	}

	tarGetRes.wg.Wait()
	fmt.Println("创建了：", len(tarGetRes.dbResult))
	return string(body), nil
}

type VidePageList struct {
	Code      int              `json:"code"`
	Msg       string           `json:"msg"`
	Page      string           `json:"page"`
	Pagecount int              `json:"pagecount"`
	Limit     string           `json:"limit"`
	Total     int              `json:"total"`
	List      []model.VideInfo `json:"list"`
}

func sendUrl(url string) {
	method := "GET"
	body, err := internal_http.HTTPRequest(url, method, []byte{})
	if err != nil {
		errorInfo := model.NewVideErrorInfo()
		// 记录失败的url
		deleteErr := errorInfo.DeleteByUrl(url)
		if deleteErr != nil {
			middleware.DefaultLog().Error(" url download is fail,record delete Url is fail", zap.Error(deleteErr))
			return
		}
		err2 := errorInfo.CreateOne(model.VideErrorInfo{
			VodUrl:     url,
			Method:     method,
			CreateTime: time.Now(),
			Error:      err.Error(),
		})
		if err2 != nil {
			middleware.DefaultLog().Error(" url download is fail,record create Url is fail", zap.Error(err2))
			return
		}
		return
	}
	videIndex := &VidePageList{}
	err = json.Unmarshal(body, videIndex)
	if err != nil {
		middleware.DefaultLog().Error("GetAllInfoIds Unmarshal is fail", zap.Error(err), zap.String(url, url))
		return
	}
	tarGetRes.rwMutex.RLock()
	defer tarGetRes.rwMutex.RUnlock()
	tarGetRes.result = append(tarGetRes.result, *videIndex)
	tarGetRes.dbResult = append(tarGetRes.dbResult, videIndex.List...)

	deleErr := model.NewVideInfo().DeleteById(videIndex.List)
	if deleErr != nil {
		middleware.DefaultLog().Error("Delete Video is fail err:", zap.Error(err), zap.String(url, url))
	}

	err = model.NewVideInfo().Create(videIndex.List)
	if err != nil {
		middleware.DefaultLog().Error("crate list is fail", zap.Error(err))
		return
	}

	list := videIndex.List
	var ids []int
	for _, s := range list {
		ids = append(ids, s.VodId)
	}
	tarGetRes.ids = append(tarGetRes.ids, ids...)

}
func ConsumeUrl(urlChan <-chan string) {
	for {
		select {
		case url, ok := <-urlChan:
			if !ok {
				return
			}
			fmt.Println("接收到URL：", url)
			sendUrl(url)
		default:
			time.Sleep(time.Second * 1)
		}

	}
}

// GetVideInfoByUrl 根据id或者url下载视频
func GetVideInfoByUrl(id int, url string) error {
	if id != 0 {
		_, err := model.NewVideInfo().GetId(id)
		if err != nil {
			middleware.DefaultLog().Error("GetVideInfoByUrl is fail,id is not exits", zap.Error(err))
			return err
		}
		url = pkg.DownloadInfoUrl + fmt.Sprintf("%d", id)
	}
	recordErr, err := SendUrlByByteRecordErr(url, pkg.DownloadUrlInfo)
	if err != nil {
		return err
	}
	var result UrlInfoResult
	err = json.Unmarshal(recordErr, &result)
	if err != nil {
		middleware.DefaultLog().Error("GetVideInfoByUrl is fail,err is ", zap.Error(err))
		return err
	}
	if len(result.List) > 0 {

		vlist := strings.Split(result.List[0].VodPlayURL, "#")
		fmt.Println(vlist)
		for _, playURL := range vlist {
			go func() {
				URL := strings.Split(playURL, "$")
				downloadVideo(URL[1], "./tmp/", 8)
			}()
		}
	}
	return nil
}
func GetVideNameLike(name string, id int) ([]*model.VideInfo, error) {
	results, err := model.NewVideInfo().GetVideInfoByNameOrId(name, id)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func downloadVideo(url, output string, chanSize int) {
	if url == "" {
		tool.PanicParameter("url")
	}
	if chanSize <= 0 {
		panic("parameter 'chanSize' must be greater than 0")
	}

	downloader, err := download.NewDownloader(download.WithUrl(url), download.WithOutput(output))
	if err != nil {
		panic(err)
	}
	if err := downloader.Start(chanSize); err != nil {
		panic(err)
	}
	fmt.Println("Done!")
}

// 发送请求并返回byte数据记录err
func SendUrlByByteRecordErr(url, lx string) ([]byte, error) {
	method := "GET"
	body, err := internal_http.HTTPRequest(url, method, []byte{})
	if err != nil {
		errorInfo := model.NewVideErrorInfo()
		// 记录失败的url
		deleteErr := errorInfo.DeleteByUrl(url)
		if deleteErr != nil {
			middleware.DefaultLog().Error(" url download is fail,record delete Url is fail", zap.Error(deleteErr))
			return nil, err
		}
		var err2 = errorInfo.CreateOne(model.VideErrorInfo{
			VodUrl:     url,
			Method:     method,
			CreateTime: time.Now(),
			Error:      err.Error(),
			Lx:         lx,
		})
		if err2 != nil {
			middleware.DefaultLog().Error(" url download is fail,record create Url is fail", zap.Error(err2))
			return nil, err
		}
		return nil, err
	}
	return body, err
}

type UrlInfoResult struct {
	Code      int                 `json:"code"`
	Msg       string              `json:"msg"`
	Page      int                 `json:"page"`
	Pagecount int                 `json:"pagecount"`
	Limit     string              `json:"limit"`
	Total     int                 `json:"total"`
	List      []model.VideUrlInfo `json:"list"`
}
