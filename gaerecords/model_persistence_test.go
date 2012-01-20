package gaerecords

import (
	"testing"
)

func TestFind(t *testing.T) {
	
	model := CreateTestModel()
	record, _ := CreatePersistedRecord(t, model)
	
	loadedRecord, err := model.Find(record.ID())
	
	if err != nil {
		t.Errorf("Model.Find: %v", err)
	}
	
	assertEqual(t, record.ID(), loadedRecord.ID())
	
}