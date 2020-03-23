package baseline

import (
	"chainsaw/tools"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func detectFiles(u *string) {
	list := [...]string{"admin.php", "admin.asp", "admin.jsp", "admin.aspx", "admin/"}
	for _,v := range list {
		entry := *u+ "/" +v
		resp := tools.DoRequest(entry, tools.ReqParam{})
		if resp.StatusCode == 200 && strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
			body, e := ioutil.ReadAll(resp.Body)
			if e != nil {
				panic(e)
			}
			str := string(body)
			if len(str) > 500 {
				str = str[:500]
			}
			fmt.Println("[*] Detected "+ entry)
			fmt.Println(str)
		}
	}
}

func detectGeneralFiles(u *string) {
	list := [...]string{"README.md", ".htaccess", "robots.txt", }
	for _,v := range list {
		entry := *u+ "/" +v
		resp := tools.DoRequest(entry, tools.ReqParam{})
		defer resp.Body.Close()
		ct := resp.Header.Get("Content-Type")
		if resp.StatusCode == 200 && !strings.Contains(ct, "text/html"){
			body, e := ioutil.ReadAll(resp.Body)
			if e != nil {
				panic(e)
			}
			str := string(body)
			if len(str) > 500 {
				str = str[:500]
			}
			fmt.Println("[*] Detected "+ entry)
			fmt.Println(str)
		}
	}
}

func crossdomain(u *string) {
	entry := *u + "/crossdomain.xml"
	resp := tools.DoRequest(entry, tools.ReqParam{})
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, e := ioutil.ReadAll(resp.Body)
		if e != nil {
			panic(e)
		}
		if strings.Contains(string(body), "cross-domain-policy") && strings.Contains(string(body), "domain=\"*\"") {
			fmt.Println("[*] Detected " + entry)
			fmt.Println(string(body))
		}
	}
}

func robots(u *string) {
	entry := *u+ "/robots.txt"
	resp := tools.DoRequest(entry, tools.ReqParam{})
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, e := ioutil.ReadAll(resp.Body)
		if e != nil {
			panic(e)
		}
		if strings.Contains(string(body), "Disallow") {
			str := string(body)
			if len(str) > 500 {
				str = str[:500]
			}
			fmt.Println("[*] Detected "+ entry)
			fmt.Println(str)
		}
	}
}

func directoryListing(u *string) bool {
	entry := *u
	resp := tools.DoRequest(entry, tools.ReqParam{})
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, e := ioutil.ReadAll(resp.Body)
		if e != nil {
			panic(e)
		}
		if strings.Contains(string(body), "Directory listing for")  {
			log.Println("Detected Directory List.", entry)
			return true
		}
	}
	return false
}