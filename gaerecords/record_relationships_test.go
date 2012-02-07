package gaerecords

import (
	"testing"
)

func TestRecordParent(t *testing.T) {
	
	Records := NewModel("TestRecordParent")
	
	record1 := Records.New()
	record2 := Records.New()
	
	assertEqual(t, false, record1.HasParent())
	assertEqual(t, record1, record1.SetParent(record2))
	assertEqual(t, true, record1.HasParent())
	assertEqual(t, record2, record1.Parent())
	
}

func TestParentRecordDatastoreKey(t *testing.T) {
	
	ParentRecords := NewModel("TestParentRecordDatastoreKey_Parents")
	ChildRecords := ParentRecords.HasMany("TestParentRecordDatastoreKey_Children")
	
	parent := ParentRecords.New()
	parent.Put()
	child := ChildRecords.New().SetParent(parent)
	child.Put()
	
	key := child.DatastoreKey()
	parentKey := key.Parent()
	
	if parentKey == nil {
		t.Errorf("key.Parent() should not be nil when SetParent() is called")
	} else {
	
		assertEqual(t, parent.ID(), parentKey.IntID())
		assertEqual(t, ParentRecords.RecordType(), parentKey.Kind())
		assertEqual(t, child.ID(), key.IntID())
		assertEqual(t, ChildRecords.RecordType(), key.Kind())
	
	}
	
}

func TestHasManyRecord_ValidOnlyWithParent(t *testing.T) {
	
	Authors := NewModel("TestHasManyRecord_ValidOnlyWithParent_Authors")
	Books := Authors.HasMany("TestHasManyRecord_ValidOnlyWithParent_Books")
	DifferentModel := NewModel("TestHasManyRecord_ValidOnlyWithParent_Different")
	
	// create a book with no parent
	
	book := Books.New()
	
	isValid, validErrors := book.IsValid()
	
	withMessage("isValid should be false when no parent is set")
	assertEqual(t, false, isValid)
	
	withMessage("Should be one error")
	if assertEqual(t, 1, len(validErrors)) {
		assertEqual(t, "gaerecords: Record expected to have a parent record of type \"TestHasManyRecord_ValidOnlyWithParent_Authors\".", validErrors[0].String())
	}
	
	// create a book with a parent of the wrong type
	
	wrongParent := DifferentModel.New()
	wrongParent.Put()
	
	book = Books.New()
	book.SetParent(wrongParent)
	isValid, validErrors = book.IsValid()
	
	withMessage("isValid should be false when wrong parent is set")
	assertEqual(t, false, isValid)
	
	withMessage("Should be one error")
	if assertEqual(t, 1, len(validErrors)) {
		assertEqual(t, "gaerecords: Record expected to have a parent record of type \"TestHasManyRecord_ValidOnlyWithParent_Authors\".", validErrors[0].String())
	}
	
	
	// create a book with a parent of the correct type
	
  rightParent := Authors.New()
  rightParent.Put()
	
	book = Books.New()
	book.SetParent(rightParent)
	isValid, validErrors = book.IsValid()
	
	withMessage("isValid should be true")
	assertEqual(t, true, isValid)
	
	withMessage("Should be 0 errors")
	assertEqual(t, 0, len(validErrors))
	
}
