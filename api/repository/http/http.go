package http

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
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
			fmt.Println("接收到URL：", url)
			if !ok {
				return
			}
			sendUrl(url)
		default:
			time.Sleep(time.Second * 1)
		}

	}
}
