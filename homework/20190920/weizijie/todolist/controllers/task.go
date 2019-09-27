package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"todolist/models"
	"todolist/session"
	"todolist/utils"
)

func TaskAction(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultManager.SessionStart(w, r)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
	} else {
		// tpl := template.Must(template.New("task.html").ParseFiles("views/task/task.html"))
		// tpl.Execute(w, models.GetTasks())
		utils.Render(w, "base.html", []string{"views/layouts/base.html", "views/task/task.html"}, models.GetTasks())
	}
}

func TaskCreateAction(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultManager.SessionStart(w, r)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
	} else {
		if r.Method == http.MethodGet {
			// tpl := template.Must(template.New("create_task.html").ParseFiles("views/task/create_task.html"))
			// tpl.Execute(w, nil)
			utils.Render(w, "base.html", []string{"views/layouts/base.html", "views/task/create_task.html"}, nil)
		} else if r.Method == http.MethodPost {

			name := r.PostFormValue("name")
			user := r.PostFormValue("user")
			desc := r.PostFormValue("desc")
			models.CreateTask(name, user, desc)

			http.Redirect(w, r, "/tasks/", http.StatusFound)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func TaskModifyAction(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultManager.SessionStart(w, r)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}
	if r.Method == http.MethodGet {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err == nil {
			task, err := models.GetTaskById(id)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				// tpl := template.Must(template.New("modify_task.html").ParseFiles("views/task/modify_task.html"))
				// tpl.Execute(w, task)
				utils.Render(w, "base.html", []string{"views/layouts/base.html", "views/task/modify_task.html"}, task)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}

	} else if r.Method == http.MethodPost {
		id, err := strconv.Atoi(r.PostFormValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		name := r.PostFormValue("name")
		desc := r.PostFormValue("desc")
		progress, err := strconv.Atoi(r.PostFormValue("progress"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		user := r.PostFormValue("user")
		status, err := strconv.Atoi(r.PostFormValue("status"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		models.ModifyTask(id, name, desc, progress, user, status)

		http.Redirect(w, r, "/tasks/", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func TaskDeleteAction(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultManager.SessionStart(w, r)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}

	if id, err := strconv.Atoi(r.FormValue("id")); err == nil {
		models.DeleteTask(id)
	} else {
		fmt.Println(err)
	}
	http.Redirect(w, r, "/tasks/", http.StatusFound)
}

func init() {

	http.HandleFunc("/tasks/", TaskAction)
	http.HandleFunc("/tasks/create/", TaskCreateAction)
	http.HandleFunc("/tasks/modify/", TaskModifyAction)
	http.HandleFunc("/tasks/delete/", TaskDeleteAction)
}
