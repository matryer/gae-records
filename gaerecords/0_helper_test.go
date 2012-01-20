package gaerecords

import (
	"testing"
)

// Asserts that two objects are equal and throws an error if not
func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected to be equal. %v != %v", a, b)
	}
}

func assertNotNil(t *testing.T, a interface{}, msg string) {
	if a == nil {
		t.Errorf("%v. Expected not to be nil.", msg)
	}
}
func assertNil(t *testing.T, a interface{}, msg string) {
	if a != nil {
		t.Errorf("%v. Expected to be nil but was: %v", msg, a)
	}
}
