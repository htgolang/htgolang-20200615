package main

import (
	"net/http"
	_ "todolist/controllers"
)

func main() {
	addr := "0.0.0.0:9999"
	http.ListenAndServe(addr, nil)
}
