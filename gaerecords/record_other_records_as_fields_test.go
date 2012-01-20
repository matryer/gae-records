package gaerecords

import (
	"testing"
)

func TestRecordAsFieldValue(t *testing.T) {
	
	// create a new model
	model := NewModel("parentRecords")
	
	// create a record
	record := model.New()
	
	// create a sub-record
	recordField := model.New()
	
	// set something on the sub-record
	recordField.SetString("Type", "Subrecord")
	
	// set something on the main record
	
	record.
		SetString("Name", "Mat").
		Set("sub", recordField).
		Put()
	
	// load the main record again
	loaded, err := model.Find(record.ID())
	
	if loaded == nil {
		t.Errorf("model.Find(%v) should find record. Record=%v Err=%v", record.ID(), loaded, err)
	} else {
		assertNotNil(t, loaded, "loaded")

		t.Logf("%v", loaded)

		assertEqual(t, loaded.Get("sub"), "Subrecord")
	}
	
}