package controllers

import (

	"net/http"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"

	"github.com/imsilence/todolist/models"
	"github.com/imsilence/todolist/forms"
	"github.com/imsilence/todolist/utils"
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

	c.Layout = "layout/base.html" //设置layout
	c.Data["nav"] = "user" //设置菜单
}


// 用户列表
func (c *UserController) Index() {
	q := strings.TrimSpace(c.GetString("q"))

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

	c.Layout = "layout/base.html" // 设置layout
	c.Data["nav"] = "user" // 设置菜单

	c.LayoutSections = map[string]string{}
	c.LayoutSections["LayoutScripts"] = "user/index_scripts.html"
	c.TplName = "user/index.html"
	c.Data["xsrf_token"] = c.XSRFToken() // csrftoken
	c.Data["q"] = q
	c.Data["users"] = users

}

// 用户创建
func (c *UserController) Create() {

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

		form := &forms.UserCreateForm{} // 用户创建表单
		valid := &validation.Validation{} //验证器


		// 解析请求参数到form中(根据form标签)
		if err := c.ParseForm(form); err != nil {
			json["text"] = err.Error()

		}else {

			// 表单验证
			if corret, err := valid.Valid(form); err != nil {

				json["text"] = err.Error()
			}else if !corret {
				json["result"] = valid.Errors

			}else {

				// 转换时间
				birthday, _ := time.Parse("2006-01-02", form.Birthday)

				// 创建结构体对象
				user := &models.User{
					Name: form.Name,
					Birthday: &birthday,
					Sex: form.Sex == 1,
					Tel: form.Tel,
					Addr: form.Addr,
					Desc: form.Desc,
				}

				//设置密码
				user.SetPassword(form.Password)

				// 插入用户
				//ormer := orm.NewOrm()
				//ormer.Insert(user)

				ormer := orm.NewOrm()

				if _, err := ormer.Insert(user); err == nil {

					json = map[string]interface{}{
						"code": 200,
						"text": "创建成功",
						"result": user,
					}
				}else {
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

// 用户修改
func (c *UserController) Modify() {

	json := map[string]interface{}{
		"code": 405,
		"text": "请求方式错误",
		"result": nil,
	}

	form := &forms.UserModifyForm{} // 修改表单
	valid := &validation.Validation{} //验证器


	if c.Ctx.Input.IsGet() {

		// 获取用户信息(显示)
		if id, err := c.GetInt("id"); err != nil {
			json["code"] = 400
			json["text"] = "获取ID失败"
			json["result"] = err.Error()

		}else{

			user := models.User{Id: id}

			if err := orm.NewOrm().Read(&user);err != nil {
				json["code"] = 400
				json["text"] = "获取数据失败"
				json["result"] = err.Error()

			}else {

				json["code"] = 200
				json["text"] = "获取数据成功"
				json["result"] = user
			}

		}

		c.Data["json"] = json
		c.ServeJSON()

	} else if c.Ctx.Input.IsPost() {
		// 解析请求参数到form中(根据form标签)
		if err := c.ParseForm(form); err != nil {
			json["code"] = 500
			json["text"] = "解析失败"
			json["result"] = err.Error()

		}else {
			// 验证表单数据
			if corret, err := valid.Valid(form); err != nil {
				json["code"] = 400
				json["result"] = err.Error()

			}else if !corret {
				json["code"] = 400

			}else {

				// 更新用户信息
				birthday, _ := time.Parse("2006-01-02", form.Birthday)

				form.User.Name = form.Name
				form.User.Birthday = &birthday
				form.User.Sex = form.Sex == 1
				form.User.Tel = form.Tel
				form.User.Addr = form.Addr
				form.User.Desc = form.Desc


				if _, err := orm.NewOrm().Update(form.User); err == nil {
					json["code"] = 200
					json["text"] = "修改成功"
					json["result"] = form.User

				}else {

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

// 用户删除
func (c *UserController) Delete() {

	json := map[string]interface{}{
		"code":   405,
		"text":   "请求方式错误",
		"result": nil,
	}
	if c.Ctx.Input.IsPost() {
		if id, err := c.GetInt("id"); err != nil {

			json["code"] = 400
			json["text"] = "获取ID失败"
			json["result"] = err.Error()

		}else{
			// 删除用户
			orm.NewOrm().Delete(&models.User{Id: id})

			// 通过flash提示用户操作结果
			json["code"] = 200
			json["text"] = "删除成功"
			json["result"] = nil

		}

		}

		c.Data["json"] = json
		c.ServeJSON()
	}

// 修改密码
func (c *UserController) Password() {

	json := map[string]interface{}{
		"code":   405,
		"text":   "请求方式错误",
		"result": nil,
	}

	form := &forms.ModifyPasswordForm{User: c.User} //修改密码表单
	valid := &validation.Validation{} // 验证器

	if c.Ctx.Input.IsPost() {

		// 解析请求参数到form中(根据form标签)
		if err := c.ParseForm(form) ; err != nil {

			json["code"] = 400
			json["text"] = "解析数据错误"
			json["result"] = err.Error()

		}else {

			// 表单验证
			if corret, err := valid.Valid(form); err != nil {
				json["code"] = 400
				json["text"] = "表单验证失败"
				json["result"] = err.Error()


			}else if !corret {

				json["code"] = 400
				json["result"] = nil

			}else{

				// 设置密码
				c.User.SetPassword(form.NewPassword)

				ormer := orm.NewOrm()

				// 只更新密码列
				ormer.Update(c.User, "Password")

				json["code"] = 200
				json["text"] = "修改密码成功"
				json["result"] = nil


			}
		}
	}
	c.Data["json"] = json
	c.ServeJSON()
}

//重置密码
func (c *UserController) ResetPassword() {

	json := map[string]interface{}{
		"code":   405,
		"text":   "请求方式错误",
		"result": nil,
	}

	if c.Ctx.Input.IsPost(){

		if id, err := c.GetInt("id"); err != nil{

			json["code"] = 400
			json["text"] = "获取ID失败"
			json["result"] = err.Error()

		} else if !c.User.IsSuper {
			json["code"] = 405
			json["text"] = "没有权限"
			json["result"] = err.Error()

		}else {

			user := models.User{Id: id}
			ormer := orm.NewOrm()
			if err := ormer.Read(&user); err == nil {
				password := utils.RandString(6) // 生成随机6位密码
				user.SetPassword(password) // 设置密码
				ormer.Update(&user, "Password") //更新数据库

				json["code"] = 200
				json["text"] = "重置密码成功"
				json["result"] = password

			}
		}

	}
	c.Data["json"] = json
	c.ServeJSON()
}
