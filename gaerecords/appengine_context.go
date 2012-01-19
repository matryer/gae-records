package gaerecords

import (
	"http"
	"appengine"
)

var appEngineContext appengine.Context

func CreateAppEngineContext(request *http.Request) appengine.Context {
	
	if appEngineContext == nil {
		// make a new one
		appEngineContext = appengine.NewContext(request)
	}
	
	return appEngineContext
}

func GetAppEngineContext() appengine.Context {
	
	if appEngineContext == nil {
		panic("gaerecords: Be sure to call CreateAppEngineContext(*http.Request) before using gaerecords capabilities")
	}
	
	return appEngineContext
	
}