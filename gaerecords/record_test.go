package gaerecords

import (
	"testing"
)

func TestNewRecord(t *testing.T) {

	model := CreateTestModel()
	record := NewRecord(model)

	assertNotNil(t, record, "new(Record)")
	assertNotNil(t, record.Fields(), "record.Fields()")
	assertEqual(t, model, record.Model())

}

func TestNoIDValue(t *testing.T) {

	assertEqual(t, int64(0), NoIDValue)

}

func TestRecordModel(t *testing.T) {
	
	var record *Record = new(Record)
	
	//assertNil(t, record.Model(), "record.Model()")
	
	model := CreateTestModel()
	record.model = model
	
	assertEqual(t, model, record.Model())
	
}

func TestRecordSetModel(t *testing.T) {
	
	var record *Record = new(Record)
	
	model := CreateTestModel()
	
	assertEqual(t, record, record.SetModel(model))	
	assertEqual(t, model, record.model)
	
}

func TestRecordString(t *testing.T) {
	
	model := NewModel("people")
	record := model.New()
	
	assertEqual(t, "{Record:model={Model:people}}", record.String())	
	
	record.setID(123)

	assertEqual(t, "{Record:model={Model:people},id=123}", record.String())	
	
}