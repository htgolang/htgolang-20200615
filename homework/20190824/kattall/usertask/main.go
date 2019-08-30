package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"net/http"
	_ "usertask/controllers"
)
func main(){
	fmt.Println(uuid.NewV4())
	addr := ":9999"
	http.ListenAndServe(addr, nil)
}
