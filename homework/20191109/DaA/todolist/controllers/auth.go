package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"net/http"
	"todolist/forms"
	"todolist/models"
)

type AuthController struct {
	BaseController
}

func (c *AuthController) Prepare(){
	//可以判断用户请求的控制器方法
	_, action := c.GetControllerAndAction()
	if action == "Login" {
		//当访问auth/login时，检查是否已经登录，若已经登录则跳转到主页面
		if session := c.GetSession("user"); session != nil {
			if id, ok := session.(int); ok {
				user := models.User{Id:id}
				ormer := orm.NewOrm()
				if ormer.Read(&user) == nil {  //read方法，默认使用第一个int类型读取数据
					//确认用户成功，跳转到主页面
					c.Redirect(beego.URLFor(beego.AppConfig.String("home")),http.StatusFound)
				}
			}
		}
	}
}

func (c *AuthController) Login() {
	//先从header中获取next参数
	form := &forms.AuthForm{Next: c.GetString("next")}
	valid := &validation.Validation{} //加载验证器

	if c.Ctx.Input.IsPost() {
		//解析参数到form结构体中
		if err := c.ParseForm(form);err == nil {
			if correct, err := valid.Valid(form);err == nil && correct {
				//用户名密码验证成功，存储session信息（标记已登录）
				c.SetSession("user",form.User.Id)

				//若url中传递next参数则跳转到next中，否则跳转到主页面
				c.Redirect(c.GetString("next",beego.URLFor(beego.AppConfig.String("home"))),http.StatusFound)
			}
		}
	}

	c.TplName = "auth/index.html"
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["form"] = form
	c.Data["validation"] = valid
}

func (c *AuthController) Logout() {
	c.DestroySession() // 销毁session
	c.Redirect(beego.URLFor(beego.AppConfig.String("logout")), http.StatusFound) // 跳转到登陆页面
}