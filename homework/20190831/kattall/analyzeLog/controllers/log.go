package controllers

import (
	"analyzeLog/modules"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

const pageCount = 15

func init(){
	http.HandleFunc("/", UploadLog)
	http.HandleFunc("/query/", QueryLog)
	http.HandleFunc("/query/ipcount/", IpCount)
	http.HandleFunc("/query/statuscount/", StatusCount)
	http.HandleFunc("/query/ipurlcount/", IpUrlCount)
}


func UploadLog(w http.ResponseWriter, r *http.Request) {
	var msg struct{
		Message string
	}
	if r.Method == http.MethodGet {
		tpl := template.Must(template.New("upload.html").ParseFiles("views/upload.html"))
		tpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		r.ParseMultipartForm(1024)
		if file, _, err:= r.FormFile("log"); err == nil {
			fmt.Println("uploadlog:", file)
			//newfile, err := os.Create(filepath.Join("files", header.Filename))
			if err == nil {
				//defer newfile.Close()
				//io.Copy(newfile, file)
				err := modules.StoreLog(file)
				fmt.Println("modules.storelog:", err)
				if err != nil {
					fmt.Println(err)
				}
				msg.Message = "上传成功"
			}
		} else {
			msg.Message = "上传失败"
		}
		tpl := template.Must(template.New("upload.html").ParseFiles("views/upload.html"))
		tpl.Execute(w, msg)
	}
}

func QueryLog(w http.ResponseWriter, r *http.Request) {
	var currentPage int
	index := r.FormValue("p")
	fmt.Println("index:", index)
	if index == "" {
		currentPage = 1
	}  else {
		currentPage, _ = strconv.Atoi(index)
	}
	fmt.Println("currentPage:", currentPage)
	fmt.Println("pages:", modules.GetPage(pageCount))
	tpl := template.Must(template.New("query.html").ParseFiles("views/query.html"))
	tpl.Execute(w, struct {
		 Pages int
		 Logs []modules.Log
	}{modules.GetPage(pageCount), modules.GetLog(pageCount, currentPage)})
}

func IpCount(w http.ResponseWriter, r *http.Request){
	tpl := template.Must(template.New("ipcount.html").ParseFiles("views/ipcount.html"))
	tpl.Execute(w, modules.GetIPCountLog())
}

func StatusCount(w http.ResponseWriter, r *http.Request){
	tpl := template.Must(template.New("statuscount.html").ParseFiles("views/statuscount.html"))
	tpl.Execute(w, modules.GetStatusCountLog())
}

func IpUrlCount(w http.ResponseWriter, r *http.Request){
	tpl := template.Must(template.New("ipurlcount.html").ParseFiles("views/ipurlcount.html"))
	tpl.Execute(w, modules.IpUrlCount())
}