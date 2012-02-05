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
	assertEqual(t, model.AddValidator(func(m *Model, r *Record) os.Error {

		called = true
		calledWithRecord = r

		return os.NewError("Oops, something went wrong")

	}), model)

	valid, errors := record.IsValid()

	assertEqual(t, false, valid)
	assertEqual(t, "Oops, something went wrong", errors[0].String())
	assertEqual(t, true, called)
	assertEqual(t, record, calledWithRecord)

}
