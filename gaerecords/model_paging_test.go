package gaerecords

import (
	"testing"
	"appengine/datastore"
)

func TestPageInfo(t *testing.T) {

	model := CreateTestModelWithPropertyType("pageInfo")

	// create some records
	for i := 0; i < 50; i++ {
		CreatePersistedRecord(t, model)
	}

	var info *PageInfo
	info, _ = model.LoadPageInfo(10)

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

}

func TestFindByPage(t *testing.T) {

	model := CreateTestModelWithPropertyType("findByPageModel")

	var records []*Record = make([]*Record, 95)

	for i := 0; i < 95; i++ {
		records[i], _ = CreatePersistedRecord(t, model)
	}

	/*
		Page 1
	*/

	// get page 1 (10 records per page)
	page1, err := model.FindByPage(1, 10)

	if err != nil {
		t.Errorf("%v", err)
		return
	}

	assertEqual(t, 10, len(page1))

	assertEqual(t, records[0].ID(), page1[0].ID())
	assertEqual(t, records[1].ID(), page1[1].ID())
	assertEqual(t, records[2].ID(), page1[2].ID())
	assertEqual(t, records[3].ID(), page1[3].ID())
	assertEqual(t, records[4].ID(), page1[4].ID())
	assertEqual(t, records[5].ID(), page1[5].ID())
	assertEqual(t, records[6].ID(), page1[6].ID())
	assertEqual(t, records[7].ID(), page1[7].ID())
	assertEqual(t, records[8].ID(), page1[8].ID())
	assertEqual(t, records[9].ID(), page1[9].ID())

	/*
		Page 2
	*/

	// get page 1 (10 records per page)
	page2, err := model.FindByPage(2, 10)

	if err != nil {
		t.Errorf("%v", err)
		return
	}

	assertEqual(t, 10, len(page2))

	assertEqual(t, records[10].ID(), page2[0].ID())
	assertEqual(t, records[11].ID(), page2[1].ID())
	assertEqual(t, records[12].ID(), page2[2].ID())
	assertEqual(t, records[13].ID(), page2[3].ID())
	assertEqual(t, records[14].ID(), page2[4].ID())
	assertEqual(t, records[15].ID(), page2[5].ID())
	assertEqual(t, records[16].ID(), page2[6].ID())
	assertEqual(t, records[17].ID(), page2[7].ID())
	assertEqual(t, records[18].ID(), page2[8].ID())
	assertEqual(t, records[19].ID(), page2[9].ID())

	/*
		Last page
	*/

	// get page 1 (10 records per page)
	pageLast, err := model.FindByPage(10, 10)

	if err != nil {
		t.Errorf("%v", err)
		return
	}

	assertEqual(t, 5, len(pageLast))

	assertEqual(t, records[90].ID(), pageLast[0].ID())
	assertEqual(t, records[91].ID(), pageLast[1].ID())
	assertEqual(t, records[92].ID(), pageLast[2].ID())
	assertEqual(t, records[93].ID(), pageLast[3].ID())
	assertEqual(t, records[94].ID(), pageLast[4].ID())

}

func TestFindByPage_WithQueryModifier(t *testing.T) {

	model := CreateTestModelWithPropertyType("findByPageModel")

	var records []*Record = make([]*Record, 95)

	for i := 0; i < 95; i++ {

		records[i], _ = CreatePersistedRecord(t, model)

		if i%2 == 0 {
			records[i].SetBool("IsEven", true)
		} else {
			records[i].SetBool("IsEven", false)
		}

		records[i].Put()

	}

	// create the query modifier func
	onlyEvenRecords := func(q *datastore.Query) {
		q.Filter("IsEven=", true)
	}

	/*
		Page 1
	*/

	// get page 1 (10 records per page)
	page1, err := model.FindByPage(1, 10, onlyEvenRecords)

	if err != nil {
		t.Errorf("%v", err)
		return
	}

	assertEqual(t, 10, len(page1))

	assertEqual(t, records[0].ID(), page1[0].ID())
	assertEqual(t, records[2].ID(), page1[1].ID())
	assertEqual(t, records[4].ID(), page1[2].ID())
	assertEqual(t, records[6].ID(), page1[3].ID())
	assertEqual(t, records[8].ID(), page1[4].ID())
	assertEqual(t, records[10].ID(), page1[5].ID())
	assertEqual(t, records[12].ID(), page1[6].ID())
	assertEqual(t, records[14].ID(), page1[7].ID())
	assertEqual(t, records[16].ID(), page1[8].ID())
	assertEqual(t, records[18].ID(), page1[9].ID())

	/*
		Page 2
	*/

	// get page 1 (10 records per page)
	page2, err := model.FindByPage(2, 10, onlyEvenRecords)

	if err != nil {
		t.Errorf("%v", err)
		return
	}

	assertEqual(t, 10, len(page2))

	assertEqual(t, records[20].ID(), page2[0].ID())
	assertEqual(t, records[22].ID(), page2[1].ID())
	assertEqual(t, records[24].ID(), page2[2].ID())
	assertEqual(t, records[26].ID(), page2[3].ID())
	assertEqual(t, records[28].ID(), page2[4].ID())
	assertEqual(t, records[30].ID(), page2[5].ID())
	assertEqual(t, records[32].ID(), page2[6].ID())
	assertEqual(t, records[34].ID(), page2[7].ID())
	assertEqual(t, records[36].ID(), page2[8].ID())
	assertEqual(t, records[38].ID(), page2[9].ID())

	/*
		Last page
	*/

	// get page 1 (10 records per page)
	pageLast, err := model.FindByPage(5, 10, onlyEvenRecords)

	if err != nil {
		t.Errorf("%v", err)
		return
	}

	if assertEqual(t, 8, len(pageLast)) {

		withMessage("last page")
		assertEqual(t, records[80].ID(), pageLast[0].ID())
		assertEqual(t, records[82].ID(), pageLast[1].ID())
		assertEqual(t, records[84].ID(), pageLast[2].ID())
		assertEqual(t, records[86].ID(), pageLast[3].ID())
		assertEqual(t, records[88].ID(), pageLast[4].ID())
		assertEqual(t, records[90].ID(), pageLast[5].ID())
		assertEqual(t, records[92].ID(), pageLast[6].ID())
		assertEqual(t, records[94].ID(), pageLast[7].ID())

	}

}
