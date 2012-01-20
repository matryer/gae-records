package gaerecords

import (
	"testing"
)

func TestLoadOneByID(t *testing.T) {
	
	model := CreateTestModel()
	record, err := CreatePersistedRecord(model)
	
	if err != nil {
		t.Errorf("CreatePersistedRecord: %v", err)
	}
	
	loadedRecord, err2 := LoadOneByID(model, record.ID())
	
	if err2 != nil {
		t.Errorf("LoadOneByID: %v", err)
	}
	if loadedRecord == nil {
		t.Errorf("LoadOneByID didn't create the record")
	}
	
	if record != nil && loadedRecord != nil {
		assertEqual(t, record.ID(), loadedRecord.ID())
	}
	
}
