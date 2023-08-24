package internal_http

import (
	"bytes"
	"crypto/tls"
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
			Timeout: 5 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
		retries: 0,
	}
}

func (r *RetryClient) Do(req *http.Request) (*http.Response, error) {
	var resp *http.Response
	var err error

	for r.retries < maxRetries {
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
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
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
