package core

import (
	"chainsaw/network"
	"log"
	"net/url"
	"time"
)

func CheckUrl(Url *url.URL) bool {
	if Url.Scheme!="http" && Url.Scheme!="https" {
		panic("Protocol missing.")
	}
	if Url.Host == "" {
		panic("Host missing.")
	}
	if Url.Port() == "80" || Url.Port() == "443" {
		log.Println("DO NOT add http's default port, it may affect the accuracy!")
	}
	resp := network.DoRequest(Url, network.ReqParam{Timeout: 10 * time.Second})
	if resp == nil {
		return false
	}
	defer resp.Body.Close()
	return true
}