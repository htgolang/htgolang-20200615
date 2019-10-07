package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/imsilence/todolist/forms"
	"github.com/imsilence/todolist/models"
)

// 任务控制器
type TaskController struct {
	LoginRequiredController
}

func (c *TaskController) Prepare() {
	c.LoginRequiredController.Prepare()

	c.Layout = "layout/base.html" // 设置layout
	c.Data["nav"] = "task" // 设置菜单
}

// 任务列表
func (c *TaskController) Index() {
	q := strings.TrimSpace(c.GetString("q", ""))

	var tasks []models.Task

	// 创建查询条件
	condition := orm.NewCondition()
	if q != "" {
		// 查询名称和描述中包含字符
		condition = condition.Or("name__icontains", q)
		condition = condition.Or("desc__icontains", q)
		condition = condition.AndCond(condition)
	}

	// 若非超级管理员只查看当前用户任务
	if !c.User.IsSuper {
		condition = condition.And("create_user__exact", c.User.Id)
	}

	orm.NewOrm().QueryTable(&models.Task{}).SetCond(condition).All(&tasks)

	c.TplName = "task/index.html"
	c.Data["tasks"] = tasks
	c.Data["q"] = q
}

// 任务创建
func (c *TaskController) Create() {
	form := &forms.TaskCreateForm{} // 任务创建表单
	valid := &validation.Validation{} // 验证器

	if c.Ctx.Input.IsPost() {

		// 解析请求参数到form中(根据form标签)
		if c.ParseForm(form) == nil {
			// 表单验证
			if corret, err := valid.Valid(form); err == nil && corret {

				// 创建任务
				task := &models.Task{
					Name:     form.Name,
					Worker: form.Worker,
					Desc:     form.Desc,
					CreateUser: c.User.Id,
				}

				ormer := orm.NewOrm()
				ormer.Insert(task)

				// 使用flash提示操作结果
				flash := beego.NewFlash()
				flash.Success("创建任务成功")
				flash.Store(&c.Controller)
				c.Redirect(beego.URLFor("TaskController.Index"), http.StatusFound)
			}
		}
	}

	c.TplName = "task/create.html"
	c.Data["xsrf_token"] = c.XSRFToken() // csrftoken
	c.Data["form"] = form
	c.Data["validation"] = valid

}


// 任务修改
func (c *TaskController) Modify() {
	form := &forms.TaskModifyForm{User: c.User} // 任务修改表单
	valid := &validation.Validation{} // 验证器
	if c.Ctx.Input.IsGet() {

		// 获取任务信息（超级管理员可查看任意任务信息，普通管理员只能查看自己创建的任务信息）
		if id, err := c.GetInt("id"); err == nil {
			task := models.Task{Id: id}
			if orm.NewOrm().Read(&task) == nil && (c.User.IsSuper || task.CreateUser == c.User.Id) {
				form.Id = task.Id
				form.Name = task.Name
				form.Progress = task.Progress
				form.Status = task.Status
				form.Worker = task.Worker
				form.Desc = task.Desc
			}
		}
	} else if c.Ctx.Input.IsPost() {
		// 任务修改

		// 解析请求参数到form中(根据form标签)
		if c.ParseForm(form) == nil {
			if corret, err := valid.Valid(form); err == nil && corret {
				// 验证任务信息（超级管理员可修改任意任务信息，普通管理员只能修改自己创建的任务信息）

				form.Task.Name = form.Name
				form.Task.Progress = form.Progress
				form.Task.Status = form.Status
				form.Task.Worker = form.Worker
				form.Task.Desc = form.Desc

				// 任务完成修改完成时间和进度
				if form.Status == models.TastStatusComplete {
					now := time.Now()
					form.Task.CompleteTime = &now
					form.Task.Progress = 100
				}

				// 任务修改
				ormer := orm.NewOrm()
				ormer.Update(form.Task)

				// 通过flash提示修改结果
				flash := beego.NewFlash()
				flash.Success("修改任务成功")
				flash.Store(&c.Controller)
				c.Redirect(beego.URLFor("TaskController.Index"), http.StatusFound)
			}
		}
	}

	c.TplName = "task/modify.html"
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["form"] = form
	c.Data["validation"] = valid
	c.Data["statusTexts"] = models.TaskStatusTexts
}

// 任务删除逻辑
func (c *TaskController) Delete() {
	if id, err := c.GetInt("id"); err == nil {
		// 超级管理员可删除任意任务信息，普通管理员只能删除自己创建的任务信息
		ormer := orm.NewOrm()
		task := models.Task{Id: id}
		if ormer.Read(&task) == nil && (c.User.IsSuper || task.CreateUser == c.User.Id) {
			ormer.Delete(&task) // 任务删除

			//通过flash提示操作结果
			flash := beego.NewFlash()
			flash.Success("删除任务成功")
			flash.Store(&c.Controller)
		}
	}
	c.Redirect(beego.URLFor("TaskController.Index"), http.StatusFound)
}
