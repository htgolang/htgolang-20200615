package utils

import (
	"html/template"
	"io"
)

// 4个参数
func Render(writer io.Writer, name string, files []string, context interface{}) {
	// tpl := template.Must(template.New(name).ParseFiles(files...))
	// tpl.Execute(writer, context)

	tpl := template.Must(template.ParseFiles(files...))
	tpl.ExecuteTemplate(writer, name, context)
}
