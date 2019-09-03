package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"user/models"
	"user/session"
)

func UserAction(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultManager.SessionStart(w, r) // 登陆session
	fmt.Println("LoginList:", sessionObj)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/users/login/", http.StatusFound)
		return
	}
	q := strings.TrimSpace(r.FormValue("q"))
	fmt.Println(q)

	tpl := template.Must(template.New("user.html").ParseFiles("views/user/user.html"))
	tpl.Execute(w, struct {
		Q     string
		Users []models.UserDescribe
	}{q, models.GetUsers(q)})
}
func UserActionLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tpl := template.Must(template.New("login.html").ParseFiles("views/user/login.html"))
		tpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		name := r.PostFormValue("name")
		password := r.PostFormValue("password")

		user, err := models.GetUserByName(name)
		if err != nil || !user.ValidatePassword(password) {
			// 登陆失败
			tpl := template.Must(template.New("login.html").ParseFiles("views/user/login.html"))
			tpl.Execute(w, struct {
				Name  string
				Error string
			}{name, "用户名或密码错误!"})
		} else {
			//登陆成功
			sessionObj := session.DefaultManager.SessionStart(w, r)
			sessionObj.Set("user", user)
			fmt.Println("LoginOK:", sessionObj)

			http.Redirect(w, r, "/users/list/", http.StatusFound)
		}
	}
}
func UserActionLoout(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultManager.SessionStart(w, r) // 登陆session
	fmt.Println("LoginList:", sessionObj)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/users/login/", http.StatusFound)
		return
	}
	session.DefaultManager.SessionDestory(w, r)
	http.Redirect(w, r, "/users/login/", http.StatusFound)
}
func UserActionCreate(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultManager.SessionStart(w, r) // 登陆session
	fmt.Println("LoginList:", sessionObj)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/users/login/", http.StatusFound)
		return
	}
	var context interface{}
	if r.Method == http.MethodPost {
		name := r.PostFormValue("name")
		password := r.PostFormValue("password")
		birthday := r.PostFormValue("birthday")
		sex := r.PostFormValue("sex")
		tel := r.PostFormValue("tel")
		addr := r.PostFormValue("addr")
		desc := r.PostFormValue("desc")
		errors := models.ValidateCreateUser(name, password, birthday, sex, tel, addr, desc)
		if len(errors) == 0 {
			models.CreateUser(name, password, birthday, sex, tel, addr, desc)
			http.Redirect(w, r, "/users/list/", http.StatusFound)
			return
		}
		context = struct {
			Error    map[string]string
			Name     string
			Password string
			Birthday string
			Tel      string
			Addr     string
			Desc     string
		}{errors, name, password, birthday, tel, addr, desc}
	}

	tpl := template.Must(template.New("createuser.html").ParseFiles("views/user/createuser.html"))
	tpl.Execute(w, context)
}

func UserActionModify(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultManager.SessionStart(w, r) // 登陆session
	fmt.Println("LoginList:", sessionObj)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/users/login/", http.StatusFound)
		return
	}

	if r.Method == http.MethodPost {
		id, err := strconv.Atoi(r.PostFormValue("id")) // PostFormValue
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		fmt.Println("PostFormValue-ID:", id)
		name := r.PostFormValue("name")
		password := r.PostFormValue("password")
		birthday := r.PostFormValue("birthday")
		sex := r.PostFormValue("sex")
		tel := r.PostFormValue("tel")
		addr := r.PostFormValue("addr")
		desc := r.PostFormValue("desc")
		fmt.Println("MethodPost:", id, password)
		models.ModifyUser(id, name, password, birthday, sex, tel, addr, desc)
		http.Redirect(w, r, "/users/list/", http.StatusFound)
	}
	id, err := strconv.Atoi(strings.TrimSpace(r.FormValue("id"))) // FormValue
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println("PostFormValue-ID:", id)
	user, err := models.GetUserByID(id)
	fmt.Println("GetUserByID：", err, user)
	if err != nil {
		http.Redirect(w, r, "/users/list/", http.StatusFound)
	} else {
		tpl := template.Must(template.New("modifyuser.html").ParseFiles("views/user/modifyuser.html"))
		tpl.Execute(w, user)
	}
}

func UserActionDelete(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultManager.SessionStart(w, r) // 登陆session
	fmt.Println("LoginList:", sessionObj)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/users/login/", http.StatusFound)
		return
	}
	id, err := strconv.Atoi(r.FormValue("id"))
	if err == nil {
		fmt.Println("删除的ID", id)
		models.DeleteUser(id)
	}
	http.Redirect(w, r, "/users/list/", http.StatusFound)
}
func init() {
	http.HandleFunc("/users/list/", UserAction)
	http.HandleFunc("/users/login/", UserActionLogin)
	http.HandleFunc("/users/logout/", UserActionLoout)
	http.HandleFunc("/users/create/", UserActionCreate)
	http.HandleFunc("/users/modify/", UserActionModify)
	http.HandleFunc("/users/delete/", UserActionDelete)
}
