package utils

import (
	"html/template"
	"io"
)

func Render(w io.Writer, files []string, name string, context interface{}) error {
	tpl := template.Must(template.ParseFiles(files...))
	return tpl.ExecuteTemplate(w, name, context)
}
