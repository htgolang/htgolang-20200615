package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"net/http"
	"html/template"
	"github.com/xlotz/usermanager/models"
	"github.com/xlotz/usermanager/session"




)

func UserAction(responseWriter http.ResponseWriter, request *http.Request)  {

		sessionObj := session.DefaultManager.SessionStart(responseWriter, request)
		if _, ok := sessionObj.Get("user"); !ok {
			http.Redirect(responseWriter, request, "/", http.StatusFound)
			return
		}
		tpl := template.Must(template.New("userlist.html").ParseFiles("views/users/userlist.html"))


		search := strings.TrimSpace(request.FormValue("search"))

		if search == "" {

			tpl.Execute(responseWriter, models.GetUsers())
		}else {
			//fmt.Println(search)
			tpl.Execute(responseWriter, models.GetUsers(search))
		}

		//// 优化

		//tpl.Execute(responseWriter,
		//	struct {
		//		Query string
		//		Users []models.User
		//	}{search, models.GetUsers(search)})


}

func LoginAction(responseWriter http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		tpl := template.Must(template.New("login.html").ParseFiles("views/users/login.html"))
		tpl.Execute(responseWriter, nil)

	}else if request.Method == http.MethodPost {
		username := request.PostFormValue("username")
		password := request.PostFormValue("passwd")

		user, err := models.AutoUsers(username)
		if err != nil || !user.ValidatePassword(password) {
			tpl := template.Must(template.New("login.html").ParseFiles("views/users/login.html"))
			tpl.Execute(responseWriter, struct {
				Name string
				Error string
			}{username, "用户名或密码错误"})

		}else {

			sessionObj := session.DefaultManager.SessionStart(responseWriter, request)
			// 登录成功
			sessionObj.Set("user", user)
			http.Redirect(responseWriter, request, "/users/list/", http.StatusFound)
		}
	}
}

func LogoutAction(responseWriter http.ResponseWriter, request *http.Request)  {
	session.DefaultManager.SessionDestory(responseWriter, request)
	http.Redirect(responseWriter, request, "/", http.StatusFound)
}

// 创建用户
func UserCreateAction(responseWriter http.ResponseWriter, request *http.Request)  {

	sessionObj := session.DefaultManager.SessionStart(responseWriter, request)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(responseWriter, request, "/users/login/", http.StatusFound)
		return
	}

	var context interface{}

	if request.Method == http.MethodPost{

		name := request.PostFormValue("username")
		brithday := request.PostFormValue("brithday")
		tel := request.PostFormValue("tel")
		addr := request.PostFormValue("addr")
		desc := request.PostFormValue("desc")
		pass := request.PostFormValue("passwd")

		errors := models.ValidateCreateUser(name, brithday, tel, addr, desc, pass)
		if len(errors) == 0 {
			models.CreateUsers(name, brithday, tel, addr, desc, pass)
			http.Redirect(responseWriter, request, "/users/list/", http.StatusFound)
			return
		}

		context = struct {
			Errors map[string]string
			Name string
			Brithday string
			Tel string
			Addr string
			Desc string
			Pass string

		}{errors, name, brithday, tel, addr, desc, pass}

		fmt.Println(context)
	}


	tpl := template.Must(template.New("createuser.html").ParseFiles("views/users/createuser.html"))
	tpl.Execute(responseWriter, context)


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
			user, err := models.GetUserById(id)
			if err != nil {
				responseWriter.WriteHeader(http.StatusBadRequest)
			}else {

				tpl := template.Must(template.New("modifyuser.html").ParseFiles("views/users/modifyuser.html"))
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
		brithday := request.PostFormValue("brithday")
		tel := request.PostFormValue("tel")
		addr := request.PostFormValue("addr")
		desc := request.PostFormValue("desc")
		pass := request.PostFormValue("passwd")

		models.ModifyUsers(id, name, brithday, tel, addr, desc, pass)
		http.Redirect(responseWriter, request, "/users/list/", http.StatusFound)
		//errors := models.ValidateCreateUser(name, brithday, tel, addr, desc, pass)
		//if len(errors) == 0 {
		//	models.ModifyUsers(id, name, brithday, tel, addr, desc, pass)
		//	http.Redirect(responseWriter, request, "/users/list/", http.StatusFound)
		//	return
		//}


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
			user, err := models.GetUserById(id)
			if err != nil {
				responseWriter.WriteHeader(http.StatusBadRequest)
			}else {
				tpl := template.Must(template.New("modifypass.html").ParseFiles("views/users/modifypass.html"))
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
		brithday := request.PostFormValue("brithday")
		tel := request.PostFormValue("tel")
		addr := request.PostFormValue("addr")
		desc := request.PostFormValue("desc")
		oldpass := request.PostFormValue("oldpass")
		newpass := request.PostFormValue("newpass")

		if err != nil {
			responseWriter.WriteHeader(http.StatusBadRequest)
		}
		models.ModifyUserPass(id, name, brithday, tel, addr, desc, oldpass, newpass)
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

		models.DeleteUsers(id)

	}else {
		fmt.Println(err)
	}
	http.Redirect(responseWriter, request, "/users/list/", http.StatusFound)
}


func init()  {
	http.HandleFunc("/", LoginAction)
	http.HandleFunc("/users/list/", UserAction)
	http.HandleFunc("/user/logout/", LogoutAction)
	http.HandleFunc("/users/create/", UserCreateAction)
	http.HandleFunc("/users/modify/", UserModifyAction)
	http.HandleFunc("/users/delete/", UserDeleteAction)
	http.HandleFunc("/users/modifypass/", UserModifyPassAction)

}
