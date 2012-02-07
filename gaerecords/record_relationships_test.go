package gaerecords

import (
	"testing"
	"os"
	"appengine/datastore"
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
	
	putErr := child.Put()
	if putErr != nil {
		t.Errorf("Couldn't put: %v", putErr)
	}
	
	// reload it
	var err os.Error
	child, err = parent.Find(ChildRecords, child.ID())
	
	if err != nil {
		t.Errorf("Couldn't find child record after Put(): %v", err)
	} else {
		
		key = child.DatastoreKey()
		parentKey = key.Parent()
	
		if parentKey == nil {
			t.Errorf("key.Parent() should not be nil when SetParent() is called")
		} else {
	
			assertEqual(t, parent.ID(), parentKey.IntID())
			assertEqual(t, ParentRecords.RecordType(), parentKey.Kind())
			assertEqual(t, child.ID(), key.IntID())
			assertEqual(t, ChildRecords.RecordType(), key.Kind())
	
		}
	
	}
	
}

func TestRecordFind(t *testing.T) {
	
	ParentRecords := NewModel("TestParentRecordDatastoreKey_Parents2")
	ChildRecords := ParentRecords.HasMany("TestParentRecordDatastoreKey_Children2")
	
	parent := ParentRecords.New()
	parent.Put()
	
	child := ChildRecords.New().SetParent(parent)
	child.SetString("name", "Timmy")
	child.Put()
	
	// load it back
	loadedChild, _ := parent.Find(ChildRecords, child.ID())
	
	assertEqual(t, child.ID(), loadedChild.ID())
	assertEqual(t, "Timmy", loadedChild.GetString("name"))
	
}

func TestRecordFindByQuery_WithQuery(t *testing.T) {
	
	ParentRecords := NewModel("TestParentRecordDatastoreKey_Parents3")
	ChildRecords := ParentRecords.HasMany("TestParentRecordDatastoreKey_Children3")
	
	parent := ParentRecords.New()
	parent.Put()
	
	child1 := ChildRecords.New().SetParent(parent)
	child1.SetString("name", "Timmy")
	child1.Put()

	child2 := ChildRecords.New().SetParent(parent)
	child2.SetString("name", "Tommy")
	child2.Put()

	// load them
	query := datastore.NewQuery(ChildRecords.RecordType())
	children, _ := ChildRecords.FindByQuery(query)
	
	assertEqual(t, child1.ID(), children[0].ID())
	assertEqual(t, "Timmy", children[0].GetString("name"))
	assertEqual(t, child2.ID(), children[1].ID())
	assertEqual(t, "Tommy", children[1].GetString("name"))
	
}

func TestRecordFindByQuery_WithQueryFunc(t *testing.T) {
	
	ParentRecords := NewModel("TestParentRecordDatastoreKey_Parents4")
	ChildRecords := ParentRecords.HasMany("TestParentRecordDatastoreKey_Children4")
	
	parent := ParentRecords.New()
	parent.Put()
	
	child1 := ChildRecords.New().SetParent(parent)
	child1.SetString("name", "Timmy")
	child1.Put()

	child2 := ChildRecords.New().SetParent(parent)
	child2.SetString("name", "Tommy")
	child2.Put()

	// load them
	var query *datastore.Query
	children, _ := ChildRecords.FindByQuery(func(q *datastore.Query){
		query = q
	})
		
	assertEqual(t, child1.ID(), children[0].ID())
	assertEqual(t, "Timmy", children[0].GetString("name"))
	assertEqual(t, child2.ID(), children[1].ID())
	assertEqual(t, "Tommy", children[1].GetString("name"))
	
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

func TestLoadChildRecord_ThenGetParent(t *testing.T) {
	
	Authors := NewModel("TestHasManyRecord_ValidOnlyWithParent_Authors5")
	Books := Authors.HasMany("TestHasManyRecord_ValidOnlyWithParent_Books5")

	darwin := Authors.New()
	darwin.SetString("name", "Charles")
	darwin.Put()
	
	originOfSpecies := Books.New().SetParent(darwin)
	originOfSpecies.Put()
	
	loadedRecords, _ := Books.FindAll()
	loaded := loadedRecords[0]
	
	parent := loaded.Parent()
	
	if parent == nil {
		t.Errorf(".Parent() of loaded record (with parent) shouldn't be nil")
	} else {
	
		assertEqual(t, darwin.ID(), parent.ID())
		assertEqual(t, darwin.GetString("name"), parent.GetString("name"))
		
	}
	
}

