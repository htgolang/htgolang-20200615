package main

import (
	"fmt"
	"net/http"
	_ "todolist/controllers"

	"github.com/satori/go.uuid"
)

func main() {
	fmt.Println(uuid.NewV4())

	addr := ":9990"
	http.ListenAndServe(addr, nil)
}
