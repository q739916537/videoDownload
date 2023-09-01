package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/lbbniu/aliyun-m3u8-downloader/pkg/download"
	"github.com/lbbniu/aliyun-m3u8-downloader/pkg/tool"
	"net/http"
	"strconv"
	"videoDownload/api/model"
	"videoDownload/api/repository/http_repo"
	"videoDownload/server"
)

const (
	path = "/home/jellyfin/%s/"
)

var (
	serial string
	ids    int
	url    string
)

func Init() {
	flag.StringVar(&serial, "s", "1", "第几季")
	flag.IntVar(&ids, "i", 0, "平台视频id")
	flag.StringVar(&url, "u", "https://kuaikan-api.com/api.php/provide/vod/from/kuaikan/", "视频地址")
}

func main() {

	server.InitMiddleWare()
	//http_repo.GetAllInfoIds()
	http.HandleFunc("/GetId", func(writer http.ResponseWriter, request *http.Request) {
		name := request.URL.Query().Get("name")
		if name == "" {
			writer.Write([]byte("name is null"))
			return
		}
		results, err := model.NewVideInfo().GetVodName(name)
		if err != nil {
			writer.Write([]byte(err.Error()))
			return
		}
		bytes, _ := json.Marshal(results)
		writer.Write(bytes)
		return
	})
	http.HandleFunc("/Download", func(writer http.ResponseWriter, request *http.Request) {
		id := request.URL.Query().Get("id")
		if id == "" {
			writer.Write([]byte("id is null"))
			return
		}
		num, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		err = http_repo.GetVideInfoByUrl(num, "")
		if err != nil {
			writer.Write([]byte(err.Error()))
			return
		}
		writer.Write([]byte("success"))
		return
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}

	//os.Setenv("https_proxy", "http://127.0.0.1:7890")
	//os.Setenv("http_proxy", "http://127.0.0.1:7890")
	//Init()
	//flag.Parse()
	//if ids == 0 {
	//	return
	//}
	//server.InitMiddleWare()
	//rr, err := getResource(url, fmt.Sprintf("%d", ids))
	//if err != nil {
	//	fmt.Printf("%+v\n", err)
	//}
	//
	//videHome := fmt.Sprintf(path, rr.List[0].VodName)
	//os.MkdirAll(videHome, os.ModePerm)
	//vlist := strings.Split(rr.List[0].VodPlayURL, "#")
	//fmt.Println(vlist)
	//db, err := db.NewDb()
	//if err != nil {
	//	panic(err)
	//}
	//
	//vr, err := db.GetVideoRecords(ids)
	//if err != nil && err != gorm.ErrRecordNotFound {
	//	panic(err)
	//}
	//
	//if err == gorm.ErrRecordNotFound {
	//	vr = &types.VideoRecords{
	//		VideoDetailList: types.VideoDetailList{
	//			VodID:     rr.List[0].VodID,
	//			VodName:   rr.List[0].VodName,
	//			VodStatus: rr.List[0].VodStatus,
	//			VodYear:   rr.List[0].VodYear,
	//			VodSerial: serial,
	//		},
	//		Downloaded: 0,
	//	}
	//
	//	err := db.InsertVideoRecords(vr)
	//	if err != nil {
	//		fmt.Printf("%+v\n", err)
	//		return
	//	}
	//
	//	vr, err = db.GetVideoRecords(ids)
	//	if err != nil && err != gorm.ErrRecordNotFound {
	//		panic(err)
	//	}
	//}
	//
	//videoSerial := len(vlist)
	//if vr.Downloaded == videoSerial {
	//	return
	//}
	//
	//setp := vr.Downloaded
	//
	//tpaht := fmt.Sprintf("./tmp/%s", rr.List[0].VodName)
	//os.RemoveAll(tpaht)
	//
	//for i := vr.Downloaded; i < videoSerial; i++ {
	//	series := strings.Split(vlist[i], "$")
	//	fmt.Printf("开始下载⏬%s - 第%s集\n", videHome, series[0])
	//	downloadVideo(series[1], tpaht, 8)
	//	ss, err := strconv.Atoi(series[0])
	//	if err != nil {
	//		os.Rename(fmt.Sprintf("%s/index.m3u8.mp4", tpaht), fmt.Sprintf("%s/%s.mp4", videHome, vr.VodName))
	//	} else {
	//		s := series[0]
	//		if ss < 10 {
	//			s = fmt.Sprintf("0%s", s)
	//		}
	//		os.Rename(fmt.Sprintf("%s/index.m3u8.mp4", tpaht), fmt.Sprintf("%s/%sS0%sE%s.mp4", videHome, vr.VodName, serial, s))
	//	}
	//	os.RemoveAll(fmt.Sprintf("%s/index.m3u8.mp4", tpaht))
	//	os.RemoveAll(fmt.Sprintf("%s/ts", tpaht))
	//	n, _ := rand.Int(rand.Reader, big.NewInt(10))
	//	setp++
	//	vr.Downloaded = setp
	//	err = db.UpdateVideoRecords(vr)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Printf("⏬%s - 第%s集,已经完成,进入休眠(%d min)\n", videHome, series[0], n.Int64())
	//	time.Sleep(time.Minute * time.Duration(n.Int64()))
	//}
	//
	//defer func() {
	//	vr.Downloaded = setp
	//	err := db.UpdateVideoRecords(vr)
	//	if err != nil {
	//		fmt.Printf("%+v\n", err)
	//	}
	//	os.RemoveAll("/data/tmp/ts")
	//}()
}

//
//func getResource(url string, ids string) (resourceResp *types.ResourceResp, err error) {
//	resp, err := http2.Get(fmt.Sprintf("%s?ids=%s&ac=detail", url, ids))
//	if err != nil || resp.StatusCode != 200 {
//		return resourceResp, errors.Wrapf(err, "url=%s,ids=%s", url, ids)
//	}
//	defer resp.Body.Close()
//	body, err := io.ReadAll(resp.Body)
//	if err != nil {
//		return resourceResp, errors.Wrapf(err, "url=%s,ids=%s", url, ids)
//	}
//	err = json.Unmarshal(body, &resourceResp)
//	if err != nil {
//		return resourceResp, errors.Wrapf(err, "url=%s,ids=%s", url, ids)
//	}
//	return resourceResp, err
//}

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
