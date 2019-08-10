package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	urllib "net/url"
	"os"
)

func main() {
	url := "https://www.baidu.com"

	// Head
	response, err := http.Get(url)
	if err == nil {
		fmt.Println(response.Proto, response.Status)
		fmt.Println(response.Header)

		io.Copy(os.Stdout, response.Body)
	}

	json := bytes.NewReader([]byte(`{"name" : "kk", "password" : "1234567"}`))
	response, err = http.Post(url, "application/json", json)

	if err == nil {
		fmt.Println(response.Proto, response.Status)
		fmt.Println(response.Header)

		io.Copy(os.Stdout, response.Body)
	}

	params := make(urllib.Values)
	params.Add("name", "kk")
	params.Add("password", "kk123")

	response, err = http.PostForm(url, params)
	if err == nil {
		fmt.Println(response.Proto, response.Status)
		fmt.Println(response.Header)

		//io.Copy(os.Stdout, response.Body)
		response.Write(os.Stdout)
	}
}
