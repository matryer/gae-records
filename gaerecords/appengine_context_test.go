package gaerecords

import (
	"testing"
	"http"
)

func TestCreateAppEngineContext(t *testing.T) {

	AppEngineContext = nil
	request := new(http.Request)
	assertNotNil(t, CreateAppEngineContext(request), "GetAppEngineContext()")

}

func TestGetAppEngineContext(t *testing.T) {

	request := new(http.Request)
	CreateAppEngineContext(request)

	assertEqual(t, AppEngineContext, GetAppEngineContext())

	AppEngineContext = nil
	UseTestAppEngineContext()

}
