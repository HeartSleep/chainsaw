package core

import "net/url"

type VulMsg struct {
	Url url.URL
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

