package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	addr := ":9999"

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		txt := `
		<!DOCTYPE html>
		<html>
			<head>
				<meta charset="utf-8" />
				<title>注册</title>
			</head>
			<body>
				<form action="/register/?a=b&c=d" method="POST" enctype="multipart/form-data">
					<label>用户名:</label><input type="text" name="username" />
					<label>密码:</label><input type="password" name="password" />
					<label>头像:</label><input type="file" name="img" />
					<input type="submit" value="注册" />
				</form>
			</body>
		</html>
		`
		fmt.Fprint(response, txt)
	})

	http.HandleFunc("/register/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Println(request.Method, request.URL)

		// 对URL/Body参数解析
		// request.ParseForm()

		request.ParseMultipartForm(1024)

		fmt.Println(request.Form)     //会放url和body所有数据
		fmt.Println(request.PostForm) // 只会放body
		fmt.Println(request.MultipartForm.Value)
		fmt.Println(request.MultipartForm.File)

		fmt.Printf("%T\n", request.MultipartForm.File["img"][0])
		if file, header, err := request.FormFile("img"); err == nil {
			fmt.Printf("%T, %T, %v", file, header, err)
			fmt.Println(header.Filename, header.Header, header.Size)
			newFile, err := os.Create("1")
			if err == nil {
				defer newFile.Close()
				io.Copy(newFile, file)
			}
			fmt.Println("ok")
		} else {
			fmt.Println(err)
		}
		// //FormValue / PostFormValue

		// fmt.Println(request.FormValue("username")) // 调用ParseForm
		// fmt.Println(request.PostFormValue("username"))

		// fmt.Println(request.Form["username"])
		// fmt.Println(request.Form["password"])
		// fmt.Println(request.Form["username"][0])
		// fmt.Println(request.Form["password"][0])
		// fmt.Println(request.Form.Get("username"))
		// fmt.Println(request.Form.Get("password"))
		// fmt.Fprintf(response, "ok")
		http.Redirect(response, request, "/login/", 301)
	})

	http.HandleFunc("/login/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(response, "用户登陆")
	})

	http.ListenAndServe(addr, nil)
}
