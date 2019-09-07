package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"todolist/models"
)

type Context struct {
	Errors   string
	Name     string
	Password string
	Birthday string
	Tel      string
	Addr     string
	Desc     string
}

//这个变量不应该放在全局，放在每个函数里边，用的时候声明成不同的结构，会更方便。
var errors string

func UserAction(w http.ResponseWriter, r *http.Request) {
	query := strings.TrimSpace(r.FormValue("query"))
	users := models.GetUsers(query)

	context := struct {
		Query string
		Users []models.User
	}{query, users}
	fmt.Println(users)

	tpl := template.Must(template.New("user.html").ParseFiles("views/users/user.html"))
	_ = tpl.Execute(w, context)
}

func UserCreateAction(w http.ResponseWriter, r *http.Request) {
	//给输入提示提供默认值
	context := Context{"", "用户名请使用英文字母和_", "", "2018-09-03", "", "", ""}

	if r.Method == http.MethodPost {
		name := r.PostFormValue("name")
		passwd := r.PostFormValue("password")
		birthday := r.PostFormValue("birthday")
		sex := r.PostFormValue("sex")
		tel := r.PostFormValue("tel")
		addr := r.PostFormValue("addr")
		desc := r.PostFormValue("desc")
		//数据验证环节，如果有不合格的注册信息，直接修改error的值。下次作业再补

		err := models.CreateUser(name, passwd, birthday, sex, tel, addr, desc)
		if err != nil {
			errors = fmt.Sprintf("%s", err)
		} else {
			http.Redirect(w, r, "/users/", http.StatusFound)
		}

		//修改模板数据
		context = Context{errors, name, passwd, birthday, tel, addr, desc}
	}
	//返回模板+数据
	tpl := template.Must(template.New("create_user.html").ParseFiles("views/users/create_user.html"))
	_ = tpl.Execute(w, context)
}

func UserModifyAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		//展示用户存储的信息
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			panic(err)
		} else {
			user, err := models.GetUserById(id)
			fmt.Println(err)
			errors = fmt.Sprintf("%s", err)
			tpl := template.Must(template.New("modify_user.html").ParseFiles("views/users/modify_user.html"))
			_ = tpl.Execute(w, struct {
				Errors string
				User   models.User
			}{errors, user})
			return
		}
		w.WriteHeader(http.StatusBadRequest)
	} else if r.Method == http.MethodPost {
		//获取提交的表单
		id := r.FormValue("id")
		name := r.FormValue("name")
		birthday := r.FormValue("birthday")
		sex := r.FormValue("sex")
		tel := r.FormValue("tel")
		addr := r.FormValue("addr")
		desc := r.FormValue("desc")

		err := models.ModifyUser(id, name, birthday, sex, tel, addr, desc)
		if err != nil {
			errors = fmt.Sprintf("%s", err)
			fmt.Fprint(w, errors)
		}
		//需要在web上展示修改错误的信息，暂时先不做
		http.Redirect(w, r, "/users/", http.StatusFound)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		//不支持其他方法
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func UserDeleteAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := models.DeleteUser(r.FormValue("id"))
		if err != nil {
			fmt.Fprint(w, err)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	http.Redirect(w, r, "/users/", http.StatusFound)
}

func LoginAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tpl := template.Must(template.New("login.html").ParseFiles("views/login.html"))
		_ = tpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		err := models.CheckUser(username, password)
		if err != nil {
			errors = fmt.Sprintf("%s", err)
			tpl := template.Must(template.New("login.html").ParseFiles("views/login.html"))
			tpl.Execute(w, struct {
				Errors   string
				UserName string
			}{errors, username})
		} else {
			//登陆成功，第一次登陆时需要生成session，这里是session的创建的地方
			http.Redirect(w, r, "/users/", http.StatusFound)
		}
	} else {
		//其他请求，跳转到登陆页面
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func LogoutAction(w http.ResponseWriter, r *http.Request) {
	//删掉存储的session，设置cooker过期时间等于-1,跳转到登陆界面
	http.Redirect(w, r, "/login/", http.StatusFound)
}

func UserModifyPasswordAction(w http.ResponseWriter, r *http.Request) {
	var username string
	/*验证session是否处于登陆状态
	验证登陆成功username = user
	验证登陆失败http.Redirect(w, r, "/login/", http.StatusFound)
	*/
	if r.Method == http.MethodGet {
		tpl := template.Must(template.New("modify_password.html").ParseFiles("views/modify_password.html"))
		_ = tpl.Execute(w, username)
	} else if r.Method == http.MethodPost {
		_ = r.PostFormValue("username")
		_ = r.PostFormValue("password")
		_ = r.PostFormValue("password2")
		/*验证两次密码是否一致
		失败的话返回UserModifyPasswordAction_Get_tpl
		成功的话LogoutAction
		*/
		http.Redirect(w, r, "/login/", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
func init() {
	http.HandleFunc("/", LoginAction)
	http.HandleFunc("/logout/", LogoutAction)
	http.HandleFunc("/users/", UserAction)
	http.HandleFunc("/users/password/", UserModifyPasswordAction)
	http.HandleFunc("/users/create/", UserCreateAction)
	http.HandleFunc("/users/modify/", UserModifyAction)
	http.HandleFunc("/users/delete/", UserDeleteAction)

}
