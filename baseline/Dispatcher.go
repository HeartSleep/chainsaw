package baseline

/**
	This module is a baseline check dispatcher.
	We will add more checklist later.
	12:11 AM, December 9, 2019 in HeFei.
 */

func Start(u string) {
	detectFiles(&u)
	robots(&u)
	directoryListing(&u)
	druid(&u)
	laravelDebug(&u)
	springActuator(&u)
}