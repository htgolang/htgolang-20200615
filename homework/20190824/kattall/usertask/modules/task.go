package modules

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type Task struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Progress int    `json:"progress"`
	User     string `json:"user"`
	Desc     string `json:"desc"`
	Status   string `json:"status"`
}

func loadTask() ([]Task, error) {
	bytes, err := ioutil.ReadFile("datas/task/tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	} else {
		var tasks []Task
		if err := json.Unmarshal(bytes, &tasks); err != nil {
			return nil, err
		} else {
			return tasks, nil
		}
	}
}

func storeTask(tasks []Task) error {
	bytes, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("datas/task/tasks.json", bytes, os.ModePerm)
}

func GetTasks() []Task {
	tasks, err := loadTask()
	if err != nil {
		panic(err)
	}
	return tasks
}

func GetTaskId() (int, error) {
	var id int
	tasks, err := loadTask()
	if err != nil {
		return -1, err
	}
	for _, task := range tasks {
		if id < task.Id {
			id = task.Id
		}
	}
	return id + 1, nil
}

func CreateTask(name, user, desc string) {
	id, err := GetTaskId()
	if err != nil {
		panic(err)
	}
	tasks, err := loadTask()
	if err != nil {
		panic(err)
	}
	task := Task{
		Id:       id,
		Name:     name,
		Progress: 0,
		User:     user,
		Desc:     desc,
		Status:   "new",
	}

	tasks = append(tasks, task)
	storeTask(tasks)
}


func GetTaskById(id int) (Task, error){
	tasks, err := loadTask()
	if err != nil {
		return Task{}, err
	}

	for _, task := range tasks {
		if task.Id == id {
			return task, nil
		}
	}
	return Task{}, errors.New("Not Found")
}

func ModifyTask(id int, name string, progress int, user, desc, status string){
	tasks, err := loadTask()
	if err != nil {
		panic(err)
	}
	new_task := make([]Task, len(tasks))
	for i, task := range tasks {
		if task.Id == id {
			task.Name = name
			task.Progress = progress
			task.User = user
			task.Desc = desc
			task.Status = status
		}
		new_task[i] = task
	}
	storeTask(new_task)
	fmt.Println("modify task: new_task:", new_task)
}

func DeleteTask(id int) {
	tasks, err := loadTask()
	if err != nil {
		panic(err)
	}
	new_task := make([]Task, 0)
	for _, task := range tasks {
		if task.Id != id {
			new_task = append(new_task, task)
		}
	}
	storeTask(new_task)
}