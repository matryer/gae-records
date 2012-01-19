package gaerecords

import (
	"testing"
)

func TestCleanup(t *testing.T) {
	
	TestContext.Close()
	
	t.Logf("<<< Test context closed >>>")
	
}