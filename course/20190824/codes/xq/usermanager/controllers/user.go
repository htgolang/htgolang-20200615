package controllers

import (

	"strings"
	"net/http"
	"html/template"
	"github.com/xlotz/usermanager/models"
	"github.com/xlotz/usermanager/session"
	//"github.com/xlotz/usermanager/controllers"



)

func UserAction(responseWriter http.ResponseWriter, request *http.Request)  {

		sessionObj := session.DefaultManager.SessionStart(responseWriter, request)
		if _, ok := sessionObj.Get("user"); !ok {
			http.Redirect(responseWriter, request, "/", http.StatusFound)
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




func init()  {
	http.HandleFunc("/", LoginAction)
	http.HandleFunc("/users/list/", UserAction)
	//http.HandleFunc("/users/create/", UserCreateAction)
	//http.HandleFunc("/users/modify/", UserModifyAction)
	//http.HandleFunc("/users/delete/", UserDeleteAction)
	//http.HandleFunc("/users/modifypass/", UserModifyPassAction)
	//http.HandleFunc("/users/register/", ReginsterUserAction)
	//http.HandleFunc("/users/query/", QueryAction)
}
