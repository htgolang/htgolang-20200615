package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"todolist/modules"
	"todolist/session"
	"todolist/utils"
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
		statusFunc := modules.SFunc()
		utils.Render(w, "base.html", []string{"views/layouts/base.html","views/task/task.html"}, modules.GetTasks(), statusFunc)
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
	utils.Render(w, "base.html", []string{"views/layouts/base.html","views/task/create_task.html"}, context, nil)
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
	utils.Render(w, "base.html", []string{"views/layouts/base.html","views/task/modify_task.html"}, context, nil)
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
