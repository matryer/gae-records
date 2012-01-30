package gaerecords

import (
	"testing"
)

func TestNewPageInfo(t *testing.T) {

	var info PageInfo

	// exact number of records
	info = NewPageInfo(50, 10, 2)
	assertEqual(t, 50, info.TotalRecords)
	assertEqual(t, 5, info.TotalPages)
	assertEqual(t, 10, info.RecordsPerPage)
	assertEqual(t, 10, info.RecordsOnLastPage)
	assertEqual(t, 1, info.FirstPage)
	assertEqual(t, 5, info.LastPage)
	assertEqual(t, true, info.HasNextPage)
	assertEqual(t, true, info.HasPreviousPage)

	// un-even number of records
	info = NewPageInfo(45, 10, 2)
	assertEqual(t, 45, info.TotalRecords)
	assertEqual(t, 5, info.TotalPages)
	assertEqual(t, 10, info.RecordsPerPage)
	assertEqual(t, 5, info.RecordsOnLastPage)
	assertEqual(t, 1, info.FirstPage)
	assertEqual(t, 5, info.LastPage)
	assertEqual(t, true, info.HasNextPage)
	assertEqual(t, true, info.HasPreviousPage)

	// less than 1 page worth of records
	info = NewPageInfo(6, 10, 1)
	assertEqual(t, 6, info.TotalRecords)
	assertEqual(t, 1, info.TotalPages)
	assertEqual(t, 10, info.RecordsPerPage)
	assertEqual(t, 6, info.RecordsOnLastPage)
	assertEqual(t, 1, info.FirstPage)
	assertEqual(t, 1, info.LastPage)
	assertEqual(t, false, info.HasNextPage)
	assertEqual(t, false, info.HasPreviousPage)

	// no records
	info = NewPageInfo(0, 10, 1)
	assertEqual(t, 0, info.TotalRecords)
	assertEqual(t, 1, info.TotalPages)
	assertEqual(t, 10, info.RecordsPerPage)
	assertEqual(t, 0, info.RecordsOnLastPage)
	assertEqual(t, 1, info.FirstPage)
	assertEqual(t, 1, info.LastPage)
	assertEqual(t, false, info.HasNextPage)
	assertEqual(t, false, info.HasPreviousPage)

}
