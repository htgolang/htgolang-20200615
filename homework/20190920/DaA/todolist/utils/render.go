package utils

import (
	"html/template"
	"io"
)

func Render(w io.Writer, html string, html_list []string, data interface{}) {
	//fmt.Println(w, html, html_list, data)
	tpl := template.Must(template.ParseFiles(html_list...))
	tpl.ExecuteTemplate(w, html, data)
}
