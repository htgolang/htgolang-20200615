package main

import (
	"net/http"
	_ "todolist/controllers"
)

func main() {
	addr := ":9999"
	http.ListenAndServe(addr, nil)
}
