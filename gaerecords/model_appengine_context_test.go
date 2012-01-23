package gaerecords

import (
  "testing"
	"http"
	"appengine"
)

func TestModelGetAppengineContext(t *testing.T) {
	
	model := CreateTestModel()
	
	withMessage("by default, model should use the global AppEngineContext")
	assertEqual(t, GetAppEngineContext(), model.AppEngineContext())
	
	// make a new appengine.Context
	request := new(http.Request)
	specificContext := appengine.NewContext(request)
	
	withMessage("model.SetAppEngineContext should chain")
	assertEqual(t, model, model.SetAppEngineContext(specificContext))
	
	withMessage("Model should use the specific AppEngineContext")
	assertEqual(t, specificContext, model.AppEngineContext())
	
	// remove it
	withMessage("model.UseGlobalAppEngineContext should chain")
	assertEqual(t, model, model.UseGlobalAppEngineContext())
	
	withMessage("by default, model should use the global AppEngineContext")
	assertEqual(t, GetAppEngineContext(), model.AppEngineContext())
	
}
