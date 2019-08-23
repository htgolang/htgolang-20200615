package main



import (
	"net/http"
	_ "usertask/controllers"
)

func main() {
	addr := ":9999"

	http.ListenAndServe(addr, nil)
}

