package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"github.com/xlotz/todolist/models"
	"github.com/xlotz/todolist/session"
)

func UserAction(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultManager.SessionStart(w, r)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/user/login/", http.StatusFound)
		return
	}

	q := strings.TrimSpace(r.FormValue("q"))

	tpl := template.Must(template.New("user.html").ParseFiles("views/user/user.html"))

	tpl.Execute(w, struct {
		Query string
		Users []models.User
	}{q, models.GetUsers(q)})
}

func LoginAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tpl := template.Must(template.New("login.html").ParseFiles("views/user/login.html"))
		tpl.Execute(w, nil)
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

	tpl := template.Must(template.New("create.html").ParseFiles("views/user/create.html"))
	tpl.Execute(w, context)
}

// 修改用户
func UserModifyAction(responseWriter http.ResponseWriter, request *http.Request)  {


	sessionObj := session.DefaultManager.SessionStart(responseWriter, request)
	if _, ok := sessionObj.Get("user"); !ok {

		http.Redirect(responseWriter, request, "/", http.StatusFound)
		return
	}

	if request.Method == http.MethodGet {

		id, err := strconv.Atoi(request.FormValue("id"))

		if err == nil {
			user, error := models.GetUserId(id)

			if error != nil {
				responseWriter.WriteHeader(http.StatusBadRequest)
			}else {

				tpl := template.Must(template.New("modifyuser.html").ParseFiles("views/user/modifyuser.html"))
				tpl.Execute(responseWriter, user)
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

		models.ModifyUserFromDB(id, name, tel, addr, descs)
		http.Redirect(responseWriter, request, "/users/list/", http.StatusFound)



	}else {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
	}


}
// 修改用户密码
func UserModifyPassAction(responseWriter http.ResponseWriter, request *http.Request)  {

	sessionObj := session.DefaultManager.SessionStart(responseWriter, request)
	if _, ok := sessionObj.Get("user"); !ok {

		http.Redirect(responseWriter, request, "/", http.StatusFound)
		return
	}

	if request.Method == http.MethodGet {

		id, err := strconv.Atoi(request.FormValue("id"))

		if err == nil {
			user, err := models.GetUserId(id)
			if err != nil {
				responseWriter.WriteHeader(http.StatusBadRequest)
			}else {
				tpl := template.Must(template.New("modifypass.html").ParseFiles("views/user/modifypass.html"))
				tpl.Execute(responseWriter, user)
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

		//name := request.PostFormValue("username")
		oldpass := request.PostFormValue("oldpass")
		newpass := request.PostFormValue("newpass")

		if err != nil {
			responseWriter.WriteHeader(http.StatusBadRequest)
		}

		if oldpass != newpass {
			panic("2次密码不相同")
		}


		models.ModifyPassFromDB(id, newpass)
		http.Redirect(responseWriter, request, "/users/list/", http.StatusFound)




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
	http.HandleFunc("/user/login/", LoginAction)
	http.HandleFunc("/user/logout/", LogoutAction)
	http.HandleFunc("/users/modify/", UserModifyAction)
	http.HandleFunc("/users/delete/", UserDeleteAction)
	http.HandleFunc("/users/modifypass/", UserModifyPassAction)
}
