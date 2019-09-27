package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"todolist/models"
	"todolist/session"
	"todolist/utils"
)

func UserAction(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultManager.SessionStart(w, r)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/user/login/", http.StatusFound)
		return
	}

	q := strings.TrimSpace(r.FormValue("q"))
	utils.Render(w, "base.html", []string{"views/layouts/base.html", "views/user/user.html"}, struct {
		Query string
		Users []models.User
	}{q, models.GetUsers(q)})
	fmt.Println(models.GetUsers(q))
}

func LoginAction(w http.ResponseWriter, r *http.Request) {
	context := struct {
		Name  string
		Error string
	}{"", ""}

	if r.Method == http.MethodGet {
		tpl := template.Must(template.New("login.html").ParseFiles("views/user/login.html"))
		tpl.Execute(w, context)
	} else if r.Method == http.MethodPost {
		name := r.PostFormValue("name")
		password := r.PostFormValue("password")

		user, err := models.GetUserByName(name)
		if err != nil || !user.ValidatePassword(password) {
			tpl := template.Must(template.New("login.html").ParseFiles("views/user/login.html"))
			tpl.Execute(w, struct {
				Name  string
				Error string
			}{name, "用户名或密码错误"})
		} else {
			sessionObj := session.DefaultManager.SessionStart(w, r)
			//登陆成功
			sessionObj.Set("user", user)
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}
}

func LogoutAction(w http.ResponseWriter, r *http.Request) {
	session.DefaultManager.SessionDestory(w, r)
	http.Redirect(w, r, "/user/login/", http.StatusFound)
}

func UserCreateAction(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultManager.SessionStart(w, r)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/user/login/", http.StatusFound)
		return
	}

	var context interface{}
	if r.Method == http.MethodPost {
		name := r.PostFormValue("name")
		password := r.PostFormValue("password")
		birthday := r.PostFormValue("birthday")
		tel := r.PostFormValue("tel")
		desc := r.PostFormValue("desc")
		addr := r.PostFormValue("addr")

		errors := models.ValidateCreateUser(name, password, birthday, tel, addr, desc)
		if len(errors) == 0 {
			models.CreateUser(name, password, birthday, tel, addr, desc)
			http.Redirect(w, r, "/users/", http.StatusFound)
			return
		}

		context = struct {
			Errors   map[string]string
			Name     string
			Password string
			Birthday string
			Tel      string
			Addr     string
			Desc     string
		}{errors, name, password, birthday, tel, addr, desc}
	}

	utils.Render(w, "base.html", []string{"views/layouts/base.html", "views/user/create.html"}, context)
}

func UserModifyAction(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultManager.SessionStart(w, r)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/user/login/", http.StatusFound)
		return
	}

	id, _ := strconv.Atoi(r.FormValue("id"))
	context := struct {
		Errors map[string]string
		User   models.User
	}{nil, models.GetUserById(id)}

	if r.Method == http.MethodGet {

	} else if r.Method == http.MethodPost {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		name := r.FormValue("name")
		tel := r.FormValue("tel")
		addr := r.FormValue("addr")
		desc := r.FormValue("desc")
		errors := models.ValidateModifyUser(name, tel, addr, desc)
		if len(errors) == 0 {
			models.ModifyUser(id, name, tel, addr, desc)
			http.Redirect(w, r, "/users/", http.StatusFound)
			return
		}
		context.Errors = errors
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
	utils.Render(w, "base.html", []string{"views/layouts/base.html", "views/user/modify.html"}, context)
}

func UserDeleteAction(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultManager.SessionStart(w, r)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/user/login/", http.StatusFound)
		return
	}

	if r.Method == http.MethodGet {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err == nil {
			models.DeleteUser(id)
		} else {
			panic(err)
		}
		http.Redirect(w, r, "/users/", http.StatusFound)
	}
}

func init() {
	http.HandleFunc("/users/", UserAction)
	http.HandleFunc("/users/create/", UserCreateAction)
	http.HandleFunc("/user/login/", LoginAction)
	http.HandleFunc("/user/logout/", LogoutAction)
	http.HandleFunc("/users/modify/", UserModifyAction)
	http.HandleFunc("/users/delete/", UserDeleteAction)
}
