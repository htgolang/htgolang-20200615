package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"net/http"
	"todolist/models"
)

type BaseController struct {
	beego.Controller
	User *models.User
}

func (c *BaseController) ValidateSession() {
	// 检查是否含有user这个session
	if session := c.GetSession("user"); session != nil {
		// 检查user的值是否是int类型,session的值是在数据库中的id
		if id, ok := session.(int); ok {
			user := models.User{Id: id}
			// 检查session的值是否正确
			if err := models.GetUser(&user); err == nil {
				// 将user的所有信息赋值给当前路由
				c.User = &user
				c.Data["user"] = &user
			} else {
				fmt.Printf("Can not get user for id %d, error: %s\n", id, err.Error())
			}
		} else {
			fmt.Printf("Invalid session id %v\n", session)
		}
	} else {
		fmt.Println("Not login")
	}
}

type LoginRequiredController struct {
	BaseController
}

func (c *LoginRequiredController) Prepare() {
	c.ValidateSession()
	// session未在数据库中匹配上，则跳转至登录界面
	if c.User == nil {
		if c.Ctx.Input.IsAjax() {
			c.Data["json"] = map[string]interface{}{
				"code": 403,
				"text": "未认证",
			}
			c.ServeJSON()
		}
		c.Redirect(beego.URLFor(beego.AppConfig.String("login")), http.StatusFound)
	} else {
		beego.ReadFromRequest(&c.Controller)
	}
}
