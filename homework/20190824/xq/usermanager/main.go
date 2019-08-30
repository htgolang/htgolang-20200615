package main

import (
"net/http"
// 初始化时加载
_ "github.com/xlotz/usermanager/controllers"
)

func main(){

	addr := ":10001"
	http.ListenAndServe(addr, nil)
}
