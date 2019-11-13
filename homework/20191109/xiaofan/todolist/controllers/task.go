package controllers

import (
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
func (c *TaskController) Prepare() {
	c.LoginRequiredController.Prepare()
}

func (c *TaskController) Index() {
	// 将基础页面返回
	c.Layout = "layout/base.html"
	c.Data["nav"] = "task"
	c.LayoutSections = map[string]string{}
	c.LayoutSections["LayoutScripts"] = "task/index_scripts.html"

	c.TplName = "task/index.html"
	c.Data["statusTexts"] = models.TaskStatusTexts
}

func (c *TaskController) List() {
	orderByColumns := map[string]bool{
		"id":       true,
		"name":     true,
		"status":   true,
		"progress": true,
		"worker":   true,
	}

	// 获取Datatables查询参数
	draw := c.GetString("draw")

	// 获取offset的值
	start, err := c.GetInt("start")
	if err != nil {
		start = 0
	}

	// 获取排序的列名
	orderBy := c.GetString("orderBy")
	if _, ok := orderByColumns[orderBy]; !ok {
		orderBy = "id"
	}

	// 获取升序还是降序
	orderDir := c.GetString("orderDir")
	if orderDir == "desc" {
		orderBy = "-" + orderBy
	}

	// 获取limit的值
	length, err := c.GetInt("length")
	if err != nil {
		length = 10
	}

	// 获取查询条件
	q := c.GetString("q")
	q = strings.TrimSpace(q)

	// 初始化查询条件，这时查询条件为空
	condition := orm.NewCondition()

	// 如果不是管理员用户，则增加一条，只查询属于自己的数据的条件
	if !c.User.IsSuper {
		condition = condition.And("create_user__exact", c.User.Id)
	}
	// 查询数据总量
	o := orm.NewOrm()
	queryset := o.QueryTable(&models.Task{})
	total, _ := queryset.SetCond(condition).Count()

	// 搜索数据条数默认为所有
	totalFilter := total

	// 如果存在查询字段，则增加查询条件
	if q != "" {
		// 初始化一个新的查询条件
		qcondition := orm.NewCondition()
		qcondition = qcondition.Or("name__icontains", q)
		qcondition = qcondition.Or("desc__icontains", q)

		// 在原来的查询条件中，添加关键字段查询
		condition = condition.AndCond(qcondition)

		// 搜索到多少条数据
		totalFilter, _ = queryset.SetCond(condition).Count()
	}

	// 生成tasks来接收查询返回值
	var tasks []*models.Task

	// 根据条件查询
	_, _ = queryset.SetCond(condition).OrderBy(orderBy).Limit(length).Offset(start).All(&tasks)
	for _, task := range tasks {
		task.Patch()
	}
	c.Data["json"] = map[string]interface{}{
		"code":            200,
		"text":            "获取任务成功",
		"result":          tasks,
		"draw":            draw,
		"recordsTotal":    total,
		"recordsFiltered": totalFilter,
	}
	c.ServeJSON()
}

func (c *TaskController) Create() {
	json := map[string]interface{}{
		"code":   405,
		"text":   "请求方式错误",
		"result": nil,
	}

	form := &forms.TaskCreateForm{}
	valid := &validation.Validation{}

	// 验证请求方法
	if c.Ctx.Input.IsPost() {
		json["code"], json["text"] = 400, "请求数据错误"
		// 验证数据输入是否正确
		if err := c.ParseForm(form); err != nil {
			json["text"] = err.Error()
		} else {
			if corret, err := valid.Valid(form); err != nil {
				json["text"] = err.Error()
			} else if !corret {
				json["result"] = valid.Errors
			} else {
				task := &models.Task{
					Name:       form.Name,
					Worker:     form.Worker,
					CreateUser: c.User.Id,
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
	c.Data["json"] = json
	c.ServeJSON()
}

func (c *TaskController) Detail() {
	json := map[string]interface{}{
		"code":   400,
		"text":   "请求数据错误",
		"result": nil,
	}

	if id, err := c.GetInt("id"); err == nil {
		task := &models.Task{Id: id}
		if orm.NewOrm().Read(task) == nil && (c.User.IsSuper || task.CreateUser == c.User.Id) {
			json["code"], json["text"], json["result"] = 200, "获取数据成功", task
		}
	}
	c.Data["json"] = json
	c.ServeJSON()
}

func (c *TaskController) Modify() {
	json := map[string]interface{}{
		"code":   405,
		"text":   "请求方式错误",
		"result": nil,
	}

	form := &forms.TaskModifyForm{User: c.User}
	valid := &validation.Validation{}

	// 验证请求方法
	if c.Ctx.Input.IsPost() {
		json["code"], json["text"] = 400, "请求数据错误"
		if err := c.ParseForm(form); err != nil {
			json["text"] = err.Error()
		} else {
			// 如果校验通过
			if correct, err := valid.Valid(form); err != nil {
				json["text"] = err.Error()
			} else if !correct {
				json["result"] = valid.Errors
			} else {
				// 给form加上用户属性
				form.User = c.User
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
	c.Data["json"] = json
	c.ServeJSON()
}

func (c *TaskController) Delete() {
	json := map[string]interface{}{
		"code":   405,
		"text":   "请求方式错误",
		"result": nil,
	}
	// 判断请求方法
	if c.Ctx.Input.IsPost() {
		json["code"], json["text"] = 400, "请求数据错误"
		// 获取url 传入的id
		id, _ := c.GetInt("id")
		task := models.Task{Id: id}
		// 根据id删除用户
		if c.User.IsSuper || c.User.Id == task.CreateUser {
			if result, err := orm.NewOrm().Delete(&task); err == nil {
				json["code"], json["text"], json["result"] = 200, "删除成功", result
			} else {
				json["code"], json["text"], json["result"] = 500, "服务器端错误", err.Error()
			}
		}
	}
	c.Data["json"] = json
	c.ServeJSON()
}
