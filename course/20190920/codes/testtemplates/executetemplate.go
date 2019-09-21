package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	//有block在前，define在后
	tpl := template.Must(template.ParseFiles("views/index.html", "views/xiangge.html"))
	tpl.ExecuteTemplate(os.Stdout, "index.html", nil)
	fmt.Println("----------------")
	for _, tpl := range tpl.Templates() {
		fmt.Println(tpl.Name())
	}
}
