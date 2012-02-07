package gaerecords

import (
	"os"
	"fmt"
)

// ValidatorFunc is a func that acts as a validator for records.  It takes the model and record,
// and returns an array of errors that are returned when Record.IsValid() is called.
type ValidatorFunc func(*Model, *Record) []os.Error

// ValidParentRecordValidator is a ValidatorFunc that ensures the record has a valid parent
// if the model.HasParentModel()
var ValidParentRecordValidator ValidatorFunc = func(m *Model, r *Record) []os.Error {
	
	// does this model have a parent model?
	if m.HasParentModel() {
		
		var valid bool = true
		
		if !r.HasParent() {
			valid = false
		} else if r.Parent().model != m.ParentModel() {			
			valid = false
		}
		
		// does the record have a matching parent record?
		if !valid {
			return []os.Error{ os.NewError(fmt.Sprintf("gaerecords: Record expected to have a parent record of type \"%v\".", m.ParentModel().RecordType())) }
		}
		
	}
	
	// all's well
	return nil
}