package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"net/http"
	"todolist/forms"
)

type AuthController struct {
	BaseController
}

// login页面的验证
func (c *AuthController) Login() {
	c.ValidateSession()
	// session验证通过就跳转到主页
	if c.User != nil {
		c.Redirect(beego.URLFor(beego.AppConfig.String("home")), http.StatusFound)
	} else {
		// 认证未通过则返回登录页面
		c.TplName = "auth/index.html"
		c.Data["form"] = forms.AuthForm{}
		c.Data["validation"] = &validation.Validation{}
	}
}

// 提交登录验证后的路由
func (c *AuthController) Auth() {
	form := &forms.AuthForm{}
	valid := &validation.Validation{}
	// 如果Parse提交的表单成功
	if err := c.ParseForm(form); err == nil {
		// 使用Valid对form进行校验，Valid是我们自定义的方法.
		if correct, err := valid.Valid(form); err == nil && correct {
			// 如果登录校验成功，设置session并跳转到主页
			c.SetSession("user", form.User.Id)
			c.Redirect(beego.URLFor(beego.AppConfig.String("home")), http.StatusFound)
		}
	}

	// 验证失败则返回登录页面的内容
	c.TplName = "auth/index.html"
	c.Data["form"] = forms.AuthForm{}
	c.Data["validation"] = valid
}

// 登出
func (c *AuthController) Logout() {
	c.ValidateSession()
	if c.User != nil {
		c.DestroySession()
	}
	c.Redirect(beego.URLFor(beego.AppConfig.String("login")), http.StatusFound)
}
