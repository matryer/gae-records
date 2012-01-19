package gaerecords

import (
	"testing"
)

func TestSetIDAndGetID(t *testing.T) {
	
	people := CreateTestModel()
	person := people.New()
	
	assertEqual(t, NoIDValue, person.ID())
	assertEqual(t, person, person.setID(123))
	assertEqual(t, int64(123), person.ID())
	
}