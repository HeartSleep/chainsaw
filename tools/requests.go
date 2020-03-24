package tools

import (
	"crypto/tls"
	"net/http"
	"net/url"
)

type ReqParam struct {
	UA string
	Timeout int
	ContentType string
	Method string
	Redirect bool
	Headers map[string]string
	Proxy *url.URL
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
	obj.Proxy, _ = url.Parse("http://127.0.0.1:8080")
}

func DoRequest(url *url.URL, param ReqParam) *http.Response {
	param.LoadDefault()
	req, _ := http.NewRequest(param.Method, url.String(), nil)
	req.Header.Set("User-Agent", param.UA)
	if param.Method == "POST" {
		req.Header.Set("Content-Type", param.ContentType)
	}
	for k, v := range param.Headers {
		req.Header.Set(k, v)
	}
	tr := &http.Transport{
		Proxy:           http.ProxyURL(param.Proxy),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	resp, err := (&http.Client{
		Transport: tr,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}).Do(req)
	if err != nil {
		panic(err)
	}
	return resp
}