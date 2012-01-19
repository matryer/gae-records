package gaerecords

import (
	"testing"
)

func TestEndToEnd(t *testing.T) {
	
	// create a record manager to manage 'Books'
	books := NewRecordManager(AppEngineContext(t), "books")
	
	// create a new book
	originOfSpecies := books.New()
	
	// set some data
	originOfSpecies.
		Set("author", "Charles Darwin").
		Set("publishDate", "24 November 1859").
		Set("originalTitle", "On the Origin of Species by Means of Natural Selection, or the Preservation of Favoured Races in the Struggle for Life")

	// save it
	
	
}