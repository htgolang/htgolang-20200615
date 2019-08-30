package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"todolist/modules"
	"todolist/session"
)

func init() {
	http.HandleFunc("/users/", UsersAction)
	http.HandleFunc("/users/login/", LoginAction)
}

func UsersAction(w http.ResponseWriter, r *http.Request) {
	q := strings.TrimSpace(r.FormValue("query"))
	tpl := template.Must(template.New("users.html").ParseFiles("views/user/users.html"))
	tpl.Execute(w, struct {
		Query string
		Users []modules.User
	}{q, modules.GetUsers(q)})
}

func LoginAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tpl := template.Must(template.New("login.html").ParseFiles("views/user/login.html"))
		tpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		fmt.Println(username, password)
		user, err := modules.GetUserByName(username)
		if err != nil || !user.VaildatePassword(password) {
			tpl := template.Must(template.New("login.html").ParseFiles("views/user/login.html"))
			tpl.Execute(w, struct {
				Name  string
				Error string
			}{username, "用户名或密码不正确"})
		} else {
			fmt.Println("ok", user)
			sessionobj := session.DefaultManager.SessionStart(w, r)
			sessionobj.Set("user", user)
			http.Redirect(w, r, "/tasks/", http.StatusFound)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
