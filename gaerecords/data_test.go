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
	
	loadedRecord, err2 := LoadOneByID(model, 1)
	
	if err2 != nil {
		t.Errorf("LoadOneByID: %v", err)
	}
	
	assertNotNil(t, record, "record")
	assertNotNil(t, loadedRecord, "loadedRecord")
	
	if record != nil && loadedRecord != nil {
		assertEqual(t, record.ID(), loadedRecord.ID())
	}
	
}