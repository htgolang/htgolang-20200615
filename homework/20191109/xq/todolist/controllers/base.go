package controllers

import (
	"net/http"

	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"

	"github.com/imsilence/todolist/models"
)

// 控制器基类
type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {

}

// 控制器基类，用于在登陆之后才可访问的控制器（组合LoginRequiredController）
type LoginRequiredController struct {
	BaseController
	User *models.User
}

func (c *LoginRequiredController) Prepare() {
	c.BaseController.Prepare()
	// 检查用户是否已经登陆（在session中可获取正确的用户id）
	if session := c.GetSession("user"); session != nil {
		if id, ok := session.(int); ok {
			user := models.User{Id: id}
			ormer := orm.NewOrm()
			if ormer.Read(&user) == nil {
				beego.ReadFromRequest(&c.Controller) // 从request中读取flash cookie数据（flash）
				c.User = &user                       // 登陆用户信息
				c.Data["user"] = &user
			}
		}
	}
	if c.User == nil {
		// 用户未登陆则跳转到登陆控制器，并使用next参数传递请求的url
		c.Redirect(beego.URLFor(beego.AppConfig.String("login"), "next", c.Ctx.Input.URL()), http.StatusFound)
	}
}
