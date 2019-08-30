package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"usertask/modules"
	"usertask/session"
)

func init(){
	http.HandleFunc("/tasks/", TaskAction)
	http.HandleFunc("/tasks/create/", TaskCreateAction)
	http.HandleFunc("/tasks/modify/", TaskModifyAction)
	http.HandleFunc("/tasks/delete/", TaskDeleteAction)
}

func TaskAction(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultManager.SessionStart(w, r)
	if _, ok := sessionObj.GET("user"); !ok {
		http.Redirect(w, r, "/users/login/", http.StatusFound)
	} else {
		tpl := template.Must(template.New("task.html").ParseFiles("views/task/task.html"))
		tpl.Execute(w, modules.GetTasks())
	}
}

func TaskCreateAction(w http.ResponseWriter, r *http.Request){
	sessionObj := session.DefaultManager.SessionStart(w, r)
	if _, ok := sessionObj.GET("user"); !ok {
		http.Redirect(w, r, "/users/login/", http.StatusFound)
		return
	}
	if r.Method == http.MethodGet {
		tpl := template.Must(template.New("create_task.html").ParseFiles("views/task/create_task.html"))
		tpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		name := r.PostFormValue("name")
		user := r.PostFormValue("user")
		desc := r.PostFormValue("desc")
		modules.CreateTask(name, user, desc)
		fmt.Println("create task: name, user, desc:", name, user, desc)
		http.Redirect(w, r, "/tasks/", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func TaskModifyAction(w http.ResponseWriter, r *http.Request){
	sessionObj := session.DefaultManager.SessionStart(w, r)
	if _, ok := sessionObj.GET("user"); !ok {
		http.Redirect(w, r, "/users/login/", http.StatusFound)
		return
	}
	if r.Method == http.MethodGet {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		task, err := modules.GetTaskById(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		tpl := template.Must(template.New("modify_task.html").ParseFiles("views/task/modify_task.html"))
		tpl.Execute(w, task)
	} else if r.Method == http.MethodPost {
		id, err := strconv.Atoi(r.PostFormValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		progress, err := strconv.Atoi(r.PostFormValue("progress"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		name := r.PostFormValue("name")
		user := r.PostFormValue("user")
		desc := r.PostFormValue("desc")
		status := r.PostFormValue("status")
		fmt.Println("id, name, progress, user, desc, status:", id, name, progress, user, desc, status)
		modules.ModifyTask(id, name, progress, user, desc, status)
		http.Redirect(w, r, "/tasks/", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func TaskDeleteAction(w http.ResponseWriter, r *http.Request){
	sessionObj := session.DefaultManager.SessionStart(w, r)
	if _, ok := sessionObj.GET("user"); !ok {
		http.Redirect(w, r, "/users/login/", http.StatusFound)
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

