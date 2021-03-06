package gaerecordsapp

import (
	"goweb"
	"http"
)

var PathToTemplates string = "gaerecordsapp/view_templates"

func mapGowebControllers() {
	
	// create a 'PeopleController'
	peopleController := new(PeopleController)
	
	// map the People resources
	goweb.MapFunc("/people/new", func(c *goweb.Context){
		peopleController.New(c)
	})
	goweb.MapFunc("/people/{{id}}/edit", func(c *goweb.Context){
		peopleController.Edit(c.PathParams["id"], c)
	})
	goweb.MapFunc("/people/{id}/confirm-delete", func(c *goweb.Context){
		peopleController.DeleteConfirm(c.PathParams["id"], c)
	})
	goweb.MapRest("/people", peopleController)
	
}

func init() {
	
	// map the controllers
	mapGowebControllers()
	
	goweb.ConfigureDefaultFormatters()
	
	// ask Goweb to kindly handle all requests
	http.Handle("/", goweb.DefaultHttpHandler)
	
}
