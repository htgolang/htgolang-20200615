package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"net/http"
	_ "todolist/controllers"
)
func main(){
	fmt.Println(uuid.NewV4())
	addr := ":9999"
	http.ListenAndServe(addr, nil)
}
