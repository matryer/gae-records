package gaerecordsapp

import (
	"http"
	"fmt"
	"github.com/hoisie/mustache.go"
)

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	
	filename := fmt.Sprint(PathToTemplates, "/", tmpl, ".html.mustache")
	layoutFilename := fmt.Sprint(PathToTemplates, "/Shared/Layout.html.mustache")
	
	w.Write([]byte(mustache.RenderFileInLayout(filename, layoutFilename, data)))
	
}
