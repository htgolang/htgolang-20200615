package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"todolist/modules"
	"todolist/session"
)

func init() {
	http.HandleFunc("/tasks/", TaskAction)
	http.HandleFunc("/tasks/create/", TaskCreateAction)
	http.HandleFunc("/tasks/modify/", TaskModifyAction)
	http.HandleFunc("/tasks/delete/", TaskDeleteAction)
}

func TaskAction(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultManager.SessionStart(w, r)
	if _, ok := sessionObj.GET("user"); !ok {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		statuFunc := template.FuncMap{
			"statusType": func(t int) string {
				if t == 0 {
					return "新建"
				} else if t == 1 {
					return "正在进行"
				} else if t == 2 {
					return "停止"
				} else if t == 3 {
					return "完成"
				} else {
					return "未知状态"
				}
			},
		}
		tpl := template.Must(template.New("task.html").Funcs(statuFunc).ParseFiles("views/task/task.html"))
		tpl.Execute(w, modules.GetTasks())
	}
}

func TaskCreateAction(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultManager.SessionStart(w, r)
	if _, ok := sessionObj.GET("user"); !ok {
		http.Redirect(w, r, "/users/login/", http.StatusFound)
		return
	}
	var context struct {
		Errors string
		Name   string
		User   string
		Desc   string
	}
	if r.Method == http.MethodPost {
		name := r.PostFormValue("name")
		user := r.PostFormValue("user")
		desc := r.PostFormValue("desc")
		err := modules.CreateTask(name, user, desc)
		if err == nil {
			http.Redirect(w, r, "/tasks/", http.StatusFound)
			return
		}
		context.Errors = fmt.Sprint(err)
		context.Name = name
		context.User = user
		context.Desc = desc
	}
	tpl := template.Must(template.New("create_task.html").ParseFiles("views/task/create_task.html"))
	tpl.Execute(w, context)
}

func TaskModifyAction(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultManager.SessionStart(w, r)
	if _, ok := sessionObj.GET("user"); !ok {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	var context struct {
		Errors string
		ID  uint
		Name   string
		Progress int
		User   string
		Desc   string
		Status int
	}
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	task, _ := modules.GetTaskById(id)
	context.ID = task.ID
	context.Name = task.Name
	context.Progress = task.Progress
	context.User = task.User
	context.Desc = task.Desc
	context.Status = task.Status
	if r.Method == http.MethodPost {
		id, _ := strconv.Atoi(r.PostFormValue("id"))
		progress, _ := strconv.Atoi(r.PostFormValue("progress"))
		name := r.PostFormValue("name")
		user := r.PostFormValue("user")
		desc := r.PostFormValue("desc")
		status, _ := strconv.Atoi(r.PostFormValue("status"))
		err = modules.ModifyTask(id, name, progress, user, desc, status)
		if err == nil {
			http.Redirect(w, r, "/tasks/", http.StatusFound)
			return
		}
		context.Errors = fmt.Sprint(err)
		context.Name = name
		context.User = user
		context.Progress = progress
		context.Desc = desc
		context.Status = status
	}
	tpl := template.Must(template.New("modify_task.html").ParseFiles("views/task/modify_task.html"))
	tpl.Execute(w, context)

}

func TaskDeleteAction(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultManager.SessionStart(w, r)
	if _, ok := sessionObj.GET("user"); !ok {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	if r.Method == http.MethodGet {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		modules.DeleteTask(id)
		fmt.Println("modify task id:", id)
		http.Redirect(w, r, "/tasks/", http.StatusFound)
	}
}
