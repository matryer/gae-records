package gaerecords

import (
	"testing"
	"gae-go-testing.googlecode.com/git/appenginetesting"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected to be equal. %v != %v", a, b)
	}
}

func CreateTestAppengineContext() *appenginetesting.Context {
	appengineContext, _ := appenginetesting.NewContext(nil)
	return appengineContext
}