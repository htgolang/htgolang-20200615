package modules

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Task struct {
	gorm.Model
	Name          string    `gorm:"type:varchar(256); unique; not null; default:''"`
	Progress      int       `gorm:"not null; default: 0"`
	User          string    `gorm:"type: varchar(64); not null; default: ''"`
	Desc          string    `gorm:"type: varchar(512); column: description; not null; default: ''"`
	Status        int       `gorm:"not null; default: 0"`
	Create_time   *time.Time `gorm:"column: create_time; type: datetime"`
	Complate_time *time.Time `gorm:"column: complate_time; type: datetime"`
}

func (t Task) TableName() string {
	return "todolist_task"
}

func GetTasks() []Task {
	var tasks []Task
	if db.Find(&tasks).Error == nil {
		return tasks
	} else {
		panic("Get Task Failed.")
	}
	return tasks
}

func CreateTask(name, user, desc string) error {
	now := time.Now()
	task := Task{
		Name:        name,
		User:        user,
		Desc:        desc,
		Create_time: &now,
	}
	if err := db.Create(&task).Error; err == nil {
		return nil
	} else {
		return err
	}
}

func GetTaskById(id int) (Task, error) {
	var task Task
	err := db.First(&task, "id = ?", id).Error
	return task, err
}

func ModifyTask(id int, name string, progress int, user, desc string, status int) error {
	var task Task
	if err := db.First(&task, "id = ?", id).Error; err == nil {
		task.Name = name
		task.Progress = progress
		task.User = user
		task.Desc = desc
		task.Status = status
		if status == 3 {
			now := time.Now()
			task.Complate_time = &now
		}
		if err := db.Save(&task).Error; err == nil {
			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}

func DeleteTask(id int) {
	var task Task
	if db.First(&task, "id = ?", id).Error == nil {
		db.Delete(&task)
	}
}
