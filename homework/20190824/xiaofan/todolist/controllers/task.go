package controllers

import (
	"html/template"
	"net/http"
	"strconv"
	"todolist/models"
	"todolist/session"
)

// task展示页面
func TaskAction(w http.ResponseWriter, r *http.Request) {
	// session认证
	sessionObj := session.DefaultSessionManager.SessionStart(w, r)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
	}

	tpl := template.Must(template.New("task.html").ParseFiles("views/tasks/task.html"))
	tpl.Execute(w, models.GetTasks())
}

func TaskCreateAction(w http.ResponseWriter, r *http.Request) {
	// session认证
	sessionObj := session.DefaultSessionManager.SessionStart(w, r)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
	}

	if r.Method == http.MethodGet {
		tpl := template.Must(template.New("create_task.html").ParseFiles("views/tasks/create_task.html"))
		tpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		name := r.PostFormValue("name")
		desc := r.PostFormValue("desc")
		user := r.PostFormValue("user")

		models.CreateTask(name, user, desc)
		http.Redirect(w, r, "/tasks/", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func TaskModifyAction(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultSessionManager.SessionStart(w, r)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
	}

	if r.Method == http.MethodGet {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err == nil {
			task, err := models.GetTaskById(id)
			if err != nil {
				w.WriteHeader(http.StatusFound)
			} else {
				tpl := template.Must(template.New("modify_task.html").ParseFiles("views/tasks/modify_task.html"))
				tpl.Execute(w, task)
			}
		} else {
			w.WriteHeader(http.StatusFound)
		}

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
		status := r.PostFormValue("status")

		models.ModifyTask(id, progress, name, user, status)
		http.Redirect(w, r, "/tasks/", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func TaskDeleteAction(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultSessionManager.SessionStart(w, r)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
	}

	if r.Method == http.MethodGet {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		models.DeleteTask(id)
		http.Redirect(w, r, "/tasks/", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func init() {
	http.HandleFunc("/tasks/", TaskAction)
	http.HandleFunc("/tasks/create/", TaskCreateAction)
	http.HandleFunc("/tasks/modify/", TaskModifyAction)
	http.HandleFunc("/tasks/delete/", TaskDeleteAction)

}
