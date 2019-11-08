package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
	"todolist/models"
	"github.com/astaxie/beego/validation"
	"todolist/forms"
)

type TaskController struct {
	LoginRequiredController
}

func (this *TaskController) Prepare() {
	this.LoginRequiredController.Prepare()

}

func (this *TaskController) Index() {
	q := this.GetString("q")
	q = strings.TrimSpace(q)

	condition := orm.NewCondition()
	if q != "" {
		condition = condition.Or("name__icontains", q)
		condition = condition.Or("desc__icontains", q)
		condition = condition.AndCond(condition)
	}

	// 如果不是管理员用户，则增加一条，只查询属于自己的数据的条件
	if !this.User.IsSuper{
		condition = condition.And("create_user__exact", this.User.Id)
	}

	// 生成tasks来接收查询返回值
	var tasks []models.Task
	// 根据条件查询
	orm.NewOrm().QueryTable(&models.Task{}).SetCond(condition).All(&tasks)

	// 将数据返回
	this.Layout = "layout/base.html"
	this.Data["nav"] = "task"

	this.LayoutSections = map[string]string{}
	this.LayoutSections["LayoutScripts"] = "task/index_scripts.html"

	this.TplName = "task/index.html"
	this.Data["xsrf_token"] = this.XSRFToken()
	this.Data["tasks"] = tasks
	this.Data["q"] = q
	this.Data["statusTexts"] = models.TaskStatusTexts
}

func (this *TaskController) Create() {
	json := map[string]interface{}{
		"code": 405,
		"text": "请求方式错误",
		"result": nil,
	}

	if this.Ctx.Input.IsPost() {
		json = map[string]interface{}{
			"code": 400,
			"text": "提交数据错误",
			"result": nil,
		}
		form := &forms.TaskCreateForm{}
		valid := &validation.Validation{}

		// 解析请求参数到form中(根据form标签)
		if err := this.ParseForm(form); err != nil {
			json["text"] = err.Error()
		} else {
			// 表达验证
			if corret, err := valid.Valid(form); err != nil {
				json["text"] = err.Error()
			} else if !corret {
				json["result"] = valid.Errors
			} else {
				task := &models.Task{
					Name:   form.Name,
					Worker: form.Worker,
					CreateUser: this.User.Id,
					Desc:   form.Desc,
				}
				o := orm.NewOrm()
				if _, err := o.Insert(task); err == nil {
					json = map[string]interface{}{
						"code":   200,
						"text":   "创建成功",
						"result": task,
					}
				} else {
						json = map[string]interface{}{
							"code": 500,
							"text": "服务端错误",
							"result": nil,
						}
				}
			}
		}
	}

	this.Data["json"] = json
	this.ServeJSON()
}


// 	this.Data["json"] = json
// 	this.ServeJSON()
// }


func (this *TaskController) Modify() {
	json := map[string]interface{} {
		"code": 405,
		"text": "请求方式错误",
		"result": nil,
	}


	form := &forms.TaskModifyForm{User: this.User}   // 任务修改表单
	valid := &validation.Validation{}

	if this.Ctx.Input.IsGet() {
		// 获取任务信息（超级管理员可查看任意任务信息，普通管理员只能查看自己创建的任务信息）
		if id, err := this.GetInt("id"); err == nil {
			task := models.Task{Id: id}
			if orm.NewOrm().Read(&task) == nil && (this.User.IsSuper || task.CreateUser == this.User.Id) {
				json["code"] = 200
				json["text"] = "获取数据成功"
				json["result"] = task
			}
		}
		this.Data["json"] = json
		this.ServeJSON()
	} else if this.Ctx.Input.IsPost() {
		// 任务修改
		json = map[string]interface{} {
			"code": 400,
			"text": "请求数据错误",
			"result": nil,
		}

		// 解析请求参数到form中(根据form标签)
		if err := this.ParseForm(form); err != nil {
			json["text"] = err.Error()
		} else {
			if corret, err := valid.Valid(form); err != nil {
				json["text"] = err.Error()
			} else if  !corret {

				json["result"] = valid.Errors

			} else {
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
				if _, err := orm.NewOrm().Update(form.Task); err == nil {
					json["code"] = 200
					json["text"] = "修改成功"
					json["result"] = form.Task

				}else {

					json["code"] = 500
					json["text"] = "服务器错误"
					json["result"] = nil
				}
			}
		} 	
	}
	this.Data["json"] = json
	this.ServeJSON()

}

// 任务删除
func (this *TaskController) Delete() {

	json := map[string]interface{}{
		"code":   405,
		"text":   "请求方式错误",
		"result": nil,
	}

	if this.Ctx.Input.IsPost() {

		if id, err := this.GetInt("id"); err == nil {

			fmt.Println(id)

			task := models.Task{Id: id}

			if this.User.IsSuper || task.CreateUser == this.User.Id {


				if orm.NewOrm().Read(&task) == nil{

					orm.NewOrm().Delete(&task)

					json["code"] = 200
					json["text"] = "删除数据成功"
					json["result"] = nil

				}

			}else {

				json["code"] = 400
				json["text"] = "获取数据失败，或你没有权限操作"
				json["result"] = nil
			}


		}else {
			json["code"] = 400
			json["text"] = "获取ID错误"
			json["result"] = err.Error()

		}

	}

	this.Data["json"] = json
	this.ServeJSON()
}