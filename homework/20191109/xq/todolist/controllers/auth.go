package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/imsilence/todolist/forms"
	"github.com/imsilence/todolist/models"
)

// 认证控制器
type AuthController struct {
	BaseController
}

func (c *AuthController) Prepare() {
	_, action := c.GetControllerAndAction()
	if action == "Login" {
		// 当访问auth/login时检查是否已经登陆，若已登陆则跳转到主页面
		if session := c.GetSession("user"); session != nil {
			if id, ok := session.(int); ok {
				user := models.User{Id: id}
				ormer := orm.NewOrm()
				if ormer.Read(&user) == nil {
					// 跳转到主页面
					c.Redirect(beego.URLFor(beego.AppConfig.String("home")), http.StatusFound)
				}
			}
		}

	}
}

// 登陆逻辑
func (c *AuthController) Login() {
	form := &forms.AuthForm{Next: c.GetString("next")} //登陆form
	valid := &validation.Validation{} //验证器

	if c.Ctx.Input.IsPost() {

		// 解析请求参数到form中(根据form标签)
		if err := c.ParseForm(form); err == nil {
			// 登陆表单验证
			if correct, err := valid.Valid(form); err == nil && correct {
				//用户名密码验证成功，存储session信息（标记已登陆）
				c.SetSession("user", form.User.Id)

				// 若url中传递next参数则跳转到next中，否则跳转到主页面
				c.Redirect(c.GetString("next", beego.URLFor(beego.AppConfig.String("home"))), http.StatusFound)
			}
		}
	}

	c.TplName = "auth/index.html"
	c.Data["xsrf_token"] = c.XSRFToken() //生成csrftoken，在beego框架层拦截csrf攻击
	c.Data["form"] = form
	c.Data["validation"] = valid
}

// 登出逻辑
func (c *AuthController) Logout() {
	c.DestroySession() // 销毁session
	c.Redirect(beego.URLFor(beego.AppConfig.String("logout")), http.StatusFound) // 跳转到登陆页面
}
