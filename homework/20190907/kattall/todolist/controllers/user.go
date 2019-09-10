package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"todolist/modules"
	"todolist/session"
)

func init() {
	http.HandleFunc("/", LoginAction)
	http.HandleFunc("/users/", UsersAction)
	http.HandleFunc("/users/logout/", LogOutAction)
	http.HandleFunc("/users/create/", UserCreateAction)
	http.HandleFunc("/users/modify/", UserModifyAction)
	http.HandleFunc("/users/delete/", UserDeleteAction)
}

func UsersAction(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.New("users.html").ParseFiles("views/user/users.html"))
	q := r.FormValue("query")
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
		name := strings.TrimSpace(r.PostFormValue("name"))
		password := strings.TrimSpace(r.PostFormValue("password"))
		fmt.Println("login, name, password:", name, password)
		user, b := modules.GetUserByName(name)
		if !b || !user.VaildatePassword(password) {
			tpl := template.Must(template.New("login.html").ParseFiles("views/user/login.html"))
			tpl.Execute(w, struct {
				Name  string
				Error string
			}{name, "用户名或密码不正确"})
		} else {
			sessionObj := session.DefaultManager.SessionStart(w, r)
			sessionObj.SET("user", user)
			http.Redirect(w, r, "/tasks/", http.StatusFound)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func LogOutAction(w http.ResponseWriter, r *http.Request) {
	session.DefaultManager.SessionDestory(w, r)
	http.Redirect(w, r, "/", http.StatusFound)
}

func UserCreateAction(w http.ResponseWriter, r *http.Request) {
	var context struct {
		Errors   map[string]string
		Name     string
		Password string
		Birthday string
		Sex      string
		Tel      string
		Desc     string
		Addr     string
	}
	sessionObj := session.DefaultManager.SessionStart(w, r)
	if _, ok := sessionObj.GET("user"); !ok {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	if r.Method == http.MethodPost {
		var sexx bool
		name := strings.TrimSpace(r.PostFormValue("name"))
		password := strings.TrimSpace(r.PostFormValue("password"))
		password2 := strings.TrimSpace(r.PostFormValue("password2"))
		birthday := strings.TrimSpace(r.PostFormValue("birthday"))
		sex, _ := strconv.Atoi(strings.TrimSpace(r.PostFormValue("sex")))
		if sex == 0 {
			sexx = true
		} else if sex == 1 {
			sexx = false
		}
		tel := strings.TrimSpace(r.PostFormValue("tel"))
		desc := strings.TrimSpace(r.PostFormValue("desc"))
		addr := strings.TrimSpace(r.PostFormValue("addr"))
		errors := modules.ValidateCreateUser(name, password, password2, birthday, tel, desc, addr)
		if len(errors) == 0 {
			modules.CreateUser(name, password, birthday, sexx, tel, desc, addr)
			http.Redirect(w, r, "/users/", http.StatusFound)
			return
		}
		context.Errors = errors
		context.Name = name
		context.Password = password
		context.Birthday = birthday
		context.Addr = addr
		context.Desc = desc
		context.Tel = tel
	}
	tpl := template.Must(template.New("create_user.html").ParseFiles("views/user/create_user.html"))
	tpl.Execute(w, context)
}

func UserModifyAction(w http.ResponseWriter, r *http.Request) {
	var context struct {
		User     modules.User
		Errors   map[string]string
		Name     string
		Birthday string
	}
	sessionObj := session.DefaultManager.SessionStart(w, r)
	if _, ok := sessionObj.GET("user"); !ok {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	id, _ := strconv.Atoi(r.FormValue("id"))
	user, _ := modules.GetUserById(id)
	if r.Method == http.MethodPost {
		var sexx bool
		id, err := strconv.Atoi(r.PostFormValue("id"))
		if err != nil {
			panic(err)
		}
		name := r.PostFormValue("name")
		birthday := r.PostFormValue("birthday")
		sex, _ := strconv.Atoi(r.PostFormValue("sex"))
		if sex == 0 {
			sexx = true
		} else if sex == 1 {
			sexx = false
		}
		tel := r.PostFormValue("tel")
		desc := r.PostFormValue("desc")
		addr := r.PostFormValue("addr")
		errors := modules.ValidateModifyUser(name, birthday)
		fmt.Println(name, birthday, sex, tel, desc, addr)
		if len(errors) == 0 {
			fmt.Println("modify user:", id, name, birthday,tel, desc, addr, sexx)
			modules.ModifyUser(id, name, birthday, tel, desc, addr, sexx)
			http.Redirect(w, r, "/users/", http.StatusFound)
			return
		}
		context.Errors = errors
		context.Name = name
		context.Birthday = birthday
	}
	tpl := template.Must(template.New("modify_user.html").ParseFiles("views/user/modify_user.html"))
	context.User = user
	tpl.Execute(w, context)
}

func UserDeleteAction(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	modules.DeleteUser(id)
	http.Redirect(w, r, "/users/", http.StatusFound)
}
