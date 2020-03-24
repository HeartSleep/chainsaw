package baseline

import "net/url"

/**
	This module is a dispatcher.
	12:11 AM, December 9, 2019 in HeFei.
 */

func Start(Url *url.URL) {
	detectGeneralFiles(Url)
	detectFiles(Url)
	crossdomain(Url)
	directoryListing(Url)
	druid(Url)
	laravelDebug(Url)
	CorsCheck(Url)
	springActuator(Url)
}