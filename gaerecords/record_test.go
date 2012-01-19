package gaerecords

import (
	"testing"
)

func TestNewRecord(t *testing.T) {
	
	model := CreateTestModel(t)
	record := NewRecord(model)
	
	assertNotNil(t, record, "new(Record)")
	assertNotNil(t, record.Fields, "record.Fields")
	assertEqual(t, model, record.Model)
	
}
