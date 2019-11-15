package forms

import (
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"

	"github.com/imsilence/todolist/models"
)

// 任务创建表单
type TaskCreateForm struct {
	Name   string `form:"name,text,名称"`
	Worker string `form:"worker,text,执行者"`
	Desc   string `form:"desc,text,描述"`
}

// 任务创建表单 验证接口（由validation.Valid调用）
func (f *TaskCreateForm) Valid(v *validation.Validation) {
	// 去除用户输入前后空白字符
	f.Name = strings.TrimSpace(f.Name)
	f.Worker = strings.TrimSpace(f.Worker)
	f.Desc = strings.TrimSpace(f.Desc)

	// 使用beego validation提供的验证器验证最小和最大长度
	v.MinSize(f.Name, 2, "name.name").Message("任务名长度必须在%d到%d之间", 2, 32)
	v.MaxSize(f.Name, 32, "name.name").Message("任务名长度必须在%d到%d之间", 2, 32)

	v.MaxSize(f.Worker, 32, "worker.worker").Message("执行者名称长度必须在%d之内", 32)
	v.MaxSize(f.Desc, 128, "desc.desc").Message("描述长度必须在%d之内", 128)
}

// 任务修改表单
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

// 任务修改表单 验证接口（由validation.Valid调用）
func (f *TaskModifyForm) Valid(v *validation.Validation) {
	// 去除用户输入前后空白字符
	f.Name = strings.TrimSpace(f.Name)
	f.Worker = strings.TrimSpace(f.Worker)
	f.Desc = strings.TrimSpace(f.Desc)

	// 验证任务是否存在及用户是否具有修改权限
	task := models.Task{Id: f.Id}
	if orm.NewOrm().Read(&task) != nil {
		v.SetError("name", "对象不存在")
		return
	} else if f.User.IsSuper || f.User.Id == task.CreateUser {
		f.Task = &task
	}

	// 验证任务状态
	if _, ok := models.TaskStatusTexts[f.Status]; !ok {
		v.SetError("status", "状态不正确")
	}

	// 使用beego validation提供的验证器验证进度范围
	v.Range(f.Progress, 0, 100, "progress.progress").Message("进度值不正确")

	// 使用beego validation提供的验证器验证最小和最大长度
	v.MinSize(f.Name, 2, "name.name").Message("任务名长度必须在%d到%d之间", 2, 32)
	v.MaxSize(f.Name, 32, "name.name").Message("任务名长度必须在%d到%d之间", 2, 32)

	v.MaxSize(f.Worker, 32, "worker.worker").Message("执行者名称长度必须在%d之内", 32)
	v.MaxSize(f.Desc, 128, "desc.desc").Message("描述长度必须在%d之内", 128)
}
