package controllers

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"todolist/models"
)

func User(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("JevonWei 用户系统")
	http.Redirect(w, r, "/login", http.StatusFound)
}

func UserOperation(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tpl := template.Must(template.New("operation.html").ParseFiles("views/user/operation.html"))
		tpl.Execute(w, nil)

	}
}

func ModifyPasswd(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		tpl := template.Must(template.New("modify_passwd.html").ParseFiles("views/user/modify_passwd.html"))
		tpl.Execute(w, models.AccountMap)
	} else if r.Method == http.MethodPost {
		name := r.PostFormValue("name")
		passwd := r.PostFormValue("passwd")
		passwd_md5 := fmt.Sprintf("%x", md5.Sum([]byte(passwd)))
		models.ModifyPasswd(name, passwd_md5)
		log.Printf("用户(%s)密码已修改", name)
		http.Redirect(w, r, "/user", http.StatusFound)
	}
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tpl := template.Must(template.New("login.html").ParseFiles("views/user/login.html"))
		tpl.Execute(w, nil)

	} else if r.Method == http.MethodPost {
		name := r.PostFormValue("name")
		passwd := r.PostFormValue("passwd")

		passwd_md5 := fmt.Sprintf("%x", md5.Sum([]byte(passwd)))

		if r.PostFormValue("Login") != "" {

			if Verify, err := models.AccountVerify(name, passwd_md5); Verify {
				models.AccountMap[name] = models.GetUsers()
				log.Printf("用户(%s)已登录", name)
				http.Redirect(w, r, "/user/operation", http.StatusFound)
			} else {

				http.Redirect(w, r, "/login/", http.StatusBadRequest)
				w.Write([]byte("用户验证失败，请重新登录"))
				log.Printf("用户(%s)验证失败:%v", name, err)
				//fmt.Println(err)
				// fmt.Printf("%T, %v\n", errors.New("Account not Exist"), errors.New("Account not Exist"))
				// fmt.Printf("%T, %v\n", err, err)
				// if err == errors.New("Verify Failed") {
				// 	fmt.Printf("%T, %v\n", err, err)
				// }
				//w.Write([]byte(string(err)))
			}
		} else if r.PostFormValue("Register") != "" {
			err := models.AccountCreate(name, passwd_md5)
			if err == nil {
				w.Write([]byte("账号注册成功，请使用该账号重新登录"))
				log.Printf("账号(%s)已注册", name)
				http.Redirect(w, r, "/login", http.StatusFound)
			} else {
				w.Write([]byte("Account Exist"))
				w.Write([]byte("\n"))
				w.Write([]byte("账号已存在，请重新注册"))
				log.Printf("账号(%s)已存在，请重新注册", name)
				http.Redirect(w, r, "/login", http.StatusBadRequest)
				//w.Write([]byte(err))
				// fmt.Printf("%T, %v\n", errors.New("Account Exist"), errors.New("Account Exist"))
				// fmt.Printf("%T, %v\n", err, err)
				// if err == errors.New("Account Exist") {
				// 	fmt.Printf("%T, %v\n", err, err)
				// }

			}

		}
	}
}

func UserAction(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.New("user.html").ParseFiles("views/user/user.html"))
	tpl.Execute(w, models.GetUsers())
	//fmt.Println(models.AccountMap)
	//tpl.Execute(w, []User{})
}

func UserCreateAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tpl := template.Must(template.New("create_user.html").ParseFiles("views/user/create_user.html"))
		tpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		name := r.PostFormValue("name")
		addr := r.PostFormValue("addr")
		tel := r.PostFormValue("tel")
		birthday := r.PostFormValue("birthday")
		desc := r.PostFormValue("desc")

		models.CreateUser(name, birthday, addr, tel, desc)

		http.Redirect(w, r, "/user/operation", http.StatusFound)
	} else {
		log.Println("UserCreate POST Faild")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func UserModifyAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err == nil {
			user, err := models.GetUserById(id)
			if err != nil {
				log.Printf("用户(%s)获取失败", user)
				w.WriteHeader(http.StatusBadRequest)
			} else {
				//fmt.Println(user.Birthday)
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
		// progress, err := strconv.Atoi(r.PostFormValue("progress"))
		// if err != nil {
		// 	w.WriteHeader(http.StatusBadRequest)
		// }
		tel := r.PostFormValue("tel")
		birthday := r.PostFormValue("birthday")

		models.ModifyUser(id, name, birthday, addr, tel, desc)

		http.Redirect(w, r, "/user/operation", http.StatusFound)
	} else {
		log.Println("UserModify POST Faild")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func UserDeleteAction(w http.ResponseWriter, r *http.Request) {
	if id, err := strconv.Atoi(r.FormValue("id")); err == nil {
		models.DeleteUser(id)
	} else {
		log.Printf("用户ID(%s)获取失败：%v", id, err)
		fmt.Println(err)
	}
	http.Redirect(w, r, "/user/", http.StatusFound)
}

func UserQueryAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tpl := template.Must(template.New("query_user.html").ParseFiles("views/user/query_user.html"))
		tpl.Execute(w, nil)
		log.Println("Query Users Success")

	} else if r.Method == http.MethodPost {
		name := r.PostFormValue("name")
		if user, err := models.GetUserByName(name); err == nil {
			tpl := template.Must(template.New("getquery_user.html").ParseFiles("views/user/getquery_user.html"))
			tpl.Execute(w, user)
			log.Printf("%s Query Success", user)
		} else {
			log.Println("Query User Not Exist")
			w.Write([]byte("查询的用户不存在"))
		}
	}

}

func init() {
	models.Init()
	http.HandleFunc("/", User)
	http.HandleFunc("/login", UserLogin)
	http.HandleFunc("/passwd/modify/", ModifyPasswd)
	http.HandleFunc("/user/", UserAction)
	http.HandleFunc("/user/query/", UserQueryAction)
	http.HandleFunc("/user/create/", UserCreateAction)
	http.HandleFunc("/user/modify/", UserModifyAction)
	http.HandleFunc("/user/delete/", UserDeleteAction)
	http.HandleFunc("/user/operation/", UserOperation)

}
