package gaerecords

import (
	"testing"
	"os"
)

func TestModelValidators(t *testing.T) {

	model := CreateTestModelWithPropertyType("TestRecordValidators")
	record := model.New()

	called := false
	var calledWithRecord *Record

	// add a validator
	assertEqual(t, model.AddValidator(func(m *Model, r *Record) []os.Error {

		called = true
		calledWithRecord = r

		return []os.Error{os.NewError("Error One"), os.NewError("Error Two")}

	}), model)

	valid, errors := record.IsValid()

	assertEqual(t, false, valid)
	assertEqual(t, "Error One", errors[0].String())
	assertEqual(t, "Error Two", errors[1].String())
	assertEqual(t, true, called)
	assertEqual(t, record, calledWithRecord)

}

func TestValidParentRecordValidator(t *testing.T) {
	
	// valid test
	model := NewModel("TestValidParentRecordValidator_NoParent")
	record := model.New()
	var errors []os.Error
	
	withMessage("Error NOT expected because model does not have parent")
	errors = ValidParentRecordValidator(model, record)
	
	assertEqual(t, 0, len(errors))
	
	// invalid test
	parentModel := NewModel("TestValidParentRecordValidator_Parent")
	model = NewModel("TestValidParentRecordValidator_WithParent")
	model.SetParentModel(parentModel)
	record = parentModel.New()
	
	errors = ValidParentRecordValidator(model, record)
	
	withMessage("Error expected because model has parent but record doesn't")
	if assertEqual(t, 1, len(errors)) {
		assertEqual(t, "gaerecords: Record expected to have a parent record of type \"TestValidParentRecordValidator_Parent\".", errors[0].String())
	}
	
	// valid test again
	record2 := model.New()
	record2.SetParent(record)
	
	errors = ValidParentRecordValidator(model, record2)
	
	withMessage("Error NOT expected because model has parent but so does record")
	assertEqual(t, 0, len(errors))
	
}
