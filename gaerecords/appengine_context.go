package gaerecords

import (
	"http"
	"appengine"
)

var AppEngineContext appengine.Context

// Creates a new appengine.Context object from the given request.
// At least one call to CreateAppEngineContext(*http.Request) is required per
// request to ensure gaerecords uses the correct context.
func CreateAppEngineContext(request *http.Request) appengine.Context {

	if AppEngineContext == nil {
		// make a new one
		AppEngineContext = appengine.NewContext(request)
	}

	return AppEngineContext
}

// Gets the current appengine.Context object used by gaerecords.
func GetAppEngineContext() appengine.Context {

	if AppEngineContext == nil {
		panic("gaerecords: Be sure to call CreateAppEngineContext(*http.Request) before using gaerecords capabilities")
	}

	return AppEngineContext

}
