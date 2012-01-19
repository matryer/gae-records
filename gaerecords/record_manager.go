package gaerecords

import (
	"appengine"
)

// Represents a kind of records that can be interacted with
type RecordManager struct {
	recordType string
	appengineContext appengine.Context
}

/*
	Creates a new record manager responsible for managing
	the specified type of data
*/
func NewRecordManager(context appengine.Context, recordType string) *RecordManager {
	m := new(RecordManager)
	m.recordType = recordType
	m.appengineContext = context
	return m
}

// Creates a new record managed by this manager
func (m *RecordManager) New() *Record {
	r := NewRecord()
	r.Manager = m
	return r
}

// Gets a string representing the record type managed by this manager
func (m *RecordManager) RecordType() string {
	return m.recordType
}
