package main

import (
	"net/http"
	_ "github.com/xlotz/todolist/controllers"
)

func main() {
	addr := "0.0.0.0:9999"
	// http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
	// 	http.NotFound(w, r)
	// })
	http.ListenAndServe(addr, nil)
}
