package gaerecords

import (
	"testing"
	"gae-go-testing.googlecode.com/git/appenginetesting"
)

func TestCleanup(t *testing.T) {
	
	// close the test context
	if TestContext != nil {
		AppEngineContext(t).(*appenginetesting.Context).Close()
		t.Logf("<<< Test context closed >>>")
	}
	
}