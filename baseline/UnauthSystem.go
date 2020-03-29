package baseline

import (
	"chainsaw/network"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"strings"
)

func springActuator(Url *url.URL) bool {
	list := [...]string{"/autoconfig", "/env", "/dump", "/health", "/info", "/mappings", "/trace"}
	for _, l := range list {
		Url.Path = l
		resp := network.DoRequest(Url, network.ReqParam{})
		if resp == nil {
			return false
		}
		if resp.StatusCode == 200 {
			log.Println("[*] Detected Spring Actuator information leak.", Url.String())
		}
		resp.Body.Close()
	}
	return false
}

func druid(Url *url.URL) bool {
	Url.Path = "/druid/index.html"
	resp := network.DoRequest(Url, network.ReqParam{})
	if resp == nil {
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, e := ioutil.ReadAll(resp.Body)
		if e != nil {
			panic(e)
		}
		if strings.Contains(string(body), "Druid Stat Index")  {
			log.Println("[*] Detected Druid unauthorized.", Url.String())
			return true
		}
	}
	return false
}

func laravelDebug(Url *url.URL) bool {
	resp := network.DoRequest(Url, network.ReqParam{Method: "POST"})
	if resp == nil {
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode == 405 {
		body, e := ioutil.ReadAll(resp.Body)
		if e != nil {
			fmt.Println(e)
			return false
		}
		if strings.Contains(string(body), "MethodNotAllowedHttpException") {
			log.Println("[*] Detected Laravel debug mode.", Url.String())
			return true
		}
	}
	return false
}