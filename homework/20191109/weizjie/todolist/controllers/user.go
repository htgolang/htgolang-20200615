package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"todolist/forms"
	"todolist/models"
	"todolist/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

// 用户控制器
type UserController struct {
	LoginRequiredController
}

func (this *UserController) Prepare() {
	this.LoginRequiredController.Prepare()
	_, action := this.GetControllerAndAction()

	// 除user/password 外其他url只有超级管理员可操作
	if action != "Password" && !this.User.IsSuper {
		// 非管理员且并非请求user/password则跳转到主页面
		this.Redirect(beego.URLFor(beego.AppConfig.String("home")), http.StatusFound)
	}

}

func (c *UserController) Index() {
	c.Layout = "layout/base.html"
	c.Data["nav"] = "user"
	c.LayoutSections = map[string]string{}
	c.LayoutSections["LayoutScripts"] = "user/index_scripts.html"
	c.TplName = "user/index.html"
	c.Data["xsrf_token"] = c.XSRFToken() // csrftoken
}

func (c *UserController) List_User() {
	orderByColumns := map[string]bool{
		"id":          true,
		"name":        true,
		"birthday":    true,
		"tel":         true,
		"addr":        true,
		"create_time": true,
	}

	draw := c.GetString("draw")
	start, err := c.GetInt("start")
	if err != nil {
		start = 0
	}
	length, err := c.GetInt("length")
	if err != nil {
		length = 10
	}

	q := strings.TrimSpace(c.GetString("q"))

	orderBy := c.GetString("orderBy")
	if _, ok := orderByColumns[orderBy]; !ok {
		orderBy = "id"
	}

	orderDir := c.GetString("orderDir")
	if orderDir == "desc" {
		orderBy = "-" + orderBy
	}

	var users []*models.User

	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&models.User{})
	total, _ := queryset.Count()

	// 创建查询条件
	condition := orm.NewCondition()

	if q != "" {
		// 创建查询条件
		// 查询名称和描述中包含字符
		condition = condition.Or("name__icontains", q)
		condition = condition.Or("tel__icontains", q)
		condition = condition.Or("addr__icontains", q)
		condition = condition.Or("desc__icontains", q)

	}
	// 查询数据
	totalFilter, _ := queryset.OrderBy(orderBy).Limit(length).Offset(start).SetCond(condition).All(&users)

	c.Data["json"] = map[string]interface{}{
		"code":            200,
		"text":            "获取用户成功",
		"result":          users,
		"draw":            draw,
		"recordsTotal":    total,
		"recordsFiltered": totalFilter,
	}
	c.ServeJSON()
}

func (c *UserController) Create() {
	json := map[string]interface{}{
		"code":   405,
		"text":   "请求方式错误",
		"result": nil,
	}

	if c.Ctx.Input.IsPost() {
		json = map[string]interface{}{
			"code":   400,
			"text":   "提交数据错误",
			"result": nil,
		}
		form := &forms.UserCreateForm{}
		valid := &validation.Validation{}

		if err := c.ParseForm(form); err != nil {
			json["text"] = err.Error()
		} else {
			// 表单验证
			if corret, err := valid.Valid(form); err != nil {
				json["text"] = err.Error()
			} else if !corret {
				json["result"] = valid.Errors
			} else {
				// 创建用户
				birthday, _ := time.Parse("2006-01-02", form.Birthday)

				user := &models.User{
					Name:     form.Name,
					Birthday: &birthday,
					Sex:      form.Sex == 1,
					Tel:      form.Tel,
					Addr:     form.Addr,
					Desc:     form.Desc,
				}
				user.SetPassword(form.Password)

				ormer := orm.NewOrm()
				if _, err := ormer.Insert(user); err == nil {
					json = map[string]interface{}{
						"code":   200,
						"text":   "创建成功",
						"result": user,
					}
				} else {
					json = map[string]interface{}{
						"code":   500,
						"text":   "服务端错误",
						"result": nil,
					}
				}
			}
		}
	}
	c.Data["json"] = json
	c.ServeJSON()
}

func (c *UserController) Modify() {
	json := map[string]interface{}{
		"code":   405,
		"text":   "请求方式错误",
		"result": nil,
	}

	form := &forms.UserModifyForm{} // 任务修改表单
	valid := &validation.Validation{}

	if c.Ctx.Input.IsGet() {
		// 获取任务信息（超级管理员可查看任意任务信息，普通管理员只能查看自己创建的任务信息）
		if id, err := c.GetInt("id"); err == nil {
			user := models.User{Id: id}
			if orm.NewOrm().Read(&user) == nil {
				json["code"] = 200
				json["text"] = "获取数据成功"
				json["result"] = user
			}
		}
		c.Data["json"] = json
		c.ServeJSON()
	} else if c.Ctx.Input.IsPost() {
		// 任务修改
		json = map[string]interface{}{
			"code":   400,
			"text":   "请求数据错误",
			"result": nil,
		}

		// 解析请求参数到form中(根据form标签)

		if err := c.ParseForm(form); err != nil {
			json["text"] = err.Error()
		} else {
			if corret, err := valid.Valid(form); err != nil {
				json["text"] = err.Error()
			} else if !corret {

				json["result"] = valid.Errors

			} else {
				// 验证任务信息（超级管理员可修改任意任务信息，普通管理员只能修改自己创建的任务信息）
				birthday, _ := time.Parse("2006-01-02", form.Birthday)
				form.User.Name = form.Name
				form.User.Birthday = &birthday
				form.User.Sex = form.Sex
				form.User.Tel = form.Tel
				form.User.Addr = form.Addr
				form.User.Desc = form.Desc

				// 更新数据库
				if _, err := orm.NewOrm().Update(form.User); err == nil {
					json["code"] = 200
					json["text"] = "修改成功"
					json["result"] = form.User

				} else {
					fmt.Println(err)
					json["code"] = 500
					json["text"] = "服务器错误"
					json["result"] = nil
				}
			}
		}
	}
	c.Data["json"] = json
	c.ServeJSON()

}

func (c *UserController) Delete() {
	json := map[string]interface{}{
		"code":   400,
		"text":   "请求数据错误",
		"result": nil,
	}

	if id, err := c.GetInt("id"); err == nil {
		ormer := orm.NewOrm()
		user := models.User{Id: id}
		if ormer.Read(&user) == nil {
			ormer.Delete(&user)
			c.Data["json"] = map[string]interface{}{
				"code":   200,
				"text":   "success",
				"result": nil,
			}
		}
	} else {
		json["text"] = err.Error()
	}
	c.ServeJSON()
}

func (c *UserController) Password() {
	json := map[string]interface{}{
		"code":   405,
		"text":   "请求方式错误",
		"result": nil,
	}
	form := &forms.ModifyPasswordForm{User: c.User}
	valid := &validation.Validation{}

	if c.Ctx.Input.IsPost() {
		json = map[string]interface{}{
			"code":   400,
			"text":   "请求数据错误",
			"result": nil,
		}
		if err := c.ParseForm(form); err != nil {
			json["text"] = err.Error()
		} else {
			if corret, err := valid.Valid(form); err != nil {
				json["text"] = err.Error()
			} else if !corret {
				json["result"] = valid.Errors
			} else {
				c.User.SetPassword(form.NewPassword)
				if result, err := orm.NewOrm().Update(c.User, "Password"); err == nil {
					json["code"], json["text"], json["result"] = 200, "密码修改成功", result
				} else {
					json["code"], json["text"], json["result"] = 500, "服务器端错误", err.Error()
				}
			}
		}
	}
	c.Data["json"] = json
	c.ServeJSON()
}

func (c *UserController) ResetPassword() {
	json := map[string]interface{}{
		"code":   400,
		"text1":  "请求数据错误",
		"result": nil,
	}
	if id, err := c.GetInt("id"); err != nil {
		json["text"] = err.Error()
	} else if c.User.IsSuper {

		user := &models.User{Id: id}
		o := orm.NewOrm()
		if err := o.Read(user); err == nil {
			password := utils.RandString(6)
			user.SetPassword(password)

			if result, err := o.Update(user, "Password"); err == nil {
				json["code"], json["text"], json["result"] = 200, fmt.Sprintln("重置用户密码为:", password), result
			} else {
				json["code"], json["text"], json["result"] = 500, "服务器端错误", err.Error()
			}
		} else {
			json["text"] = err.Error()
		}
	}
	c.Data["json"] = json
	c.ServeJSON()
}
