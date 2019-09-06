package main

import (
	_ "analyzeLog/controllers"
	"fmt"
	"net/http"
	"os"
)

func main(){
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Println("启动发生错误。")
		os.Exit(-1)
	}
}
