package baseline

import (
	"chainsaw/tools"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func springActuator(u *string) bool {
	list := [...]string{"/autoconfig", "/beans", "/env", "/configprops", "/dump", "/health", "/info", "/mappings", "/metrics", "/shutdown", "/trace"}
	for _, l := range list {
		entry := *u + l
		resp := tools.DoRequest(entry, tools.ReqParam{})
		defer resp.Body.Close()
		if resp.StatusCode == 200 {
			log.Println("[*] Detected Spring Actuator information leak.", entry)
		}
	}
	return false
}

func druid(u *string) bool {
	entry := *u+"/druid/index.html"
	resp := tools.DoRequest(entry, tools.ReqParam{})
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, e := ioutil.ReadAll(resp.Body)
		if e != nil {
			panic(e)
		}
		if strings.Contains(string(body), "Druid Stat Index")  {
			log.Println("[*] Detected Druid unauthorized.", entry)
			return true
		}
	}
	return false
}

func laravelDebug(u *string) bool {
	resp := tools.DoRequest(*u, tools.ReqParam{Method: "POST"})
	defer resp.Body.Close()
	if resp.StatusCode == 405 {
		body, e := ioutil.ReadAll(resp.Body)
		if e != nil {
			fmt.Println(e)
			return false
		}
		if strings.Contains(string(body), "MethodNotAllowedHttpException") {
			log.Println("[*] Detected Laravel debug mode.", *u)
			return true
		}
	}
	return false
}