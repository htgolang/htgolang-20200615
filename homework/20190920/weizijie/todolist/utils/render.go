package utils

import (
	"net/http"
	"text/template"
)

func Render(w http.ResponseWriter, name string, files []string, context interface{}) {
	tpl := template.Must(template.New(name).ParseFiles(files...))
	tpl.Execute(w, context)
	// tpl := template.Must(template.ParseFiles(files...))
	// tpl.ExecuteTemplate(w, name, context)
}
