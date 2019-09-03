package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Task struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	Progress     int       `json:"progress"`
	User         string    `json:"user"`
	Desc         string    `json:"desc"`
	Status       int       `json:"status"`
	CreateTime   time.Time `json:"create_time"`
	CompleteTime time.Time `json:"complete_time"`
}

func GetTasks() []Task {
	db, err := sql.Open("mysql", dsn)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}
	sql := "select * from todolist_task"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	tasks := make([]Task, 0)
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.Id, &task.Name, &task.Progress, &task.User, &task.Desc, &task.Status, &task.CreateTime, &task.CompleteTime); err != nil {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

func CreateTask(name, user, desc string) {
	db, err := sql.Open("mysql", dsn)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}

	if err := db.Ping(); err != nil {
		fmt.Println(err)
	}

	sql := "insert into todolist_task(name, progress, user, `desc`, status, create_time) values(?,?,?,?,?,now())"
	_, err = db.Exec(sql, name, 0, user, desc, 0)
	if err != nil {
		fmt.Println(err)
	}

}

func GetTaskById(id int) (Task, error) {
	db, err := sql.Open("mysql", dsn)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	sql := "select id, name, user, progress, status, `desc` from todolist_task where id=?;"

	var task Task
	row := db.QueryRow(sql, id)

	err = row.Scan(&task.Id, &task.Name, &task.User, &task.Progress, &task.Status, &task.Desc)
	if err != nil {
		fmt.Println(err)
		return Task{}, errors.New("not found")
	}
	return task, nil

}

func ModifyTask(id, progress int, name, user string, status int) {
	db, err := sql.Open("mysql", dsn)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}
	sql := "update todolist_task set progress=?, name=?, user=?, status=? where id=?"
	_, err = db.Exec(sql, progress, name, user, status, id)

	if err != nil {
		panic(err)
	}
}

func DeleteTask(id int) {
	db, err := sql.Open("mysql", dsn)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	sql := "delete from todolist_task where id=?"
	_, err = db.Exec(sql, id)
	if err != nil {
		panic(err)
	}
}
