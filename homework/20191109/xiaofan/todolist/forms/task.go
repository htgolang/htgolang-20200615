package forms

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"strings"
	"todolist/models"
)

// 创建任务的结构体
type TaskCreateForm struct {
	Name   string `form:"name,text,名称"`
	Worker string `form:"worker,text,执行者"`
	Desc   string `form:"desc,text,描述"`
}

// 创建任务结构体的验证
func (c *TaskCreateForm) Valid(v *validation.Validation) {
	c.Name = strings.TrimSpace(c.Name)
	c.Worker = strings.TrimSpace(c.Worker)
	c.Desc = strings.TrimSpace(c.Desc)

	v.MinSize(c.Name, 2, "name.name").Message("用户名称必须在%d到%d个字符之间", 2, 32)
	v.MaxSize(c.Name, 32, "name.name").Message("用户名称必须在%d到%d个字符之间", 2, 32)

	v.MinSize(c.Worker, 2, "worker.worker").Message("执行者名称必须在%d到%d个字符之间", 2, 32)
	v.MaxSize(c.Worker, 32, "worker.worker").Message("执行者名称必须在%d到%d个字符之间", 2, 32)

	v.MaxSize(c.Desc, 128, "desc.desc").Message("描述必须在%d个字符之内", 128)
}

// 修改任务的结构体
type TaskModifyForm struct {
	Id       int    `form:"id,hidden,ID"`
	Name     string `form:"name,text,名称"`
	Status   int    `form:"status,select,状态"`
	Progress int    `form:"progress,range,进度"`
	Worker   string `form:"worker,text,执行者"`
	Desc     string `form:"desc,text,描述"`

	Task *models.Task
	User *models.User
}

// 修改任务结构体的验证
func (c *TaskModifyForm) Valid(v *validation.Validation) {
	c.Name = strings.TrimSpace(c.Name)
	c.Worker = strings.TrimSpace(c.Worker)
	c.Desc = strings.TrimSpace(c.Desc)

	v.MinSize(c.Name, 2, "name.name").Message("用户名称必须在%d到%d个字符之间", 2, 32)
	v.MaxSize(c.Name, 32, "name.name").Message("用户名称必须在%d到%d个字符之间", 2, 32)

	v.MinSize(c.Worker, 2, "worker.worker").Message("执行者名称必须在%d到%d个字符之间", 2, 32)
	v.MaxSize(c.Worker, 32, "worker.worker").Message("执行者名称必须在%d到%d个字符之间", 2, 32)

	v.MaxSize(c.Desc, 128, "desc.desc").Message("描述必须在%d个字符之内", 128)
	v.Range(c.Progress, 0, 100, "progress.progress").Message("进度不正确")

	task := models.Task{Id: c.Id}
	if orm.NewOrm().Read(&task) == nil {
		if c.User.IsSuper || c.User.Id == task.CreateUser {
			c.Task = &task
		}
	} else {
		_ = v.SetError("name", "对象不存在")
	}
	if _, ok := models.TaskStatusTexts[c.Status]; !ok {
		_ = v.SetError("status", "状态不正确")
	}
}
