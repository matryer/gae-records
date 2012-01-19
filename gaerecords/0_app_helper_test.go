package gaerecords

import (
	"testing"
	"appengine"
	"gae-go-testing.googlecode.com/git/appenginetesting"
)

var TestContext *appenginetesting.Context

// Creates a test appengine context object
func AppEngineContext(t *testing.T) appengine.Context {
	
	if TestContext == nil {
		t.Logf("<<< Test context created >>>")
		TestContext, _ = appenginetesting.NewContext(nil)
	}
	
	return TestContext
	
}

func CreateTestRecord(t *testing.T) *Record {
	return NewRecord(CreateTestModel(t))
}

func CreateTestModel(t *testing.T) *Model {
	return new(Model)
}