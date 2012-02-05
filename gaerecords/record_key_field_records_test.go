package gaerecords

import (
	"testing"
)

func TestSetRecordField(t *testing.T) {
	
	model := CreateTestModelWithPropertyType("test_key_field_rcords_model1")
	
	husband := model.New()
	husband.SetString("name", "Mat")
	husband.Put()
	
	wife := model.New()
	wife.SetString("name", "Laurie")
	
	assertEqual(t, wife, wife.SetRecordField("spouse", husband))
	
	if !wife.HasField("spouse_key") {
		
		t.Errorf("spouse_key should be set by SetRecordField")
		
	} else {
		
		if wife.GetKeyField("spouse_key") == nil {
			t.Errorf("GetKeyField('spouse_key') should NOT be nil")
		} else {
			assertEqual(t, "test_key_field_rcords_model1", wife.GetKeyField("spouse_key").Kind())
			assertEqual(t, husband.ID(), wife.GetKeyField("spouse_key").IntID())
		}
		
	}
	
}

func TestGetRecordField(t *testing.T) {
	
	model := CreateTestModelWithPropertyType("test_key_field_rcords_model2")
	
	husband := model.New()
	husband.SetString("name", "Mat")
	husband.Put()
	
	wife := model.New()
	wife.SetString("name", "Laurie")
	wife.SetRecordField("spouse", husband)
	wife.Put()
	
	// load the wife back again
	wife, _ = model.Find(wife.ID())
	
	// get the husband
	loadedHusband, _ := wife.GetRecordField(model, "spouse")
	
	assertEqual(t, husband.GetString("name"), loadedHusband.GetString("name"))
	
}
