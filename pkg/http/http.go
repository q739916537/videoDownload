package internal_http

import (
	"fmt"
	io "io"
	"net/http"
	"time"
)

var maxRetries = 3

type RetryClient struct {
	client  *http.Client
	retries int
}

func NewRetryClient() *RetryClient {
	return &RetryClient{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		retries: 0,
	}
}

func (r *RetryClient) Do(req *http.Request) (*http.Response, error) {
	var resp *http.Response
	var err error
	for r.retries < maxRetries {
		req.Header.Add("authority", "kuaikan-api.com")
		req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
		req.Header.Add("accept-language", "zh-CN,zh;q=0.9")
		req.Header.Add("cache-control", "max-age=0")
		req.Header.Add("sec-ch-ua", "\"Not/A)Brand\";v=\"99\", \"Google Chrome\";v=\"115\", \"Chromium\";v=\"115\"")
		req.Header.Add("sec-ch-ua-mobile", "?0")
		req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
		req.Header.Add("sec-fetch-dest", "document")
		req.Header.Add("sec-fetch-mode", "navigate")
		req.Header.Add("sec-fetch-site", "none")
		req.Header.Add("sec-fetch-user", "?1")
		req.Header.Add("upgrade-insecure-requests", "1")
		req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")
		resp, err = r.client.Do(req)
		if err == nil {
			break
		}
		r.retries++
		time.Sleep(time.Second * time.Duration(r.retries))
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func HTTPRequest(url, method string, body []byte) ([]byte, error) {
	client := NewRetryClient()
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func main() {
	url := "https://api.example.com/data"
	method := "GET"
	body := []byte{}

	response, err := HTTPRequest(url, method, body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response:", response)
}
