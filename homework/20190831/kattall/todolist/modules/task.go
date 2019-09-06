package modules

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

type Task struct {
	Id            int
	Name          string
	Progress      int
	User          string
	Desc          string
	Status        int
	Create_time   time.Time
	Complate_time time.Time
}

func GetTasks() []Task {
	var tasks []Task
	var task Task
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
		os.Exit(-1)
	}

	if err = db.Ping(); err != nil {
		panic(err)
		os.Exit(-1)
	}
	defer db.Close()

	rows, err := db.Query("select * from task")
	for rows.Next() {
		rows.Scan(&task.Id, &task.Name, &task.Progress, &task.User, &task.Desc, &task.Status, &task.Create_time, &task.Complate_time)
		fmt.Println("task:", task.Id, task.Name, task.Progress, task.User, task.Desc, task.Status, task.Create_time, task.Complate_time)
		tasks = append(tasks, task)
	}
	fmt.Println("getTasks:", tasks)
	return tasks
}

func CreateTask(name, user, desc string) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
		os.Exit(-1)
	}
	if err = db.Ping(); err != nil {
		panic(err)
		os.Exit(-1)
	}
	defer db.Close()
	_, err = db.Exec("insert into task(name, user, `desc`, create_time) values(?,?,?,now())", name, user, desc)
	if err != nil {
		panic(err)
	}
}

func GetTaskById(id int) (Task, error) {
	var task Task
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
		os.Exit(-1)
	}
	if err = db.Ping(); err != nil {
		panic(err)
		os.Exit(-1)
	}
	defer db.Close()

	rows := db.QueryRow("select * from task where id = ?", id)
	err = rows.Scan(&task.Id, &task.Name, &task.Progress, &task.User, &task.Desc, &task.Status, &task.Create_time, &task.Complate_time)
	fmt.Println("task:", task.Id, task.Name, task.Progress, task.User, task.Desc, task.Status, task.Create_time, task.Complate_time)
	return task, nil
}

func ModifyTask(id int, name string, progress int, user, desc, status string) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
		os.Exit(-1)
	}
	if err = db.Ping(); err != nil {
		panic(err)
		os.Exit(-1)
	}
	defer db.Close()

	_, err = db.Exec("update task set name=?, progress=?, user=?, `desc`=?, status=? where id = ?", name, progress, user, desc, status, id)
	if err != nil {
		panic(err)
	}
}

func DeleteTask(id int) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
		os.Exit(-1)
	}
	if err = db.Ping(); err != nil {
		panic(err)
		os.Exit(-1)
	}
	defer db.Close()

	_, err = db.Exec("delete from task where id = ?", id)
	if err != nil {
		panic(err)
	}
}

