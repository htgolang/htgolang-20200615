package forms

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"strings"
	"todolist/models"
)

type TaskCreateForm struct {
	Name   string `form:"name,text,名称"`
	Worker string `form:"worker,text,执行者"`
	Desc   string `form:"desc,text,描述"`
}

func (this *TaskCreateForm) Valid(v *validation.Validation) {
	this.Name = strings.TrimSpace(this.Name)
	this.Worker = strings.TrimSpace(this.Worker)
	this.Desc = strings.TrimSpace(this.Desc)

	v.MinSize(this.Name, 2, "name.name").Message("用户名称必须在%d到%d个字符之间", 2, 32)
	v.MaxSize(this.Name, 32, "name.name").Message("用户名称必须在%d到%d个字符之间", 2, 32)

	v.MinSize(this.Worker, 2, "worker.worker").Message("执行者名称必须在%d到%d个字符之间", 2, 32)
	v.MaxSize(this.Worker, 32, "worker.worker").Message("执行者名称必须在%d到%d个字符之间", 2, 32)

	v.MaxSize(this.Desc, 128, "desc.desc").Message("描述必须在%d个字符之内", 128)
}

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

func (this *TaskModifyForm) Valid(v *validation.Validation) {
	this.Name = strings.TrimSpace(this.Name)
	this.Worker = strings.TrimSpace(this.Worker)
	this.Desc = strings.TrimSpace(this.Desc)

	v.MinSize(this.Name, 2, "name.name").Message("用户名称必须在%d到%d个字符之间", 2, 32)
	v.MaxSize(this.Name, 32, "name.name").Message("用户名称必须在%d到%d个字符之间", 2, 32)

	v.MinSize(this.Worker, 2, "worker.worker").Message("执行者名称必须在%d到%d个字符之间", 2, 32)
	v.MaxSize(this.Worker, 32, "worker.worker").Message("执行者名称必须在%d到%d个字符之间", 2, 32)

	v.MaxSize(this.Desc, 128, "desc.desc").Message("描述必须在%d个字符之内", 128)
	v.Range(this.Progress, 0, 100, "progress.progress").Message("进度不正确")

	task := models.Task{Id: this.Id}
	if orm.NewOrm().Read(&task) == nil {
		if this.User.IsSuper || this.User.Id == task.CreateUser {
			this.Task = &task
		}
	} else {
		v.SetError("name", "对象不存在")
	}
	if _, ok := models.TaskStatusTexts[this.Status]; !ok {
		v.SetError("status", "状态不正确")
	}
}
