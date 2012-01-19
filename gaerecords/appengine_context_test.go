package gaerecords

import (
	"testing"
	"http"
)

func TestCreateAppEngineContext(t *testing.T) {
	
	appEngineContext = nil
	request := new(http.Request)
	assertNotNil(t, CreateAppEngineContext(request), "GetAppEngineContext()")
	
}

func TestGetAppEngineContext(t *testing.T) {
	
	request := new(http.Request)
	CreateAppEngineContext(request)
	
	assertEqual(t, appEngineContext, GetAppEngineContext())
	
}