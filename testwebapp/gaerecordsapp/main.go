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
	goweb.MapRest("/people", peopleController)
	
}

func init() {
	
	// map the controllers
	mapGowebControllers()
	
	goweb.ConfigureDefaultFormatters()
	
	// ask Goweb to kindly handle all requests
	http.Handle("/", goweb.DefaultHttpHandler)
	
}
