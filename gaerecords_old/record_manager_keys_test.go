package gaerecords

import (
	"testing"
)

func TestGetKey(t *testing.T) {
	
	people := CreateTestPeopleRecordManager(t)
	key := people.NewKey()
	
	assertEqual(t, people.RecordType(), key.Kind())
	assertEqual(t, true, key.Incomplete())
	
}

func TestNewKeyWithID(t *testing.T) {
	
	people := CreateTestPeopleRecordManager(t)
	key := people.NewKeyWithID(123)
	
	assertEqual(t, people.RecordType(), key.Kind())
	assertEqual(t, int64(123), key.IntID())
	assertEqual(t, false, key.Incomplete())
	
}