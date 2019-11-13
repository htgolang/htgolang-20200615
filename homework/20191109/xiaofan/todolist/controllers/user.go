package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"net/http"
	"strings"
	"time"
	"todolist/forms"
	"todolist/models"
	"todolist/utils"
)

// 用户控制器
type UserController struct {
	LoginRequiredController
}

func (c *UserController) Prepare() {
	c.LoginRequiredController.Prepare()
	_, action := c.GetControllerAndAction()

	// 除user/password 外其他url只有超级管理员可操作
	if action != "Password" && !c.User.IsSuper {
		// 非管理员且并非请求user/password则跳转到主页面
		c.Redirect(beego.URLFor(beego.AppConfig.String("home")), http.StatusFound)
	}

}

func (c *UserController) Index() {
	c.Data["nav"] = "user"
	c.Layout = "layout/base.html"
	c.TplName = "user/index.html"
	c.LayoutSections = map[string]string{}
	c.LayoutSections["LayoutScripts"] = "user/index_scripts.html"
}

func (c *UserController) List() {
	orderByColumns := map[string]bool{
		"name":     true,
		"birthday": true,
		"tel":      true,
		"addr":     true,
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

	// 获取查询字段
	q := strings.TrimSpace(c.GetString("q"))

	// 查询数据总数
	queryset := orm.NewOrm().QueryTable(&models.User{})
	total, _ := queryset.Count()

	// 创建条件并判断是否将添加加入进去
	condition := orm.NewCondition()
	if q != "" {
		condition = condition.Or("name__icontains", q)
		condition = condition.Or("tel__icontains", q)
		condition = condition.Or("addr__icontains", q)
		condition = condition.Or("desc__icontains", q)
	}
	// 查询数据并返回数据量
	var users []*models.User
	totalFilter, _ := queryset.OrderBy(orderBy).Limit(length).Offset(start).SetCond(condition).All(&users)
	c.Data["json"] = map[string]interface{}{
		"code":            200,
		"text":            "获取任务成功",
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
	form := &forms.UserCreateForm{}
	valid := &validation.Validation{}

	if c.Ctx.Input.IsPost() {
		json["code"], json["text"] = 400, "请求数据错误"
		if err := c.ParseForm(form); err != nil {
			json["text"] = err.Error()
		} else {
			if corret, err := valid.Valid(form); err != nil {
				json["text"] = err.Error()
			} else if !corret {
				json["result"] = valid.Errors
			} else {
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
				if result, err := orm.NewOrm().Insert(user); err == nil {
					json["code"], json["text"], json["result"] = 200, "创建成功", result
				} else {
					json["code"], json["text"], json["result"] = 500, "服务器端错误", err.Error()
				}
			}
		}
	}

	c.Data["json"] = json
	c.ServeJSON()
}

func (c *UserController) Detail() {
	json := map[string]interface{}{
		"code":   400,
		"text":   "请求数据错误",
		"result": nil,
	}

	if id, err := c.GetInt("id"); err == nil {
		user := &models.User{Id: id}
		if orm.NewOrm().Read(user) == nil && (c.User.IsSuper || user.Id == c.User.Id) {
			json["code"], json["text"], json["result"] = 200, "获取数据成功", user
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
	form := &forms.UserModifyForm{}
	valid := &validation.Validation{}

	if c.Ctx.Input.IsPost() {
		json["code"], json["text"] = 400, "请求数据错误"
		if err := c.ParseForm(form); err != nil {
			json["text"] = err.Error()
		} else {
			if corret, err := valid.Valid(form); err != nil {
				json["text"] = err.Error()
			} else if !corret {
				json["result"] = valid.Errors
			} else {
				birthday, _ := time.Parse("2016-01-02", form.Birthday)
				form.User.Name = form.Name
				form.User.Birthday = &birthday
				form.User.Sex = form.Sex == 1
				form.User.Tel = form.Tel
				form.User.Addr = form.Addr
				form.User.Desc = form.Desc

				// 更新数据库
				if result, err := orm.NewOrm().Update(form.User); err == nil {
					json["code"], json["text"], json["result"] = 200, "修改用户成功", result
				} else {
					json["code"], json["text"], json["result"] = 500, "服务器端错误", err.Error()
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
	if id, err := c.GetInt("id"); err != nil {
		json["text"] = err.Error()
	} else {
		if result, err := orm.NewOrm().Delete(&models.User{Id: id}); err == nil {
			json["code"], json["text"], json["result"] = 200, "用户删除成功", result
		} else {
			json["code"], json["text"], json["result"] = 500, "服务器端错误", err.Error()
		}
	}

	c.Data["json"] = json
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
		json["code"], json["text"] = 400, "请求数据错误"
		if err := c.ParseForm(form); err != nil {
			json["text"] = err.Error()
		} else {
			if corret, err := valid.Valid(form); err != nil {
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
		"text":   "请求数据错误",
		"result": nil,
	}
	if id, err := c.GetInt("id"); err != nil {
		json["text"] = err.Error()
	} else if c.User.IsSuper {

		user := &models.User{Id: id}
		o := orm.NewOrm()
		if err := o.Read(user); err == nil {
			password := utils.RandomString(6)
			user.SetPassword(password)

			if result, err := o.Update(user, "Password"); err == nil {
				json["code"], json["text"], json["result"] = 200, fmt.Sprintf("重置用户: %s,密码为: %s", user.Name, password), result
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
