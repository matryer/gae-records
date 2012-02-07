package gaerecords

import (
	"http"
	"appengine"
)

// The AppEngineContext that will be used by gaerecords by default.  Use
// CreateAppEngineContext helper method, or directly assign to this variable before
// calling methods that interact with the datastore.
//
// You can assign an appengine.Context to a specific model at any time,
// in which case that appengine.Context will be preferred.
var AppEngineContext appengine.Context

// Creates a new appengine.Context object from the given request.
// At least one call to CreateAppEngineContext(*http.Request) is required per
// request to ensure gaerecords uses the correct context.  Or you can directly
// assign to the AppEngineContext variable.
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
		Panic("Be sure to call CreateAppEngineContext(*http.Request) before using gaerecords capabilities")
	}

	return AppEngineContext

}
