package gaerecords

import (
	"testing"
)

func TestGetKey(t *testing.T) {
	
	people := CreateTestPeopleRecordManager(t)
	key := people.GetKey(nil)
	
	assertEqual(t, people.RecordType(), key.Kind())
	
	key2 := people.GetKey(key)
	assertEqual(t, key, key2.Parent())
	
}

func TestGetKeyWithID(t *testing.T) {
	
	people := CreateTestPeopleRecordManager(t)
	key := people.GetKeyWithID(123, nil)
	
	assertEqual(t, people.RecordType(), key.Kind())
	assertEqual(t, int64(123), key.IntID())
	
	key2 := people.GetKeyWithID(456, key)
	assertEqual(t, key, key2.Parent())
	assertEqual(t, int64(456), key2.IntID())
	
}