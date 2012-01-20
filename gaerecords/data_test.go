package gaerecords

import (
	"testing"
)

func TestLoadOneByID(t *testing.T) {
	
	model := CreateTestModelWithPropertyType("modeltwo")
	record, err := CreatePersistedRecord(t, model)
	
	if err != nil {
		t.Errorf("CreatePersistedRecord: %v", err)
	}
	if record == nil {
		t.Errorf("CreatePersistedRecord didn't create the record")
	}
	
	loadedRecord, err2 := LoadOneByID(model, record.ID())
	
	if err2 != nil {
		t.Errorf("LoadOneByID: %v", err2)
	}
	if loadedRecord == nil {
		t.Errorf("LoadOneByID didn't create the record")
	}
	
	if record != nil && loadedRecord != nil {
		
		assertEqual(t, record.ID(), loadedRecord.ID())
		
		assertEqual(t, "Mat", record.GetString("name"))
		assertEqual(t, int64(29), record.GetInt64("age"))
		assertEqual(t, model.RecordType(), record.Model().RecordType())
		
	}
	
}
