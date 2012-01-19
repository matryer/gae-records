package gaerecords

type Model struct {}

func (m *Model) New() *Record {
	return NewRecord(m)
}
