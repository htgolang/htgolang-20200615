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

		tpl.Execute(responseWriter, models.GetUserListFromDB(search))


}

func LoginAction(responseWriter http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		tpl := template.Must(template.New("login.html").ParseFiles("views/users/login.html"))
		tpl.Execute(responseWriter, nil)

	}else if request.Method == http.MethodPost {
		username := request.PostFormValue("username")
		password := request.PostFormValue("passwd")

		user, err := models.AutoUserFromDB(username, password)

		if err != nil {
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


	if request.Method == http.MethodPost{

		name := request.PostFormValue("username")
		brithday := request.PostFormValue("brithday")
		tel := request.PostFormValue("tel")
		addr := request.PostFormValue("addr")
		desc := request.PostFormValue("desc")
		pass := request.PostFormValue("passwd")
		sex := request.PostFormValue("sex")

		err := models.GetUserIdFromDB(name)
		if err == nil {
			models.InsertUserToDB(name, pass, sex, brithday, tel, addr, desc)
			http.Redirect(responseWriter, request, "/users/list/", http.StatusFound)
			return
		}

	}


	tpl := template.Must(template.New("createuser.html").ParseFiles("views/users/createuser.html"))
	tpl.Execute(responseWriter, nil)


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
		fmt.Println(id)

		if err == nil {
			user, error := models.GetUserFromDB(id)

			if error != nil {
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
		descs := request.PostFormValue("desc")
		//pass := request.PostFormValue("passwd")
		sex := request.PostFormValue("sex")

		fmt.Println(id, name, sex, brithday, tel, addr, descs)

		models.ModifyUserFromDB(id, name, sex, brithday, tel, addr, descs)
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
			user, err := models.GetUserFromDB(id)
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
		oldpass := request.PostFormValue("oldpass")
		newpass := request.PostFormValue("newpass")

		if err != nil {
			responseWriter.WriteHeader(http.StatusBadRequest)
		}
		user, err := models.AutoUserFromDB(name, oldpass)
		if err != nil {
			panic(err)

		}
		fmt.Println(user)

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


func init()  {
	http.HandleFunc("/", LoginAction)
	http.HandleFunc("/users/list/", UserAction)
	http.HandleFunc("/user/logout/", LogoutAction)
	http.HandleFunc("/users/create/", UserCreateAction)
	http.HandleFunc("/users/modify/", UserModifyAction)
	http.HandleFunc("/users/delete/", UserDeleteAction)
	http.HandleFunc("/users/modifypass/", UserModifyPassAction)

}
