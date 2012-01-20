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
	subrecord := model.New()
	
	// set something on the sub-record
	subrecord.SetString("Type", "Subrecord")
	
	// set something on the main record
	
	record.
		SetString("Name", "Mat").
		Set("sub", subrecord).
		Put()
		
	// load the main record again
	loaded, _ := model.Find(record.ID())
	
	assertNotNil(t, loaded, "loaded")
	//assertEqual(t, loaded.Get("sub"), "Subrecord")
	
}