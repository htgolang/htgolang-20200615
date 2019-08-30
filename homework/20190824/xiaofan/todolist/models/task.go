package models

import (
	"encoding/json"
	"errors"
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

func loadTasks() ([]Task, error) {

	if bytes, err := ioutil.ReadFile("datas/tasks.json"); err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	} else {
		var tasks []Task
		if err := json.Unmarshal(bytes, &tasks); err == nil {
			return tasks, nil
		} else {
			return nil, err
		}
	}

}

func storeTasks(tasks []Task) error {
	bytes, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("datas/tasks.json", bytes, 0x066)
}

func GetTasks() []Task {
	tasks, err := loadTasks()
	if err == nil {
		return tasks
	}
	panic(err)
}

func GetTaskId() (int, error) {
	tasks, err := loadTasks()
	if err != nil {
		return -1, err
	}
	var id int
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
	task := Task{
		Id:       id,
		Name:     name,
		User:     user,
		Desc:     desc,
		Progress: 0,
		Status:   "new",
	}
	tasks, err := loadTasks()
	if err != nil {
		panic(err)
	}
	tasks = append(tasks, task)
	storeTasks(tasks)
}

func GetTaskById(id int) (Task, error) {
	tasks, err := loadTasks()
	if err != nil {
		panic(err)
	}
	for _, task := range tasks {
		if id == task.Id {
			return task, nil
		}

	}
	return Task{}, errors.New("not found")
}

func ModifyTask(id, progress int, name, user, status string) {
	tasks, err := loadTasks()
	if err != nil {
		panic(err)
	}

	for i, task := range tasks {
		if task.Id == id {
			tasks[i].Name = name
			tasks[i].Progress = progress
			tasks[i].User = user
			tasks[i].Status = status
		}
	}
	storeTasks(tasks)
}

func DeleteTask(id int) {
	tasks, err := loadTasks()
	if err != nil {
		panic(err)
	}
	newTask := make([]Task, 0)
	for _, task := range tasks {
		if task.Id == id {
			continue
		}
		newTask = append(newTask, task)
	}

	storeTasks(newTask)
}
