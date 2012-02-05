package gaerecords

import (
	"testing"
)

func TestAddModel(t *testing.T) {
	
	model := new(Model)
	model.recordType = "TheRecordType"
	
	// reset models
	models = nil
	
	addModel(model)
	
	if assertEqual(t, 1, len(models)) {

		// make sure the model is added to the 'Models' array
		assertEqual(t, model, models["TheRecordType"])
	
	}
	
}

func TestGetModelByRecordType(t *testing.T) {
	
	// reset models
	models = nil
	
	model1 := NewModel("TestGetModelByRecordType-Model1")
	model2 := NewModel("TestGetModelByRecordType-Model2")
	model3 := NewModel("TestGetModelByRecordType-Model3")
	
	assertEqual(t, model1, getModelByRecordType("TestGetModelByRecordType-Model1"))
	assertEqual(t, model2, getModelByRecordType("TestGetModelByRecordType-Model2"))
	assertEqual(t, model3, getModelByRecordType("TestGetModelByRecordType-Model3"))
	
}

func TestNewModel_Models(t *testing.T) {

	// reset models
	models = nil

	model := NewModel("TestNewModel_Models")

	assertEqual(t, model.RecordType(), "TestNewModel_Models")

	if assertEqual(t, 1, len(models)) {

		// make sure the model is added to the 'Models' array
		assertEqual(t, model, models["TestNewModel_Models"])
		
	}

}
