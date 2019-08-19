package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"todolist/models"
)

func UserAction(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.New("user.html").ParseFiles("views/user.html"))
	tpl.Execute(w, models.GetUsers())
}

func CreateUserAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tpl := template.Must(template.New("create_user.html").ParseFiles("views/create_user.html"))
		tpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {

		name := r.PostFormValue("name")
		password := fmt.Sprintf("%x", []byte(r.PostFormValue("password")))
		bir := r.PostFormValue("birthday")
		tel := r.PostFormValue("tel")
		addr := r.PostFormValue("addr")
		desc := r.PostFormValue("desc")

		models.CreateUser(name, password, bir, tel, addr, desc)
		http.Redirect(w, r, "/users/", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func ModifyUserAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err == nil {
			user, err := models.GetUserById(id)
			if err != nil {
				w.WriteHeader(http.StatusFound)
			} else {
				tpl := template.Must(template.New("modify_user.html").ParseFiles("views/modify_user.html"))
				tpl.Execute(w, user)
			}
		} else {
			w.WriteHeader(http.StatusFound)
		}

	} else if r.Method == http.MethodPost {
		id, err := strconv.Atoi(r.PostFormValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		name := r.PostFormValue("name")
		bir := r.PostFormValue("birthday")
		tel := r.PostFormValue("tel")
		addr := r.PostFormValue("addr")
		desc := r.PostFormValue("desc")

		models.ModifyUser(id, name, bir, tel, addr, desc)
		http.Redirect(w, r, "/users/", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func DeleteUserAction(w http.ResponseWriter, r *http.Request) {
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

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tpl := template.Must(template.New("login.html").ParseFiles("views/login.html"))
		tpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		username := r.PostFormValue("username")
		password := fmt.Sprintf("%x", []byte(r.PostFormValue("password")))

		if ok := models.CheckUser(username, password); ok == false {
			w.WriteHeader(http.StatusForbidden)
		} else {
			models.SetCookie(username, r.RemoteAddr, r.UserAgent())
			http.Redirect(w, r, "/tasks/", http.StatusFound)

		}
	}
}

func ModifyPasswordAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// 获取cookie
		cookie := models.GetCookie(r.RemoteAddr, r.UserAgent())
		if cookie == "" {
			http.Redirect(w, r, "/login/", http.StatusFound)
		} else {
			tpl := template.Must(template.New("modify_password.html").ParseFiles("views/modify_password.html"))
			tpl.Execute(w, cookie)
		}
	} else if r.Method == http.MethodPost {
		fmt.Println("2")
		username := r.PostFormValue("username")
		password := fmt.Sprintf("%x", []byte(r.PostFormValue("password")))

		models.ModifyPassword(username, password)
		http.Redirect(w, r, "/login/", http.StatusFound)

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func init() {
	http.HandleFunc("/login/", Login)
	http.HandleFunc("/users/", UserAction)
	http.HandleFunc("/users/create/", CreateUserAction)
	http.HandleFunc("/users/modify/", ModifyUserAction)
	http.HandleFunc("/users/delete/", DeleteUserAction)
	http.HandleFunc("/users/password/", ModifyPasswordAction)
}
