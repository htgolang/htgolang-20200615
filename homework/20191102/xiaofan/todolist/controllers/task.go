package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"strings"
	"time"
	"todolist/forms"
	"todolist/models"
)

type TaskController struct {
	LoginRequiredController
}

// 访问前验证，session存在则继续，否则跳转至登录界面
func (this *TaskController) Prepare() {
	this.LoginRequiredController.Prepare()
}

func (this *TaskController) Index() {
	// 获取查询字段
	q := this.GetString("q")
	q = strings.TrimSpace(q)

	// 初始化查询字段，这时查询条件为空
	condition := orm.NewCondition()

	// 如果存在查询字段，则检查name和desc是否包含查询字段
	if q != "" {
		condition = condition.Or("name__icontains", q)
		condition = condition.Or("desc__icontains", q)
		condition = condition.AndCond(condition)
	}

	// 如果不是管理员用户，则增加一条，只查询属于自己的数据的条件
	if !this.User.IsSuper {
		condition = condition.And("create_user__exact", this.User.Id)
	}

	// 生成tasks来接收查询返回值
	var tasks []models.Task
	// 根据条件查询
	_, _ = orm.NewOrm().QueryTable(&models.Task{}).SetCond(condition).All(&tasks)

	// 将数据返回
	this.Layout = "layout/base.html"
	this.Data["nav"] = "task"
	this.LayoutSections = map[string]string{}
	this.LayoutSections["LayoutScripts"] = "task/index_scripts.html"
	this.TplName = "task/index.html"
	this.Data["tasks"] = tasks
	this.Data["q"] = q
	this.Data["statusTexts"] = models.TaskStatusTexts
}

func (this *TaskController) Create() {
	json := map[string]interface{}{
		"code":   405,
		"text":   "请求方式错误",
		"result": nil,
	}

	form := &forms.TaskCreateForm{}
	valid := &validation.Validation{}

	// 验证请求方法
	if this.Ctx.Input.IsPost() {
		json["code"], json["text"] = 400, "请求数据错误"
		// 验证数据输入是否正确
		if err := this.ParseForm(form); err != nil {
			json["text"] = err.Error()
		} else {
			fmt.Println(form)
			if corret, err := valid.Valid(form); err != nil {
				json["text"] = err.Error()
			} else if !corret {
				json["result"] = valid.Errors
			} else {
				task := &models.Task{
					Name:       form.Name,
					Worker:     form.Worker,
					CreateUser: this.User.Id,
					Desc:       form.Desc,
				}
				// 验证通过则插入用户数据
				o := orm.NewOrm()
				if _, err := o.Insert(task); err == nil {
					json["code"], json["text"], json["result"] = 200, "创建成功", task
				} else {
					json["code"], json["text"], json["result"] = 500, "服务器端错误", nil
				}
			}
		}
	}

	//  失败则返回错误信息至创建页面
	this.Data["json"] = json
	this.ServeJSON()
}

func (this *TaskController) Detail() {
	json := map[string]interface{}{
		"code":   400,
		"text":   "请求数据错误",
		"result": nil,
	}

	if id, err := this.GetInt("id"); err == nil {
		task := &models.Task{Id: id}
		if orm.NewOrm().Read(task) == nil && (this.User.IsSuper || task.CreateUser == this.User.Id) {
			json["code"], json["text"], json["result"] = 200, "获取数据成功", task
		}
	}
	this.Data["json"] = json
	this.ServeJSON()
}

func (this *TaskController) Modify() {
	json := map[string]interface{}{
		"code":   405,
		"text":   "请求方式错误",
		"result": nil,
	}

	form := &forms.TaskModifyForm{User: this.User}
	valid := &validation.Validation{}

	// 验证请求方法
	if this.Ctx.Input.IsPost() {
		json["code"], json["text"] = 400, "请求数据错误"
		if err := this.ParseForm(form); err != nil {
			json["text"] = err.Error()
		} else {
			// 如果校验通过
			if correct, err := valid.Valid(form); err != nil {
				json["text"] = err.Error()
			} else if !correct {
				json["result"] = valid.Errors
			} else {
				// 给form加上用户属性
				form.User = this.User
				// 给任务结构体赋值
				form.Task.Name = form.Name
				form.Task.Progress = form.Progress
				form.Task.Status = form.Status
				form.Task.Worker = form.Worker
				form.Task.Desc = form.Desc

				// 如果进度为完成，则将当前时间写入完成时间
				if form.Status == models.TaskStatusComplete {
					now := time.Now()
					form.Task.CompleteTime = &now
					form.Task.Progress = 100
				}

				// 更新数据
				if _, err := orm.NewOrm().Update(form.Task); err == nil {
					json["code"], json["text"], json["result"] = 200, "修改成功", form.Task
				} else {
					json["code"], json["text"], json["result"] = 500, "服务器端错误", nil
				}
			}
		}
	}

	// 将json数据返回
	this.Data["json"] = json
	this.ServeJSON()
}

func (this *TaskController) Delete() {
	json := map[string]interface{}{
		"code":   405,
		"text":   "请求方式错误",
		"result": nil,
	}
	// 判断请求方法
	if this.Ctx.Input.IsPost() {
		json["code"], json["text"] = 400, "请求数据错误"
		// 获取url 传入的id
		id, _ := this.GetInt("id")
		task := models.Task{Id: id}
		// 根据id删除用户
		if this.User.IsSuper || this.User.Id == task.CreateUser {
			if result, err := orm.NewOrm().Delete(&task); err == nil {
				json["code"], json["text"], json["result"] = 200, "删除成功", result
			} else {
				json["code"], json["text"], json["result"] = 500, "服务器端错误", err.Error()
			}
		}
	}
	this.Data["json"] = json
	this.ServeJSON()
}
