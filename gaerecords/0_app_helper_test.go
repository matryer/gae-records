package gaerecords

import (
	"os"
	"testing"
	"appengine/datastore"
	"gae-go-testing.googlecode.com/git/appenginetesting"
)

// Creates a test appengine context object
func UseTestAppEngineContext() {

	appEngineContext, _ = appenginetesting.NewContext(nil)

}

func CreateTestRecord() *Record {
	return NewRecord(CreateTestModel())
}

func CreateTestModel() *Model {
	return NewModel("model")
}
func CreateTestModelWithPropertyType(kind string) *Model {
	return NewModel(kind)
}

func CreatePersistedRecord(t *testing.T, model *Model) (*Record, os.Error) {

	context := GetAppEngineContext()
	person := model.New()
	key := person.DatastoreKey()

	person.
		SetString("name", "Mat").
		SetInt64("age", 29)

	newKey, err := datastore.Put(context, key, datastore.PropertyLoadSaver(person))

	t.Logf(">> Created new record with ID: '%v'", newKey)

	if err != nil {
		return nil, err
	}

	// set the person ID
	person.setID(newKey.IntID())

	return person, nil

}

func TestSetup(t *testing.T) {

	UseTestAppEngineContext()
	t.Logf("<<< Test context created %v >>>", appEngineContext)

}
