package gaerecords

import (
	"testing"
)

func TestModel(t *testing.T) {
	
	model := CreateTestModel(t)
	assertNotNil(t, model, "new(Model)")
	
}

func TestNew(t *testing.T) {
	
	model := CreateTestModel(t)
	record := model.New()
	
	assertNotNil(t, record, "Record shouldn't be nil")
	
}