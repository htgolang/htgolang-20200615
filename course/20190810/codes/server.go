package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func test02(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Time: %d", time.Now().Unix())
}

// 处理器
type Test03 struct{}

func (t Test03) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Time: %s", time.Now().Format("2006-01-02 15:04:05"))
}

func main() {

	//定义处理器函数
	http.HandleFunc("/test01/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hi KK"))
	})

	http.HandleFunc("/test02/", test02)

	http.Handle("/test03/", Test03{})
	http.Handle("/test04/", &Test03{})

	http.HandleFunc("/request/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.UserAgent())
		fmt.Println(r.Referer())

		fmt.Println(r.Method, r.URL, r.Proto)
		fmt.Println(r.Header)

		fmt.Println("体:")
		// bytes := make([]byte, 1024)
		// n, _ := r.Body.Read(bytes)
		// fmt.Println(string(bytes[:n]))

		io.Copy(os.Stdout, r.Body)

		w.Write([]byte("Request"))

	})

	// http.Dir => 类型转换
	http.Handle("/", http.FileServer(http.Dir(".")))

	err := http.ListenAndServe("0.0.0.0:9999", nil)
	if err != nil {
		fmt.Println(err)
	}
}
