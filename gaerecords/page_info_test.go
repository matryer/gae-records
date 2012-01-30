package gaerecords

import (
	"testing"
)

func TestNewPageInfo(t *testing.T) {

	info := NewPageInfo(50, 10)

	withMessage("TotalRecords")
	assertEqual(t, 50, info.TotalRecords)
	assertEqual(t, 5, info.TotalPages)
	assertEqual(t, 10, info.RecordsPerPage)

	assertEqual(t, true, info.HasNextPage(1))
	assertEqual(t, true, info.HasNextPage(2))
	assertEqual(t, true, info.HasNextPage(3))
	assertEqual(t, true, info.HasNextPage(4))
	assertEqual(t, false, info.HasNextPage(5))

	assertEqual(t, false, info.HasPreviousPage(1))
	assertEqual(t, true, info.HasPreviousPage(2))
	assertEqual(t, true, info.HasPreviousPage(3))
	assertEqual(t, true, info.HasPreviousPage(4))
	assertEqual(t, true, info.HasPreviousPage(5))

	assertEqual(t, 1, info.FirstPage())
	assertEqual(t, 5, info.LastPage())

	info = NewPageInfo(95, 10)
	assertEqual(t, 5, info.RecordsOnLastPage())

	info = NewPageInfo(96, 10)
	assertEqual(t, 6, info.RecordsOnLastPage())

	info = NewPageInfo(91, 10)
	assertEqual(t, 1, info.RecordsOnLastPage())

	info = NewPageInfo(90, 10)
	assertEqual(t, 10, info.RecordsOnLastPage())

	info = NewPageInfo(100, 10)
	assertEqual(t, 10, info.RecordsOnLastPage())

}
