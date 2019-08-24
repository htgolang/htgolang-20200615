package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"todolist/models"
	"todolist/session"
)

func LoginAction(w http.ResponseWriter, r *http.Request) {
	var datas struct {
		Name  string
		Error string
	}

	if r.Method == http.MethodGet {
		tpl := template.Must(template.New("login.html").ParseFiles("views/login.html"))
		tpl.Execute(w, datas)
	} else if r.Method == http.MethodPost {
		name := r.PostFormValue("name")
		password := r.PostFormValue("password")
		if user, err := models.GetUserByName(name); err == nil && user.ValidatePassword(password) {
			session := session.DefaultSessionManager.StartSession(w, r)
			session.Set("user", user)
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			tpl := template.Must(template.New("login.html").ParseFiles("views/login.html"))
			datas.Name = name
			datas.Error = "用户名或密码错误"
			tpl.Execute(w, datas)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func TestAction(w http.ResponseWriter, r *http.Request) {

	session := session.DefaultSessionManager.StartSession(w, r)
	fmt.Println(session.Get("user"))
	fmt.Fprint(w, "ok")
}

func LogoutAction(w http.ResponseWriter, r *http.Request) {
	session.DefaultSessionManager.DesotrySession(w, r)
}

func init() {
	http.HandleFunc("/user/login/", LoginAction)
	http.HandleFunc("/user/test/", TestAction)
	http.HandleFunc("/user/logout/", LogoutAction)
}
