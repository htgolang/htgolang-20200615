package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/v1/", func(w http.ResponseWriter, r *http.Request) {

		// 显示数据
		cxt := `
		显示数据: {{ . }}
		`

		tpl, err := template.New("tpl").Parse(cxt)
		if err == nil {
			tpl.Execute(w, "Hello World")
		} else {
			fmt.Println(err)
		}
	})

	http.HandleFunc("/v2/", func(w http.ResponseWriter, r *http.Request) {

		// 显示数据
		cxt := `
		显示数据: {{ . }}

		显示数据: {{ index . 1 }}
		`

		tpl, err := template.New("tpl").Parse(cxt)
		if err == nil {
			tpl.Execute(w, []string{"Hello World", "kk"})
		} else {
			fmt.Println(err)
		}
	})

	http.HandleFunc("/v3/", func(w http.ResponseWriter, r *http.Request) {

		// 显示数据
		cxt := `
		显示数据: {{ . }}

		显示数据: {{ .a }}
		`

		tpl, err := template.New("tpl").Parse(cxt)
		if err == nil {
			tpl.Execute(w, map[string]string{"a": "Hello World", "b": "kk"})
		} else {
			fmt.Println(err)
		}
	})

	http.HandleFunc("/v4/", func(w http.ResponseWriter, r *http.Request) {

		// 显示数据
		cxt := `
		显示数据: {{ . }}

		{{ if . }}
		True
		{{ else }}
		False
		{{ end }}

		`

		tpl, err := template.New("tpl").Parse(cxt)
		if err == nil {
			tpl.Execute(w, false)
		} else {
			fmt.Println(err)
		}
	})

	http.HandleFunc("/v5/", func(w http.ResponseWriter, r *http.Request) {

		// 显示数据
		cxt := `
		显示数据: {{ . }}

		{{ if ge . 90 }}
		优秀
		{{ else if ge . 60 }}
		及格
		{{ else }}
		不及格
		{{ end }}
		`
		// ge >=
		// gt >
		// le <=
		// lt <
		// eq ==
		// ne !=

		tpl, err := template.New("tpl").Parse(cxt)
		if err == nil {
			tpl.Execute(w, 50)
		} else {
			fmt.Println(err)
		}
	})

	http.HandleFunc("/v6/", func(w http.ResponseWriter, r *http.Request) {

		// 显示数据
		cxt := `
		显示数据: {{ . }}

		{{ range . }}
		元素: {{ . }}
		{{ end }}
		`

		tpl, err := template.New("tpl").Parse(cxt)
		if err == nil {
			tpl.Execute(w, []string{"kk", "小贩", "xq"})
		} else {
			fmt.Println(err)
		}
	})

	http.HandleFunc("/v7/", func(w http.ResponseWriter, r *http.Request) {

		// 显示数据
		cxt := `
		{{ $name := "kk" }}
		显示数据: {{ . }} {{ $name }}

		{{ range $key, $value := . }}
		元素: {{ . }} {{ $key }} {{ $value }}
		{{ end }}
		`

		tpl, err := template.New("tpl").Parse(cxt)
		if err == nil {
			tpl.Execute(w, map[string]string{"a": "kk", "b": "小贩", "c": "xq"})
		} else {
			fmt.Println(err)
		}
	})

	http.HandleFunc("/v8/", func(w http.ResponseWriter, r *http.Request) {

		// 显示数据
		cxt := `
		{{ index . 0  }}
		{{ 0|index .  }}
		{{ len .  }}
		{{ .|len  }}
		{{ printf "%T" . }}
		{{ .|printf "%T" }}
		{{ "abc"|upper }}
		{{ "abc"|upper02 }}
		{{ upper "abc" }}
		{{ upper02 "abc" }}
		`

		funcs := template.FuncMap{
			"upper": func(txt string) string {
				return strings.ToUpper(txt)
			},
			"upper02": strings.ToUpper,
		}

		tpl, err := template.New("tpl").Funcs(funcs).Parse(cxt)
		if err == nil {
			tpl.Execute(w, []string{"kk", "小贩", "xq"})
		} else {
			fmt.Println(err)
		}
	})

	http.HandleFunc("/v9/", func(w http.ResponseWriter, r *http.Request) {
		tpl, err := template.New("tpl.html").ParseFiles("tpl.html")
		if err == nil {
			fmt.Println(tpl)
			tpl.Execute(w, map[string]map[string]string{"tasks": {"a": "kk", "B": "小贩", "C": "xq"}})
		} else {
			fmt.Println(err)
		}
	})

	http.ListenAndServe(":9999", nil)
}
