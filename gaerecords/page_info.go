package gaerecords

// Internal value for the first page
var firstPage int = 1

// PageInfo is a struct that holds details about paging for a model.  Best practice is to
// load this object directly from the model using the model.LoadPageInfo() method.
//
// You can use this object without a model by using the NewPageInfo method.
type PageInfo struct {
	TotalPages     int
	TotalRecords   int
	RecordsPerPage int
}

// NewPageInfo creates a new PageInfo instance containing the paging details for
// the specified values.
//
// totalRecords - the total number of records
// recordsPerPage - the number of records per page
func NewPageInfo(totalRecords, recordsPerPage int) *PageInfo {

	info := new(PageInfo)

	info.TotalRecords = totalRecords
	info.RecordsPerPage = recordsPerPage
	info.TotalPages = totalRecords / recordsPerPage

	return info

}

// HasNextPage gets whether there is another page after the specified
// pageNumber.
func (p *PageInfo) HasNextPage(pageNumber int) bool {
	return pageNumber < p.TotalPages
}

// HasPreviousPage gets whether there is another page before the specified
// pageNumber.
func (p *PageInfo) HasPreviousPage(pageNumber int) bool {
	return pageNumber > firstPage
}

// FirstPage gets the page number for the first page.
func (p *PageInfo) FirstPage() int {
	return firstPage
}

// LastPage gets the page number for the last page.
func (p *PageInfo) LastPage() int {
	return p.TotalPages
}

// RecordsOnLastPage gets the number of records on the last page.
// Unless the number of records fits exactly into the pages, the last page will
// usually contain fewer records and this method calculates the value.
func (p *PageInfo) RecordsOnLastPage() int {

	// calculate the floating records after paging
	difference := p.TotalRecords % p.RecordsPerPage

	// if none are left over, the last page is full
	if difference == 0 {
		difference = p.RecordsPerPage
	}

	return difference

}
