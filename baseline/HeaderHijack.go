package baseline

import (
	"chainsaw/network"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"time"
)

/**
	Include: Origin, Cache Poison, etc.
 */

func CorsCheck(Url *url.URL) bool{
	evilOrigin := Url.Scheme+"://"+Url.Host+".evil"+Url.Host
	resp := network.DoRequest(Url, network.ReqParam{Headers: map[string]string{"Origin": evilOrigin}})
	if resp.Header.Get("Access-Control-Allow-Origin") == evilOrigin {
		log.Println("[*] CORS wrong config."+Url.String()+"--"+evilOrigin)
		return true
	}
	Url.Path = "/api"
	resp = network.DoRequest(Url, network.ReqParam{Headers: map[string]string{"Origin": evilOrigin}})
	if resp.Header.Get("Access-Control-Allow-Origin") == evilOrigin {
		log.Println("[*] CORS wrong config."+Url.String()+"--"+evilOrigin)
		return true
	}
	return false
}

func genIpAddr() string {
	rand.Seed(time.Now().Unix())
	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
}

func Cache(Url *url.URL) bool {
	ip := genIpAddr()
	list := map[string]string {
		"X-Forwarded-For": ip,
		"X-Host": ip,
		"X-Forwarded-Host": ip,
		"Cache-Control": "no-store",
		"X-Forwarded-Scheme": "nohttps",
	}
	resp := network.DoRequest(Url, network.ReqParam{Headers: list})
	if resp == nil {
		return false
	}
	defer resp.Body.Close()

	return true
}