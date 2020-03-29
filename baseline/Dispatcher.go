package baseline

import "net/url"

/**
	Dispatcher
	12:11 AM, December 9, 2019 in HeFei.
 */

func renewUrl(Url *url.URL) *url.URL {
	newUrl := &url.URL{}
	*newUrl = *Url
	return newUrl
}

func Start(Url *url.URL) {
	detectGeneralFiles(renewUrl(Url))
	detectFiles(renewUrl(Url))
	crossdomain(renewUrl(Url))
	druid(renewUrl(Url))
	laravelDebug(renewUrl(Url))
	CorsCheck(renewUrl(Url))
	springActuator(renewUrl(Url))
}