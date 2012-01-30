package gaerecords

// Internal value for the first page
var firstPage int = 1

// PageInfo is a struct that holds details about paging for a model.  Best practice is to
// load this object directly from the model using the model.LoadPageInfo() method.
//
// You can use this object without a model by using the NewPageInfo method.
type PageInfo struct {

	// TotalPages represents the total number of pages
	TotalPages int

	// TotalRecords represents the total number of records
	TotalRecords int

	// RecordsPerPage represents the number of records per page
	RecordsPerPage int

	// CurrentPage represents the current page number
	CurrentPage int

	// HasPreviousPage gets whether there are any pages before the CurrentPage
	HasPreviousPage bool

	// HasNextPage gets whether there are any pages after the CurrentPage
	HasNextPage bool

	// FirstPage gets the page number of the first page (always 1)
	FirstPage int

	// LastPage gets the page number of the last page
	LastPage int

	// RecordsOnLastPage gets the number of records on the last page
	RecordsOnLastPage int
}

// NewPageInfo creates a new PageInfo instance containing the paging details for
// the specified values.
//
//   totalRecords - the total number of records
//   recordsPerPage - the number of records per page
//   currentPage - the current page number
func NewPageInfo(totalRecords, recordsPerPage, currentPage int) PageInfo {

	var info PageInfo

	info.TotalRecords = totalRecords
	info.RecordsPerPage = recordsPerPage
	info.TotalPages = totalRecords / recordsPerPage
	info.RecordsOnLastPage = totalRecords % recordsPerPage

	// if there are any records on the last page, increase
	// the page count
	if info.RecordsOnLastPage > 0 {
		info.TotalPages++
	}

	// we'll always assume one page (even if there are no records)
	if info.TotalPages == 0 {
		info.TotalPages = 1
	}

	info.CurrentPage = currentPage
	info.HasPreviousPage = currentPage > firstPage
	info.HasNextPage = currentPage < info.TotalPages

	info.FirstPage = firstPage
	info.LastPage = info.TotalPages

	if info.RecordsOnLastPage == 0 && info.TotalRecords > 0 {
		info.RecordsOnLastPage = recordsPerPage
	}

	return info

}
