package controllers

import (
	"fmt"
	"html/template"
	"loganalyze/models"
	_ "loganalyze/models"
	"net/http"
)

func init() {
	http.HandleFunc("/", LogUpload)
	http.HandleFunc("/query/", LogQuery)

}

func LogUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tpl := template.Must(template.New("upload.html").ParseFiles("views/upload.html"))
		_ = tpl.Execute(w, nil)
		return
	} else if r.Method == http.MethodPost {
		_ = r.ParseMultipartForm(1024)

		if file, header, err := r.FormFile("logfile"); err == nil {
			fmt.Println(header)
			models.LogStore(file)
			_, _ = w.Write([]byte(`<p>success</p><div><a href="/">返回首页</a></div><div><a href="/query/">查询日志</a></div>`))
			return
		}
	}
	http.Redirect(w, r, r.Referer(), 404)
	return
}

func LogQuery(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tpl := template.Must(template.New("query.html").ParseFiles("views/query.html"))
		_ = tpl.Execute(w, models.LogQuery())
		return
	}
	http.Redirect(w, r, r.Referer(), 404)
	return
}
