package main

import (
	"net/http"
	_ "user/controllers"
)

func main(){
	addr := "0.0.0.0:9000"
	http.ListenAndServe(addr,nil)
}