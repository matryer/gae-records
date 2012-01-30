package gaerecords

import (
	"testing"
)

var NextTestMessage string = ""

func withMessage(m string) {
	NextTestMessage = m
}
func getMessage() string {
	m := NextTestMessage
	NextTestMessage = ""
	return m
}

// Asserts that two objects are equal and throws an error if not
func assertEqual(t *testing.T, a interface{}, b interface{}) bool {
	m := getMessage()
	if a != b {
		t.Errorf("Expected to be equal. %v != %v. %v", a, b, m)
		return false
	}
	return true
}

func assertNotNil(t *testing.T, a interface{}, msg string) bool {
	m := getMessage()
	if a == nil {
		t.Errorf("%v. Expected not to be nil. %v", msg, m)
		return false
	}
	return true
}
func assertNil(t *testing.T, a interface{}, msg string) bool {
	m := getMessage()
	if a != nil {
		t.Errorf("%v. Expected to be nil but was: %v. %v", msg, a, m)
		return false
	}
	return true
}
