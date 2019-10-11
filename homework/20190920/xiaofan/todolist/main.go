package main

import (
	"fmt"
	"net/http"
	_ "todolist/controllers"
)

func main() {
	addr := ":9999"
	serve := http.ListenAndServe(addr, nil)
	fmt.Println(serve)
}
