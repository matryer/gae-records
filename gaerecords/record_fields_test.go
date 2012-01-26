package gaerecords

import (
	"testing"
	"appengine"
	"appengine/datastore"
)

func TestSet(t *testing.T) {

	people := CreateTestModel()
	person := people.New()

	// Set() should chain
	assertEqual(t, person, person.Set("name", "Mat"))

	// did field update?
	assertEqual(t, "Mat", person.Fields()["name"])

}

func TestGet(t *testing.T) {

	people := CreateTestModel()
	person := people.New()

	person.Fields()["age"] = 29

	assertEqual(t, 29, person.Get("age"))

}

func TestGetAndSetString(t *testing.T) {

	people := CreateTestModel()
	person := people.New()

	assertEqual(t, person, person.SetString("name", "Mat"))
	assertEqual(t, "Mat", person.Fields()["name"])
	assertEqual(t, "Mat", person.GetString("name"))

	err := person.Put()
	if err != nil {
		t.Errorf("Failed to Put: %v", err)
	}

}

func TestGetAndSetInt(t *testing.T) {

	people := CreateTestModel()
	person := people.New()

	assertEqual(t, person, person.SetInt64("age", 27))
	assertEqual(t, int64(27), person.Fields()["age"])
	assertEqual(t, int64(27), person.GetInt64("age"))

	err := person.Put()
	if err != nil {
		t.Errorf("Failed to Put: %v", err)
	}

}

func TestGetAndSetFloat64(t *testing.T) {

	people := CreateTestModel()
	person := people.New()

	assertEqual(t, person, person.SetFloat64("age", 27.5))
	assertEqual(t, float64(27.5), person.Fields()["age"])
	assertEqual(t, float64(27.5), person.GetFloat64("age"))

	err := person.Put()
	if err != nil {
		t.Errorf("Failed to Put: %v", err)
	}

}

func TestGetAndSetBool(t *testing.T) {

	people := CreateTestModel()
	person := people.New()

	assertEqual(t, person, person.SetBool("field", true))
	assertEqual(t, true, person.Fields()["field"])
	assertEqual(t, true, person.GetBool("field"))

	err := person.Put()
	if err != nil {
		t.Errorf("Failed to Put: %v", err)
	}

}

func TestGetAndSetKeyField(t *testing.T) {

	people := CreateTestModel()
	person := people.New()

	var key *datastore.Key = datastore.NewIncompleteKey(GetAppEngineContext(), "Entity", nil)

	assertEqual(t, person, person.SetKeyField("field", key))
	assertEqual(t, key, person.Fields()["field"])
	assertEqual(t, key, person.GetKeyField("field"))

	err := person.Put()
	if err != nil {
		t.Errorf("Failed to Put: %v", err)
	}

}

func TestGetAndSetTimeField(t *testing.T) {

	people := CreateTestModel()
	person := people.New()

	var time datastore.Time = 1

	assertEqual(t, person, person.SetTime("field", time))
	assertEqual(t, time, person.Fields()["field"])
	assertEqual(t, time, person.GetTime("field"))

	err := person.Put()
	if err != nil {
		t.Errorf("Failed to Put: %v", err)
	}

}

func TestGetAndSetBlobKeyField(t *testing.T) {

	people := CreateTestModel()
	person := people.New()

	var key appengine.BlobKey = "blob"

	assertEqual(t, person, person.SetBlobKey("field", key))
	assertEqual(t, key, person.Fields()["field"])
	assertEqual(t, key, person.GetBlobKey("field"))

	err := person.Put()
	if err != nil {
		t.Errorf("Failed to Put: %v", err)
	}

}

func TestDifferentNumbericalValueTypes(t *testing.T) {

	people := CreateTestModel()
	person := people.New()

	err := person.
		//Set("1", int(1)).
		//Set("2", uint8(1)).     
		//Set("3", uint16(1)).    
		//Set("4", uint32(1)).    
		//Set("5", uint64(1)).    
		//Set("6", int8(1)).      
		//Set("7", int16(1)).     
		//Set("8", int32(1)).     
		Set("9", int64(1)).
		//Set("10", float32(1.1)).   
		Set("11", float64(1.1)).
		//Set("12", complex64(1.1)). 
		//Set("13", complex128(1.1)).
		Set("14", true).
		Put()

	if err != nil {
		t.Errorf("TestDifferentValueTypes failed with error: %v", err)
	}

}

func TestGetMultipleItem(t *testing.T) {

	people := CreateTestModel()
	person := people.New()

	person.
		Set("many-strings", []string{"one", "two", "three"}).
		Put()

	person, _ = people.Find(person.ID())

	assertEqual(t, "one", person.GetMultipleItem("many-strings", 0))
	assertEqual(t, "two", person.GetMultipleItem("many-strings", 1))
	assertEqual(t, "three", person.GetMultipleItem("many-strings", 2))

}

func TestGetMultipleLen(t *testing.T) {

	people := CreateTestModel()
	person := people.New()

	person.
		Set("many-strings", []string{"one", "two", "three"}).
		Set("many-float64s", []float64{float64(1.1), float64(2.2), float64(3.3), float64(4.4), float64(5.5)}).
		Put()

	person, _ = people.Find(person.ID())

	assertEqual(t, 3, person.GetMultipleLen("many-strings"))
	assertEqual(t, 5, person.GetMultipleLen("many-float64s"))

}

func TestPersistingMultipleValues(t *testing.T) {

	people := CreateTestModel()
	person := people.New()

	err := person.
		Set("many-strings", []string{"one", "two", "three"}).
		Set("many-int64s", []int64{1, 2, 3, 4, 5}).
		Set("many-float64s", []float64{float64(1.1), float64(2.2), float64(3.3), float64(4.4), float64(5.5)}).
		Set("many-bools", []bool{true, false, true, false}).
		Put()

	if err != nil {
		t.Errorf("Error when putting multiple values: %v", err)
	}

	// load the record back out
	loaded, err := people.Find(person.ID())

	if err != nil {
		t.Errorf("Error when loading record with multiple values: %v", err)
	}

	assertEqual(t, "one", loaded.GetMultiple("many-strings")[0])
	assertEqual(t, "two", loaded.GetMultiple("many-strings")[1])
	assertEqual(t, "three", loaded.GetMultiple("many-strings")[2])

	assertEqual(t, int64(1), loaded.GetMultiple("many-int64s")[0])
	assertEqual(t, int64(2), loaded.GetMultiple("many-int64s")[1])
	assertEqual(t, int64(3), loaded.GetMultiple("many-int64s")[2])
	assertEqual(t, int64(4), loaded.GetMultiple("many-int64s")[3])
	assertEqual(t, int64(5), loaded.GetMultiple("many-int64s")[4])

	assertEqual(t, float64(1.1), loaded.GetMultiple("many-float64s")[0])
	assertEqual(t, float64(2.2), loaded.GetMultiple("many-float64s")[1])
	assertEqual(t, float64(3.3), loaded.GetMultiple("many-float64s")[2])
	assertEqual(t, float64(4.4), loaded.GetMultiple("many-float64s")[3])
	assertEqual(t, float64(5.5), loaded.GetMultiple("many-float64s")[4])

	assertEqual(t, true, loaded.GetMultiple("many-bools")[0])
	assertEqual(t, false, loaded.GetMultiple("many-bools")[1])
	assertEqual(t, true, loaded.GetMultiple("many-bools")[2])
	assertEqual(t, false, loaded.GetMultiple("many-bools")[3])

}

func TestSetMultiple_StronglyTypedVarients(t *testing.T) {

	people := CreateTestModel()
	person := people.New()

	key1 := people.NewKeyWithID(123)
	key2 := people.NewKeyWithID(456)
	key3 := people.NewKeyWithID(789)
	var time1 datastore.Time = datastore.Time(10)
	var time2 datastore.Time = datastore.Time(20)
	var time3 datastore.Time = datastore.Time(30)
	blobKey1 := appengine.BlobKey("one")
	blobKey2 := appengine.BlobKey("one")
	blobKey3 := appengine.BlobKey("one")

	assertEqual(t, person, person.SetMultipleStrings("many-strings", []string{"one", "two", "three"}))
	assertEqual(t, person, person.SetMultipleInt64s("many-int64s", []int64{1, 2, 3, 4, 5}))
	assertEqual(t, person, person.SetMultipleFloat64s("many-float64s", []float64{float64(1.1), float64(2.2), float64(3.3), float64(4.4), float64(5.5)}))
	assertEqual(t, person, person.SetMultipleBools("many-bools", []bool{true, false, true, false}))
	assertEqual(t, person, person.SetMultipleKeys("many-keys", []*datastore.Key{key1, key2, key3}))
	assertEqual(t, person, person.SetMultipleTimes("many-times", []datastore.Time{time1, time2, time3}))
	assertEqual(t, person, person.SetMultipleBlobKeys("many-blob-keys", []appengine.BlobKey{blobKey1, blobKey2, blobKey3}))

	err := person.Put()

	if err != nil {
		t.Errorf("%v", err)
	} else {

		// reload the item
		loaded, _ := people.Find(person.ID())

		assertEqual(t, "one", loaded.GetMultiple("many-strings")[0])
		assertEqual(t, "two", loaded.GetMultiple("many-strings")[1])
		assertEqual(t, "three", loaded.GetMultiple("many-strings")[2])

		assertEqual(t, int64(1), loaded.GetMultiple("many-int64s")[0])
		assertEqual(t, int64(2), loaded.GetMultiple("many-int64s")[1])
		assertEqual(t, int64(3), loaded.GetMultiple("many-int64s")[2])
		assertEqual(t, int64(4), loaded.GetMultiple("many-int64s")[3])
		assertEqual(t, int64(5), loaded.GetMultiple("many-int64s")[4])

		assertEqual(t, float64(1.1), loaded.GetMultiple("many-float64s")[0])
		assertEqual(t, float64(2.2), loaded.GetMultiple("many-float64s")[1])
		assertEqual(t, float64(3.3), loaded.GetMultiple("many-float64s")[2])
		assertEqual(t, float64(4.4), loaded.GetMultiple("many-float64s")[3])
		assertEqual(t, float64(5.5), loaded.GetMultiple("many-float64s")[4])

		assertEqual(t, true, loaded.GetMultiple("many-bools")[0])
		assertEqual(t, false, loaded.GetMultiple("many-bools")[1])
		assertEqual(t, true, loaded.GetMultiple("many-bools")[2])
		assertEqual(t, false, loaded.GetMultiple("many-bools")[3])

		assertEqual(t, key1.String(), loaded.GetMultiple("many-keys")[0].(*datastore.Key).String())
		assertEqual(t, key2.String(), loaded.GetMultiple("many-keys")[1].(*datastore.Key).String())
		assertEqual(t, key3.String(), loaded.GetMultiple("many-keys")[2].(*datastore.Key).String())

		assertEqual(t, time1, loaded.GetMultiple("many-times")[0])
		assertEqual(t, time2, loaded.GetMultiple("many-times")[1])
		assertEqual(t, time3, loaded.GetMultiple("many-times")[2])

		assertEqual(t, blobKey1, loaded.GetMultiple("many-blob-keys")[0])
		assertEqual(t, blobKey2, loaded.GetMultiple("many-blob-keys")[1])
		assertEqual(t, blobKey3, loaded.GetMultiple("many-blob-keys")[2])

	}

}
