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
	LoginController //继承login控制器，针对所有控制的操作，都可以在login基础控制器内实现
}

func (c *TaskController) Prepare(){
	c.LoginController.Prepare() //task控制器属于登录后操作，所以需要加载login控制器，进行登录验证
}

func (c *TaskController) Index(){
	c.Layout = "layout/base.html"  //还需要调用beego控制器的layout方法，进行页面基本布局
	c.Data["nav"] = "task" //这个是当前处于哪个栏目里的标记，传递给nav.html页面
	c.LayoutSections = map[string]string{}
	c.LayoutSections["LayoutScripts"] = "task/index_scripts.html"

	c.TplName = "task/index.html" //任务列表页的布局样式
	c.Data["xsrf_token"] = c.XSRFToken() //防csrftoken攻击,原理还没理解,需要重新听课
	c.Data["statusTexts"] = models.TaskStatusTexts
}

func (c *TaskController) List() {
	orderByColumns := map[string]bool{
		"id":true,
		"name":true,
		"status":true,
		"progress":true,
		"worker":true,
	}
	draw := c.GetString("draw")
	start, err := c.GetInt("start")
	if  err != nil {
		start = 0
	}
	length, err := c.GetInt("length")
	if  err != nil {
		length = 10
	}
	q := strings.TrimSpace(c.GetString("q","")) //去空格,获取参数(c.getstring)
	orderBy := c.GetString("orderBy")
	if _,ok := orderByColumns[orderBy];!ok {
		orderBy = "id"
	}
	orderDir := c.GetString("orderDir")

	if orderDir == "desc" {
		orderBy = "-" + orderBy
	}

	var tasks []*models.Task  //做一个任务模型切片，准备填充表单。

	condition := orm.NewCondition()

	if !c.User.IsSuper {    //如果不是超管，就在查询条件里，添加任务的创建用户ID=查询用户ID的规则
		condition = condition.And("create_user__exact", c.User.Id)
	}
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&models.Task{})
	total,_ := queryset.SetCond(condition).Count()
	totalFilter := total

	if q != "" {
		qcondition := orm.NewCondition()
		qcondition = qcondition.Or("name__icontains",q)
		qcondition = qcondition.Or("desc__icontains",q)
		qcondition = qcondition.Or("worker__icontains",q)
		condition = condition.AndCond(qcondition)

		totalFilter,_ = queryset.SetCond(condition).Count()
	}

	queryset.SetCond(condition).OrderBy(orderBy).Limit(length).Offset(start).All(&tasks)

	for _,task := range tasks {
		task.Patch()
	}
	c.Data["json"] = map[string]interface{}{
		"code": 200,
		"text": "获取数据成功",
		"result": tasks,
		"draw": draw,
		"recordsTotal": total,
		"recordsFiltered": totalFilter,
	}
	c.ServeJSON()
}

func (c *TaskController) Create(){

	json := map[string]interface{}{
		"code": 405,
		"text": "请求方式错误",
		"result": nil,
	}

	//提交表单数据，采用post方式
	if c.Ctx.Input.IsPost() {
		json = map[string]interface{}{
			"code": 400,
			"text": "提交数据错误",
			"result": nil,
		}
		form := &forms.TaskCreateForm{} //声明创建表单数据模型变量
		valid := &validation.Validation{} //加载验证器模块
		if err := c.ParseForm(form); err != nil {
			json["text"] = err.Error()
		} else {
			if corret,err := valid.Valid(form);err != nil {
				json["text"] = err.Error()
			} else if !corret {
				json["result"] = valid.Errors
			} else {
				//填充task模型
				task := &models.Task{
					Name:         form.Name,
					Worker:       form.Worker,
					Desc:         form.Desc,
					CreateUser:   c.User.Id,   //我现在还没写登录模块，所以这里没有这个ID
				}

				ormer := orm.NewOrm()
				if _,err := ormer.Insert(task);err == nil {
					json = map[string]interface{}{
						"code": 200,
						"text": "创建成功",
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
	//刷新创建表单页面，是get方式
	c.Data["json"] = json
	c.ServeJSON()
}

func (c *TaskController) Detail(){
	json := map[string]interface{}{
		"code": 400,
		"text": "请求数据错误",
		"result": nil,
	}

	if id,err := c.GetInt("id");err == nil {
		task := models.Task{Id:id}
		//没有c.User的数据，这个数据是登录时填充的用户数据，现在没写登录模块
		if orm.NewOrm().Read(&task) == nil &&  (c.User.IsSuper || task.CreateUser == c.User.Id)  {
			json = map[string]interface{}{
				"code": 200,
				"text": "请求数据成功",
				"result": task,
			}
		}
	}
	c.Data["json"] = json
	c.ServeJSON()
}

func (c *TaskController) Modify(){
	json := map[string]interface{}{
		"code": 405,
		"text": "请求方式错误",
		"result": nil,
	}

	if c.Ctx.Input.IsPost() {
		json = map[string]interface{}{
			"code": 400,
			"text": "提交数据错误",
			"result": nil,
		}
		form := &forms.TaskModifyForm{User: c.User} //声明修改表单数据模型变量
		valid := &validation.Validation{} //加载验证器
		if err := c.ParseForm(form);err != nil {
			json["text"] = err.Error()
		} else {
			if corret,err := valid.Valid(form);err != nil {
				json["text"] = err.Error()
			} else if !corret {
				json["result"] = valid.Errors
			} else {
				form.Task.Name = form.Name
				form.Task.Progress = form.Progress
				form.Task.Status = form.Status
				form.Task.Worker = form.Worker
				form.Task.Desc = form.Desc

				if form.Status == models.TaskStatusComplete {
					now := time.Now()
					form.Task.CompleteTime = &now
					form.Task.Progress = 100
				}

				ormer := orm.NewOrm()
				if _,err := ormer.Update(form.Task);err == nil {
					json = map[string]interface{}{
						"code": 200,
						"text": "修改成功",
						"result": form.Task,
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
	c.Data["json"] = json
	c.ServeJSON()
}

func (c *TaskController) Delete() {
	if id,err := c.GetInt("id");err == nil {
		//超级管理员可以删所有，普通管理员只能删除自己创建的
		task := models.Task{Id:id}
		ormer := orm.NewOrm()
		if ormer.Read(&task) == nil && (c.User.IsSuper || task.CreateUser == c.User.Id) {
			ormer.Delete(&task)
		}
	}

	c.Data["json"] = map[string]interface{}{
		"code": 200,
		"text": "删除成功",
	}
	c.ServeJSON()
}