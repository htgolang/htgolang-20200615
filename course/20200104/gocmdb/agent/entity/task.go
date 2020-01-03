package entity

import (
	"encoding/json"
)

type Task struct {
	Id int `json:"id"`

	Plugin  string `json:"plugin"`
	Params  string `json:"params"`
	Timeout int    `json:"timeout"`
}

func NewTask(taskMap interface{}) (Task, error) {
	var task Task
	if taskBytes, err := json.Marshal(taskMap); err != nil {
		return task, err
	} else if err = json.Unmarshal(taskBytes, &task); err != nil {
		return task, err
	}
	return task, nil
}

type Result struct {
	TaskId int    `json:"task_id"`
	Status int    `json:"status"`
	Result string `json:"result"`
	Err    string `json:"err"`
}

func NewResult(task Task, result interface{}, err error) Result {
	bytes, _ := json.Marshal(result)
	errInfo := ""
	status := 0
	if err != nil {
		status = 1
		errInfo = err.Error()
	}
	return Result{
		TaskId: task.Id,
		Status: status,
		Result: string(bytes),
		Err:    errInfo,
	}
}
