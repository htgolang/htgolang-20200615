package main

import (
	"log"
	_ "loganalyze/controllers"
	"net/http"
	"os"
)

func main() {
	addr := ":8888"
	serve := http.ListenAndServe(addr, nil)

	if serve != nil {
		log.Printf("listen %s error", addr)
		os.Exit(-1)
	}
}
