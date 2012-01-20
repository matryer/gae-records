package gaerecords

import (
	"testing"
)

func TestNewModel(t *testing.T) {

	model := NewModel("kind")

	assertEqual(t, model.RecordType(), "kind")

}

func TestNew(t *testing.T) {

	model := CreateTestModel()
	record := model.New()

	assertNotNil(t, record, "Record shouldn't be nil")

}

func TestModelString(t *testing.T) {
	
	model := NewModel("people")
	
	assertEqual(t, "{Model:people}", model.String())
	
}