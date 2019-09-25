package utils

import (
	"io"
	"html/template"
)

func Render(writer io.Writer, name string, files []string, context interface{}, funcMap template.FuncMap){
	tpl := template.Must(template.New("tpl").Funcs(funcMap).ParseFiles(files...))
	tpl.ExecuteTemplate(writer, name, context)
}
