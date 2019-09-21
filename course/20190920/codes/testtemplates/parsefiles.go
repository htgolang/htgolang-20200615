package main

import (
	"html/template"
	"os"
)

func main() {
	//有block在前，define在后
	tpl := template.Must(template.New("index.html").ParseFiles("views/index.html", "views/xiangge.html"))
	tpl.Execute(os.Stdout, nil)
}
