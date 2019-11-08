package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"net/http"
	"strings"
	"time"
	"todolist/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"todolist/forms"
	"todolist/utils"
)

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

// 用户列表
func (this *UserController) Index() {
	q := strings.TrimSpace(this.GetString("q"))

	var users []models.User

	// 创建查询条件
	condition := orm.NewCondition()
	if q != "" {
		// 查询名称、电话、地址、描述中包含字符
		condition = condition.Or("name__icontains", q)
		condition = condition.Or("tel__icontains", q)
		condition = condition.Or("addr__icontains", q)
		condition = condition.Or("desc__icontains", q)
	}

	orm.NewOrm().QueryTable(&models.User{}).SetCond(condition).All(&users)

	this.Layout = "layout/base.html"
	this.Data["nav"] = "user"
	this.LayoutSections = map[string]string{}
	this.LayoutSections["LayoutScripts"] = "user/index_scripts.html"
	this.Data["q"] = q
	this.Data["users"] = users
	this.TplName = "user/index.html"
}

// 用户创建
func (this *UserController) Create() {
	json := map[string]interface{}{
		"code":   405,
		"text":   "请求方式错误",
		"result": nil,
	}

	form := &forms.UserCreateForm{} // 用户创建表单
	valid := &validation.Validation{} // 验证器

	if this.Ctx.Input.IsPost() {
		json["code"], json["text"] = 400, "请求数据错误"

		if err := this.ParseForm(form); err != nil {
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

	this.Data["json"] = json
	this.ServeJSON()
}

func (this *UserController) Detail() {
	json := map[string]interface{}{
		"code":   400,
		"text":   "请求数据错误",
		"result": nil,
	}

	if id, err := this.GetInt("id"); err == nil {

		user := &models.User{Id: id}
		if orm.NewOrm().Read(user) == nil && (this.User.IsSuper || user.Id == this.User.Id) {
			json["code"], json["text"], json["result"] = 200, "获取数据成功", user
		}
	}
	this.Data["json"] = json
	this.ServeJSON()
}

func (this *UserController) Modify() {
	json := map[string]interface{}{
		"code":   405,
		"text":   "请求方式错误",
		"result": nil,
	}
	form := &forms.UserModifyForm{}
	valid := &validation.Validation{}

	if this.Ctx.Input.IsPost() {
		json["code"], json["text"] = 400, "请求数据错误"
		if err := this.ParseForm(form); err != nil {
			json["text"] = err.Error()
		} else {
			if corret, err := valid.Valid(form); err != nil {
				json["text"] = err.Error()
			} else if !corret {
				json["result"] = valid.Errors
			} else {
				birthday, _ := time.Parse("2006-01-02", form.Birthday)
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

	this.Data["json"] = json
	this.ServeJSON()
}


// 用户删除
func (this *UserController) Delete() {
	json := map[string]interface{}{
		"code":   400,
		"text":   "请求数据错误",
		"result": nil,
	}

	if id, err := this.GetInt("id"); err != nil {
		json["text"] = err.Error()
	} else {
		// 删除用户
		if result, err := orm.NewOrm().Delete(&models.User{Id: id}); err == nil {
			json["code"], json["text"], json["result"] = 200, "用户删除成功", result
		} else {
			json["code"], json["text"], json["result"] = 500, "服务器端错误", err.Error()
		}
	}
	this.Data["json"] = json
	this.ServeJSON()
}

// 修改密码
func (this *UserController) Password() {
	json := map[string]interface{}{
		"code":   405,
		"text":   "请求方式错误",
		"result": nil,
	}
	form := &forms.ModifyPasswordForm{User: this.User}
	valid := &validation.Validation{}

	if this.Ctx.Input.IsPost() {
		json["code"], json["text"] = 400, "请求数据错误"
		if err := this.ParseForm(form); err != nil {
			json["text"] = err.Error()
		} else {
			if corret, err := valid.Valid(form); err != nil {
			} else if !corret {
				json["result"] = valid.Errors
			} else {
				this.User.SetPassword(form.NewPassword)
				if result, err := orm.NewOrm().Update(this.User, "Password"); err == nil {
					json["code"], json["text"], json["result"] = 200, "密码修改成功", result
				} else {
					json["code"], json["text"], json["result"] = 500, "服务器端错误", err.Error()
				}
			}
		}
	}
	this.Data["json"] = json
	this.ServeJSON()
}


//重置密码
func (this *UserController) ResetPassword() {
	json := map[string]interface{}{
		"code":   400,
		"text":   "请求数据错误",
		"result": nil,
	}
	if id, err := this.GetInt("id"); err != nil {
		json["text"] = err.Error()
	} else if this.User.IsSuper {

		user := &models.User{Id: id}
		o := orm.NewOrm()
		if err := o.Read(user); err == nil {
			password := utils.RandomString(6)
			user.SetPassword(password)

			if result, err := o.Update(user, "Password"); err == nil {
				json["code"], json["text"], json["result"] = 200, fmt.Sprintf("重置用户<%s>密码为<%s>", user.Name, password), result
			} else {
				json["code"], json["text"], json["result"] = 500, "服务器端错误", err.Error()
			}
		} else {
			json["text"] = err.Error()
		}
	}
	this.Data["json"] = json
	this.ServeJSON()
}