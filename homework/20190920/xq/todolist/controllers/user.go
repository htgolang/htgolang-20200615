package controllers

import (
	"fmt"
	"net/http"

	"strconv"
	"strings"
	"github.com/xlotz/todolist/models"
	"github.com/xlotz/todolist/session"
	"github.com/xlotz/todolist/utils"
)

func UserAction(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultManager.SessionStart(w, r)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/user/login/", http.StatusFound)
		return
	}

	q := strings.TrimSpace(r.FormValue("q"))

	utils.Render(w, []string{"views/layouts/base.html", "views/user/user.html"}, "base.html", struct {
		Query string
		Users []models.User
	}{q, models.GetUsers(q)})

}

func LoginAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		utils.Render(w, []string{"views/user/login.html"}, "login.html", nil)
	} else if r.Method == http.MethodPost {
		name := r.PostFormValue("name")
		password := r.PostFormValue("password")

		user, err := models.GetUserByName(name)
		if err != nil || !user.ValidatePassword(password) {
			utils.Render(w, []string{"views/user/login.html"}, "login.html", struct {
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

	utils.Render(w, []string{"views/layouts/base.html", "views/user/create.html"}, "base.html", context)
}

func ModifyUserAction(responseWriter http.ResponseWriter, request *http.Request) {
	sessionObj := session.DefaultManager.SessionStart(responseWriter, request)
	if _, ok := sessionObj.Get("user"); !ok {

		http.Redirect(responseWriter, request, "/", http.StatusFound)
		return
	}

	if request.Method == http.MethodGet {

		id, err := strconv.Atoi(request.FormValue("id"))

		if err == nil {
			user, error := models.GetUserById(id)
			fmt.Println(user)

			if error != nil {
				responseWriter.WriteHeader(http.StatusBadRequest)
			}else {
				utils.Render(responseWriter, []string{"views/layouts/base.html", "views/user/modifyuser.html"}, "base.html", user)
				//tpl := template.Must(template.New("modifyuser.html").ParseFiles("views/user/modifyuser.html"))
				//tpl.Execute(responseWriter, user)
			}
		} else {
			responseWriter.WriteHeader(http.StatusBadRequest)
		}


	}else if request.Method == http.MethodPost{


		id, err := strconv.Atoi(request.PostFormValue("id"))

		if err != nil {
			responseWriter.WriteHeader(http.StatusBadRequest)
			return
		}

		name := request.PostFormValue("username")
		//birthday := request.PostFormValue("brithday")
		tel := request.PostFormValue("tel")
		addr := request.PostFormValue("addr")
		descs := request.PostFormValue("desc")
		//pass := request.PostFormValue("passwd")
		sex := request.PostFormValue("sex")


		fmt.Println(id, name, sex, tel, addr, descs)

		models.ModifyUser(id, name, tel, addr, descs)
		http.Redirect(responseWriter, request, "/users/", http.StatusFound)



	}else {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// 删除用户
func UserDeleteAction(responseWriter http.ResponseWriter, request *http.Request)  {

	sessionObj := session.DefaultManager.SessionStart(responseWriter, request)
	if _, ok := sessionObj.Get("user"); !ok {

		http.Redirect(responseWriter, request, "/", http.StatusFound)
		return
	}

	if id, err := strconv.Atoi(request.FormValue("id")); err == nil  {

		models.DeleteUserFromDB(id)

	}else {
		fmt.Println(err)
	}
	http.Redirect(responseWriter, request, "/users/list/", http.StatusFound)
}

func init() {
	http.HandleFunc("/users/", UserAction)
	http.HandleFunc("/users/create/", UserCreateAction)
	http.HandleFunc("/users/modify/", ModifyUserAction)
	http.HandleFunc("/users/delete/", UserDeleteAction)
	http.HandleFunc("/user/login/", LoginAction)
	http.HandleFunc("/user/logout/", LogoutAction)

}
