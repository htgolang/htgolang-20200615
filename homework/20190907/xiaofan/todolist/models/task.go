package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Task struct {
	gorm.Model
	Name         string    `gorm:"type:varchar(256); not null; default:''"`
	Progress     int       `gorm:"not null; default:0"`
	User         string    `gorm:"type:varchar(32); not null; default:''"`
	Desc         string    `gorm:"column:description; type:varchar(512); not null; default:''"`
	Status       int       `gorm:"not null; default:0"`
	CreateTime   time.Time `gorm:"column:create_time; type:datetime"`
	CompleteTime time.Time `gorm:"column:complete_time; type:datetime"`
}

func GetTasks() []Task {
	var tasks []Task
	db.Find(&tasks)
	return tasks
}

func CreateTask(name, user, desc string) {
	db.Create(&Task{
		Name:       name,
		User:       user,
		Desc:       desc,
		CreateTime: time.Now(),
	})
}

func GetTaskById(id int) (Task, error) {
	var task Task
	err := db.First(&task, "id=?", id).Error
	return task, err

}

func ModifyTask(id, progress int, name, user string, status int) {
	var task Task
	if db.First(&task, "id=?", id).Error == nil {
		task.Name = name
		task.User = user
		task.Progress = progress
		task.Status = status
	}
	db.Save(&task)
}

func DeleteTask(id int) {
	var task Task
	if db.First(&task, "id=?", id).Error == nil {
		db.Delete(&task)
	}
}
