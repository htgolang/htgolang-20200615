package controllers

import (
"fmt"
//"go/ast"
"html/template"
"net/http"
"github.com/xlotz/todolist/models"
"strconv"
)

// 用户登陆
func UserLoginAction(responseWriter http.ResponseWriter, request *http.Request)  {

	if request.Method == http.MethodGet {
		tpl := template.Must(template.New("login.html").ParseFiles("views/users/login.html"))
		tpl.Execute(responseWriter, nil)
	}else if request.Method == http.MethodPost{
		name := request.PostFormValue("username")
		pass := request.PostFormValue("passwd")
		if models.AutoUsers(name, pass) {

			http.Redirect(responseWriter, request, "/users/list/", http.StatusFound)

		}else {
			http.Redirect(responseWriter, request, "/", http.StatusFound)
		}


	}

}
// 用户登出
func UserLogoutAction(responseWriter http.ResponseWriter, request *http.Request)  {

		//tpl := template.Must(template.New("login.html").ParseFiles("views/users/login.html"))
		//tpl.Execute(responseWriter, nil)
		http.Redirect(responseWriter, request, "/", http.StatusFound)
}
// 用户列表
func UserAction(responseWriter http.ResponseWriter, request *http.Request)  {

	tpl := template.Must(template.New("userlist.html").ParseFiles("views/users/userlist.html"))

	tpl.Execute(responseWriter, models.GetUsers())
}
// 创建用户
func UserCreateAction(responseWriter http.ResponseWriter, request *http.Request)  {

	if request.Method == http.MethodGet {
		tpl := template.Must(template.New("createuser.html").ParseFiles("views/users/createuser.html"))
		tpl.Execute(responseWriter, nil)

	}else if request.Method == http.MethodPost{

		name := request.PostFormValue("username")
		brithday := request.PostFormValue("brithday")
		tel := request.PostFormValue("tel")
		addr := request.PostFormValue("addr")
		desc := request.PostFormValue("desc")
		pass := request.PostFormValue("passwd")

		models.CreateUsers(name, brithday, tel, addr, desc, pass)
		http.Redirect(responseWriter, request, "/users/list/", http.StatusFound)

	}else {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
	}

}
// 修改用户
func UserModifyAction(responseWriter http.ResponseWriter, request *http.Request)  {

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
		}

		name := request.PostFormValue("username")
		brithday := request.PostFormValue("brithday")
		tel := request.PostFormValue("tel")
		addr := request.PostFormValue("addr")
		desc := request.PostFormValue("desc")
		pass := request.PostFormValue("passwd")
		//progress, err := strconv.Atoi(request.PostFormValue("progress"))
		if err != nil {
			responseWriter.WriteHeader(http.StatusBadRequest)
		}
		models.ModifyUsers(id, name, brithday, tel, addr, desc, pass)
		http.Redirect(responseWriter, request, "/users/list/", http.StatusFound)

	}else {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
	}

}
// 修改用户密码
func UserModifyPassAction(responseWriter http.ResponseWriter, request *http.Request)  {

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
	if id, err := strconv.Atoi(request.FormValue("id")); err == nil  {

		models.DeleteUsers(id)

	}else {
		fmt.Println(err)
	}
	http.Redirect(responseWriter, request, "/users/list/", http.StatusFound)
}
// 注册用户
func ReginsterUserAction(responseWriter http.ResponseWriter, request *http.Request){
	if request.Method == http.MethodGet {
		tpl := template.Must(template.New("registeruser.html").ParseFiles("views/users/registeruser.html"))
		tpl.Execute(responseWriter, nil)

	}else if request.Method == http.MethodPost{

		name := request.PostFormValue("username")
		brithday := request.PostFormValue("brithday")
		tel := request.PostFormValue("tel")
		addr := request.PostFormValue("addr")
		desc := request.PostFormValue("desc")
		pass := request.PostFormValue("passwd")

		models.ReginsterUsers(name, brithday, tel, addr, desc, pass)
		http.Redirect(responseWriter, request, "/", http.StatusFound)

	}else {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func init()  {
	http.HandleFunc("/", UserLoginAction)
	http.HandleFunc("/users/list/", UserAction)
	http.HandleFunc("/users/create/", UserCreateAction)
	http.HandleFunc("/users/modify/", UserModifyAction)
	http.HandleFunc("/users/delete/", UserDeleteAction)
	http.HandleFunc("/users/modifypass/", UserModifyPassAction)
	http.HandleFunc("/users/register/", ReginsterUserAction)
}

