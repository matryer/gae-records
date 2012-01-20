package gaerecords

import (
	"http"
	"appengine"
)

var appEngineContext appengine.Context

// Creates a new appengine.Context object from the given request.
// At least one call to CreateAppEngineContext(*http.Request) is required per
// request to ensure gaerecords uses the correct context.
func CreateAppEngineContext(request *http.Request) appengine.Context {

	if appEngineContext == nil {
		// make a new one
		appEngineContext = appengine.NewContext(request)
	}

	return appEngineContext
}

// Gets the current appengine.Context object used by gaerecords.
func GetAppEngineContext() appengine.Context {

	if appEngineContext == nil {
		panic("gaerecords: Be sure to call CreateAppEngineContext(*http.Request) before using gaerecords capabilities")
	}

	return appEngineContext

}
