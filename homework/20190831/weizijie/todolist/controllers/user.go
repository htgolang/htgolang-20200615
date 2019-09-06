package controllers

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"todolist/models"
	"todolist/session"
)

func Manager(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultManager.SessionStart(w, r)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}

	if r.Method == http.MethodGet {
		tpl := template.Must(template.New("manager.html").ParseFiles("views/user/manager.html"))
		tpl.Execute(w, nil)
	}
}

func UserAction(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultManager.SessionStart(w, r)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}

	q := strings.TrimSpace(r.FormValue("query"))
	tpl := template.Must(template.New("user.html").ParseFiles("views/user/user.html"))
	tpl.Execute(w, struct {
		Query string
		Users []models.User_str
	}{q, models.GetUsers(q)})
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tpl := template.Must(template.New("login.html").ParseFiles("views/user/login.html"))
		tpl.Execute(w, nil)

	} else if r.Method == http.MethodPost {
		name := r.PostFormValue("name")
		passwd := r.PostFormValue("passwd")

		passwd_md5 := fmt.Sprintf("%x", md5.Sum([]byte(passwd)))
		user, err := models.GetUserByName(name)

		if r.PostFormValue("Login") != "" {

			if err != nil || !user.ValidatePassword(passwd_md5) {
				tpl := template.Must(template.New("login.html").ParseFiles("views/user/login.html"))
				tpl.Execute(w, struct {
					Name  string
					Error string
				}{name, "用户名密码错误"})
			} else {
				sessionObj := session.DefaultManager.SessionStart(w, r)
				sessionObj.Set("user", user)
				models.Login_User = user.Name
				log.Printf("用户(%s)已登录", name)
				http.Redirect(w, r, "/user/manager/", http.StatusFound)
			}

		} else if r.PostFormValue("Register") != "" {
			models.Register = true
			sessionObj := session.DefaultManager.SessionStart(w, r)
			sessionObj.Set("user", user)
			http.Redirect(w, r, "/user/create/", http.StatusFound)
		}
	}
}

func UserCreateAction(w http.ResponseWriter, r *http.Request) {

	sessionObj := session.DefaultManager.SessionStart(w, r)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
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
		sex_str := r.PostFormValue("sex")
		sex, _ := strconv.ParseBool(sex_str)

		passwd_md5 := fmt.Sprintf("%x", md5.Sum([]byte(password)))

		errors := models.ValidateCreateUser(name, password, birthday, tel, addr, desc)
		if len(errors) == 0 {
			models.CreateUser(name, passwd_md5, birthday, tel, addr, desc, sex)
			if models.Register {
				http.Redirect(w, r, "/login/", http.StatusFound)
			} else {
				http.Redirect(w, r, "/user/", http.StatusFound)
			}
			models.Register = false
			return
		}

		context = struct {
			Errors   map[string]string
			Name     string
			Password string
			Sex      bool
			Birthday string
			Tel      string
			Addr     string
			Desc     string
		}{errors, name, passwd_md5, sex, birthday, tel, addr, desc}
	}
	tpl := template.Must(template.New("create.html").ParseFiles("views/user/create.html"))
	tpl.Execute(w, context)
}

func UserModifyAction(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultManager.SessionStart(w, r)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}

	if r.Method == http.MethodGet {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err == nil {
			user, err := models.GetUserById(id)
			if err != nil {
				log.Printf("用户(%s)获取失败", user)
				w.WriteHeader(http.StatusBadRequest)
			} else {
				tpl := template.Must(template.New("modify_user.html").ParseFiles("views/user/modify_user.html"))
				tpl.Execute(w, user)
			}
		} else {
			log.Printf("用户ID(%d)获取失败", id)
			w.WriteHeader(http.StatusBadRequest)
		}
	} else if r.Method == http.MethodPost {
		id, err := strconv.Atoi(r.PostFormValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		name := r.PostFormValue("name")
		desc := r.PostFormValue("desc")
		addr := r.PostFormValue("addr")
		tel := r.PostFormValue("tel")
		birthday := r.PostFormValue("birthday")
		sex_str := r.PostFormValue("sex")
		sex, _ := strconv.ParseBool(sex_str)

		models.ModifyUser(id, name, birthday, addr, tel, desc, sex)
		http.Redirect(w, r, "/user/", http.StatusFound)

	} else {
		log.Println("UserModify POST Faild")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func UserDeleteAction(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultManager.SessionStart(w, r)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}

	if id, err := strconv.Atoi(r.FormValue("id")); err == nil {
		models.DeleteUser(id)
	} else {
		log.Printf("用户ID(%s)获取失败：%v", id, err)
		fmt.Println(err)
	}
	http.Redirect(w, r, "/user/", http.StatusFound)
}

func LogoutAction(w http.ResponseWriter, r *http.Request) {
	session.DefaultManager.SessionDestory(w, r)
	http.Redirect(w, r, "/login/", http.StatusFound)
}

func ModifyPasswd(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tpl := template.Must(template.New("modify_passwd.html").ParseFiles("views/user/modify_passwd.html"))
		tpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		name := r.PostFormValue("name")
		passwd := r.PostFormValue("passwd")
		passwd_md5 := fmt.Sprintf("%x", md5.Sum([]byte(passwd)))
		models.ModifyPasswd(name, passwd_md5)
		log.Printf("用户(%s)密码已修改", name)
		http.Redirect(w, r, "/user", http.StatusFound)
	}
}

func init() {
	models.Init()
	http.HandleFunc("/", UserLogin)
	http.HandleFunc("/login/", UserLogin)
	http.HandleFunc("/user/", UserAction)
	http.HandleFunc("/user/logout/", LogoutAction)
	http.HandleFunc("/user/create/", UserCreateAction)
	http.HandleFunc("/user/modify/", UserModifyAction)
	http.HandleFunc("/user/delete/", UserDeleteAction)
	http.HandleFunc("/user/manager/", Manager)
	http.HandleFunc("/passwd/modify/", ModifyPasswd)
}
