package main

import (
	"net/http"
	_ "user/controllers"
)

func main(){
	addr := ":9000"
	http.ListenAndServe(addr,nil)
}