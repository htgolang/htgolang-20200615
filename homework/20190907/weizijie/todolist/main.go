package main

import (
	"net/http"
	"fmt"
	"github.com/satori/go.uuid"
	_ "todolist/controllers"
)

func main() {
	fmt.Println(uuid.NewV4())
	
	addr := "0.0.0.0:9999"
	http.ListenAndServe(addr, nil)
}
