package main

import (
	"html/template"
	"os"
)

func main() {
	tpl := template.Must(template.ParseFiles("views/xiangge.html", "views/include.html"))
	tpl.ExecuteTemplate(os.Stdout, "include.html", nil)
}
