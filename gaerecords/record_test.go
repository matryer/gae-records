package gaerecords

import (
	"testing"
)

func TestNewRecord(t *testing.T) {
	
	model := CreateTestModel(t)
	record := NewRecord(model)
	
	assertNotNil(t, record, "new(Record)")
	assertNotNil(t, record.Fields(), "record.Fields()")
	assertEqual(t, model, record.Model())
	
}

func TestNoIDValue(t *testing.T) {
	
	assertEqual(t, int64(0), NoIDValue)
	
}