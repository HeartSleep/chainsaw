package tools

import (
	"net/http"
)

type ReqParam struct {
	UA string
	Timeout int
	ContentType string
	Method string
	Redirect bool
	Headers []string
}

func (obj *ReqParam) LoadDefault() {
	if obj.Method == "" {
		obj.Method = "GET"
	}
	if obj.UA == "" {
		obj.UA = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36"
	}
	if obj.Timeout == 0 {
		obj.Timeout = 5
	}
	if obj.Method == "POST" && obj.ContentType == "" {
		obj.ContentType = "application/x-www-form-urlencoded"
	}
}

func  DoRequest(url string, param ReqParam) *http.Response {
	param.LoadDefault()
	req, _ := http.NewRequest(param.Method, url, nil)
	req.Header.Set("User-Agent", param.UA)
	if param.Method == "POST" {
		req.Header.Set("Content-Type", param.ContentType)
	}
	resp, err := (&http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}).Do(req)
	if err != nil {
		panic(err)
	}
	return resp
}