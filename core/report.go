package core

import (
	"chainsaw/network"
	"net/url"
)

type VulMsg struct {
	Url url.URL
	Param network.ReqParam
	Module string
	Detail map[string]string
}

func Report(msg VulMsg) bool{
	//TODO
	if false {
		return false
	}
	return true
}

func WriteToFile() {

}
