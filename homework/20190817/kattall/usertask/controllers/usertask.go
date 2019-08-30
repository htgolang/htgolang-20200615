package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"
	"usertask/modules"
)

func init(){
	http.HandleFunc("/", UserLoginAction)
	http.HandleFunc("/user/listUser/", ListUserAction)
	http.HandleFunc("/user/createUser/", CreateUserAction)
	http.HandleFunc("/user/modifyUser/", ModifyUser)
	http.HandleFunc("/user/deleteUser/", DeleteUser)
	http.HandleFunc("/user/modifyUserPassword/", ModifyUserPassword)

	http.HandleFunc("/user/userTaskList/", UserTaskList)
	http.HandleFunc("/user/userTaskCreate/", UserTaskCreate)
}

// 用户登陆页面  只有get方法
func UserLoginAction(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
		tpl := template.Must(template.New("userLogin.html").ParseFiles("views/userLogin.html"))
		tpl.Execute(w, nil)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// 用户列表 只有get方法
func ListUserAction(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
		tpl := template.Must(template.New("userList.html").ParseFiles("views/userList.html"))
		fmt.Println("userList:", modules.ListUser())
		tpl.Execute(w, modules.ListUser())
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// 创建用户
// get  返回createUser.html   点击创建后, post请求
// post 添加用户(添加用户判断用户名是否存在), 然后返回任务列表。默认列表为空
func CreateUserAction(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
		tpl := template.Must(template.New("userCreate.html").ParseFiles("views/userCreate.html"))
		tpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		username := r.PostFormValue("username")
		password := fmt.Sprintf("%x", []byte(r.PostFormValue("password")))
		birthday, err := time.Parse("2006-01-02", r.PostFormValue("birthday"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		addr := r.PostFormValue("addr")
		desc := r.PostFormValue("desc")
		fmt.Printf("username: %s, password: %s, birthday: %s, addr: %s, desc: %s\n", username, password, birthday, addr, desc)
		if err := modules.CreateUser(username, password, birthday, addr, desc); err != nil {
			w.WriteHeader(http.StatusForbidden)
		} else {
			http.Redirect(w, r, "/user/listUser/", http.StatusFound)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func UserTaskList(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
		tpl := template.Must(template.New("userTaskList.html").ParseFiles("views/userTaskList.html"))
		tpl.Execute(w, nil)
	} else {
		username := r.PostFormValue("username")
		password := fmt.Sprintf("%x", []byte(r.PostFormValue("password")))
		fmt.Println(modules.CheckUserPwd(username, password))
		if user, err := modules.CheckUserPwd(username, password); err != nil {
			w.WriteHeader(http.StatusForbidden)
		} else {
			tpl := template.Must(template.New("userTaskList.html").ParseFiles("views/userTaskList.html"))
			tpl.Execute(w, user)
		}
	}
}

// *** 这里创建task后 302跳转还有点问题, 原来得用户名带不过去
func UserTaskCreate(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
		username := r.FormValue("username")
		tpl := template.Must(template.New("userTaskCreate.html").ParseFiles("views/userTaskCreate.html"))
		tpl.Execute(w, username)
	} else {
		username := r.PostFormValue("username")
		name := r.PostFormValue("name")
		desc := r.PostFormValue("desc")
		fmt.Println(username, name, desc)
		if err := modules.CreateTask(username, name, desc); err != nil {
			w.WriteHeader(http.StatusForbidden)
		} else {
			//w.Write([]byte("添加成功."))
			// 这里创建task后 302跳转还有点问题, 原来得用户名带不过去
			http.Redirect(w, r, "/user/userTaskList/", http.StatusFound)
		}
	}
}

func ModifyUser(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		user, err := modules.GetUserById(id)
		fmt.Println("modify user:", user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		tpl := template.Must(template.New("userModify.html").ParseFiles("views/userModify.html"))
		tpl.Execute(w, user)
	} else if r.Method == http.MethodPost {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		username := r.PostFormValue("username")
		birthday, err := time.Parse("2006-01-02", r.PostFormValue("birthday"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		addr := r.PostFormValue("addr")
		desc := r.PostFormValue("desc")
		modules.ModifyUser(id, username, birthday, addr, desc)
		http.Redirect(w, r, "/user/listUser/", http.StatusFound)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		err = modules.DeleteUser(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			http.Redirect(w, r, "/user/listUser/", http.StatusFound)
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func ModifyUserPassword(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
		username := r.FormValue("username")
		user, err := modules.GetUserByName(username)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		tpl := template.Must(template.New("userModifyPassword.html").ParseFiles("views/userModifyPassword.html"))
		tpl.Execute(w, user)
	} else if r.Method == http.MethodPost {
		username := r.PostFormValue("username")
		password := fmt.Sprintf("%x", []byte(r.PostFormValue("password")))
		err := modules.ModifyUserPassword(username, password)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			http.Redirect(w, r, "/user/userTaskList/", http.StatusFound)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}