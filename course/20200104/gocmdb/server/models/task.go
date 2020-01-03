package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Task struct {
	Id   int    `orm:"column(id);" json:"id"`
	UUID string `orm:"column(uuid);size(64);" json:"uuid"`

	Plugin  string `orm:"column(plugin);size(32);" json:"plugin"`
	Params  string `orm:"column(params);type(text);" json:"params"`
	Timeout int    `orm:"column(timeout);" json:"timeout"`

	Status        int        `orm:"column(status);" json:"status"`
	CreatedTime   *time.Time `orm:"column(created_time);auto_now_add;" json:"created_time"`
	CompletedTime *time.Time `orm:"column(completed_time);null;" json:"completed_time"`
	DeletedTime   *time.Time `orm:"column(deleted_time);null;" json:"deleted_time"`

	Result *Result `orm:"column(result);reverse(one);" json:"result"`
}

type TaskManager struct{}

func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

func (m *TaskManager) Create(uuid string, plugin string, params string, timeout int) error {
	ormer := orm.NewOrm()
	task := &Task{
		UUID:    uuid,
		Plugin:  plugin,
		Params:  params,
		Timeout: timeout,
		Status:  TaskStatusNew,
	}
	if _, err := ormer.Insert(task); err != nil {
		return err
	}
	return nil
}

func (m *TaskManager) GetByUuid(uuid string) []*Task {
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(new(Task))

	condition := orm.NewCondition()
	condition = condition.And("deleted_time__isnull", true)
	condition = condition.And("uuid__exact", uuid)
	condition = condition.And("status__in", TaskStatusNew)

	var result []*Task
	queryset.SetCond(condition).All(&result)
	queryset.SetCond(condition).Update(orm.Params{"status": TaskStatusExecing})
	return result
}

func (m *TaskManager) GetByIdAndUuid(id int, uuid string) *Task {
	ormer := orm.NewOrm()
	task := &Task{Id: id, UUID: uuid}
	if err := ormer.Read(task, "id", "uuid"); err == nil {
		return task
	}
	return nil
}

func (m *TaskManager) Result(uuid string, result *Result) error {
	ormer := orm.NewOrm()
	task := m.GetByIdAndUuid(result.TaskId, uuid)
	if task == nil {
		return fmt.Errorf("针对终端%s任务%s不存在", uuid, result.TaskId)
	}
	now := time.Now()

	task.Status = TaskStatusSuccess
	if result.Status != 0 {
		task.Status = TaskStatusFailure
	}
	task.CompletedTime = &now
	if _, err := ormer.Update(task); err != nil {
		return err
	}

	result.Task = task
	if _, err := ormer.Insert(result); err != nil {
		return err
	}
	return nil
}

var DefaultTaskManager = NewTaskManager()

type Result struct {
	Id     int   `orm:"column(id);" json:"id"`
	Task   *Task `orm:"column(task);rel(one);" json:"task"`
	TaskId int   `orm:"-" json:"task_id"`

	Status int    `orm:"-" json:"status"`
	Result string `orm:"column(result);type(text);" json:"result"`
	Err    string `orm:"column(err);type(text);" json:"err"`

	CreatedTime *time.Time `orm:"column(created_time);auto_now_add;" json:"created_time"`
	DeletedTime *time.Time `orm:"column(deleted_time);null;" json:"deleted_time"`
}

func init() {
	orm.RegisterModel(new(Task), new(Result))
}
