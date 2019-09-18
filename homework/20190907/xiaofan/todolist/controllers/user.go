package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"todolist/models"
	"todolist/session"
)

// 用户列表
func UserAction(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultSessionManager.SessionStart(w, r)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}

	query := strings.TrimSpace(r.PostFormValue("query"))
	users := models.GetUsers(query)

	context := struct {
		Q     string
		Users []models.User
	}{query, users}

	tpl := template.Must(template.New("user.html").ParseFiles("views/users/user.html"))
	_ = tpl.Execute(w, context)

}

// 创建用户
func CreateUserAction(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultSessionManager.SessionStart(w, r)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}

	var context interface{}
	if r.Method == http.MethodPost {
		name := r.PostFormValue("name")
		password := r.PostFormValue("password")
		bir := r.PostFormValue("bir")
		tel := r.PostFormValue("tel")
		addr := r.PostFormValue("addr")
		desc := r.PostFormValue("desc")
		sex := r.PostFormValue("sex")

		errors := models.ValidateCreateUser(name, password, bir, tel)
		if len(errors) == 0 {
			models.CreateUser(name, password, sex, bir, tel, addr, desc)
			http.Redirect(w, r, "/users/", http.StatusFound)
		}
		context = struct {
			Errors   map[string]string
			Name     string
			Password string
			Birthday string
			Tel      string
		}{errors, name, password, bir, tel}

	}

	tpl := template.Must(template.New("create_user.html").ParseFiles("views/users/create_user.html"))
	_ = tpl.Execute(w, context)
}

// 修改用户
func ModifyUserAction(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultSessionManager.SessionStart(w, r)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}

	var context struct {
		User     models.User
		Errors   map[string]string
		Name     string
		Birthday string
		Tel      string
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	user, err := models.GetUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if r.Method == http.MethodPost {
		id, err := strconv.Atoi(r.PostFormValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		name := r.PostFormValue("name")
		bir := r.PostFormValue("bir")
		tel := r.PostFormValue("tel")
		addr := r.PostFormValue("addr")
		desc := r.PostFormValue("desc")

		errors := models.ValidateModifyUser(name, bir, tel)
		if len(errors) == 0 {
			models.ModifyUser(id, name, bir, tel, addr, desc)
			http.Redirect(w, r, "/users/", http.StatusFound)
		}

		context.Errors = errors
		context.Name = name
		context.Birthday = bir
		context.Tel = tel
	}

	context.User = user
	tpl := template.Must(template.New("modify_user.html").ParseFiles("views/users/modify_user.html"))
	err = tpl.Execute(w, context)
	fmt.Println(err)
	return
}

// 删除用户
func DeleteUserAction(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultSessionManager.SessionStart(w, r)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}

	if r.Method == http.MethodGet {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		models.DeleteUser(id)
		http.Redirect(w, r, "/users/", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

// 用户登录
func LoginAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tpl := template.Must(template.New("login.html").ParseFiles("views/login.html"))
		_ = tpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		if ok := models.CheckUser(username, password); ok == false {
			resp := struct {
				Username string
				Resp     string
			}{username, "用户名或密码不正确，请重新输入"}

			tpl := template.Must(template.New("login.html").ParseFiles("views/login.html"))
			_ = tpl.Execute(w, resp)

		} else {
			// 根据w,r生成一个session，并且设置过期时间
			sessionObj := session.DefaultSessionManager.SessionStart(w, r)
			sessionObj.Set("user", username)
			http.Redirect(w, r, "/tasks/", http.StatusFound)
		}
	}
}

// 修改个人密码
func ModifyPasswordAction(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultSessionManager.SessionStart(w, r)
	var username interface{}
	if user, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	} else {
		username = user
	}

	if r.Method == http.MethodGet {
		tpl := template.Must(template.New("modify_password.html").ParseFiles("views/users/modify_password.html"))
		_ = tpl.Execute(w, username)
	} else if r.Method == http.MethodPost {
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		models.ModifyPassword(username, password)
		http.Redirect(w, r, "/login/", http.StatusFound)

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

// 用户注销
func LogoutAction(w http.ResponseWriter, r *http.Request) {
	session.DefaultSessionManager.SessionDestory(w, r)
	http.Redirect(w, r, "/login/", http.StatusFound)
}

func init() {
	http.HandleFunc("/users/", UserAction)
	http.HandleFunc("/login/", LoginAction)
	http.HandleFunc("/logout/", LogoutAction)
	http.HandleFunc("/users/create/", CreateUserAction)
	http.HandleFunc("/users/modify/", ModifyUserAction)
	http.HandleFunc("/users/delete/", DeleteUserAction)
	http.HandleFunc("/users/password/", ModifyPasswordAction)
}
