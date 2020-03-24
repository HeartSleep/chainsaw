package baseline

import (
	"chainsaw/tools"
	"log"
	"net/url"
)

func CorsCheck(Url *url.URL) bool{
	evilOrigin := "https://"+Url.Host+".evil"+Url.Host
	resp := tools.DoRequest(Url, tools.ReqParam{Headers: map[string]string{"Origin": evilOrigin}})
	if resp.Header.Get("Access-Control-Allow-Origin") == evilOrigin {
		log.Println("[*] CORS wrong config."+Url.String()+"--"+evilOrigin)
		return true
	}
	Url.Path = "/api"
	resp = tools.DoRequest(Url, tools.ReqParam{Headers: map[string]string{"Origin": evilOrigin}})
	if resp.Header.Get("Access-Control-Allow-Origin") == evilOrigin {
		log.Println("[*] CORS wrong config."+Url.String()+"--"+evilOrigin)
		return true
	}
	return false
}