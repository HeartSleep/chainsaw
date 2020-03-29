package baseline

import (
	"chainsaw/network"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

func detectFiles(Url *url.URL) {
	list := [...]string{"admin.php", "admin.asp", "admin.jsp", "admin.aspx", "admin/"}
	for _,v := range list {
		Url.Path = v
		resp := network.DoRequest(Url, network.ReqParam{})
		if resp == nil {
			continue
		}
		if resp.StatusCode == 200 && strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
			body, e := ioutil.ReadAll(resp.Body)
			if e != nil {
				panic(e)
			}
			str := string(body)
			if len(str) > 500 {str = str[:500]}
			fmt.Println("[*] Detected "+ Url.String())
			fmt.Println(str)
		}
		resp.Body.Close()
	}
}

func detectGeneralFiles(Url *url.URL) {
	list := [...]string{"README.md", ".htaccess", "robots.txt", }
	for _,v := range list {
		Url.Path = v
		resp := network.DoRequest(Url, network.ReqParam{})
		ct := resp.Header.Get("Content-Type")
		if resp.StatusCode == 200 && !strings.Contains(ct, "text/html"){
			body, e := ioutil.ReadAll(resp.Body)
			if e != nil {
				panic(e)
			}
			str := string(body)
			str = strings.Replace(str, "\n", " ", -1)
			if len(str) > 500 {
				str = str[:500]
			}
			fmt.Println("[*] Detected "+ Url.String())
			fmt.Println(str)
		}
		resp.Body.Close()
	}
}

func crossdomain(Url *url.URL) {
	Url.Path = "/crossdomain.xml"
	resp := network.DoRequest(Url, network.ReqParam{})
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, e := ioutil.ReadAll(resp.Body)
		if e != nil {
			panic(e)
		}
		if strings.Contains(string(body), "cross-domain-policy") && strings.Contains(string(body), "domain=\"*\"") {
			fmt.Println("[*] Detected " + Url.String())
			fmt.Println(string(body))
		}
	}
}

// TODO Should detect with data flow
//func directoryListing(Url *url.URL) bool {
//	resp := network.DoRequest(Url, network.ReqParam{})
//	defer resp.Body.Close()
//	if resp.StatusCode == 200 {
//		body, e := ioutil.ReadAll(resp.Body)
//		if e != nil {
//			panic(e)
//		}
//		if strings.Contains(string(body), "Directory listing for")  {
//			log.Println("Detected Directory List.", Url.String())
//			return true
//		}
//	}
//	return false
//}