package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
	//"todolist/models"
	//"github.com/imslience/todolist/form"
	"github.com/astaxie/beego/validation"
	"todolist/forms"
)

type AuthController struct {
	BaseController
}

func (this *AuthController) Login() {
	this.ValidateSession()
	if this.User != nil {
		this.Redirect("/", http.StatusFound)
	} else {
		this.TplName = "auth/index.html"
		this.Data["form"] = forms.AuthForm{}
		this.Data["validation"] = &validation.Validation{}
	}
}

func (this *AuthController) Auth() {
	form := &forms.AuthForm{}
	valid := &validation.Validation{}

	if err := this.ParseForm(form); err == nil {
		if correct, err := valid.Valid(form); err == nil && correct {
			this.SetSession("user", form.User.Id)
			this.Redirect("/", http.StatusFound)
		}
	}

	this.TplName = "auth/index.html"
	this.Data["form"] = form
	this.Data["validation"] = valid
}

// 登出逻辑
func (this *AuthController) Logout() {
	this.DestroySession() // 销毁session
	this.Redirect(beego.URLFor(beego.AppConfig.String("logout")), http.StatusFound) // 跳转到登陆页面
}