package main

import (
	"log"
	"net/http"
	_ "github.com/xlotz/todolist/controllers"
)

func main() {
	addr := "0.0.0.0:8888"
	// http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
	// 	http.NotFound(w, r)
	// })
	log.Fatal(http.ListenAndServe(addr, nil))
}
