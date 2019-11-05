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

func (this *BaseController) ValidateSession() {
	// 检查是否含有user这个session
	if session := this.GetSession("user"); session != nil {
		// 检查user的值是否是int类型,session的值是在数据库中的id
		if id, ok := session.(int); ok {
			user := models.User{Id: id}
			// 检查session的值是否正确
			if err := models.GetUser(&user); err == nil {
				// 将user的所有信息赋值给当前路由
				this.User = &user
				this.Data["user"] = &user
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

func (this *LoginRequiredController) Prepare() {
	this.ValidateSession()
	// session未在数据库中匹配上，则跳转至登录界面
	if this.User == nil {
		this.Redirect(beego.URLFor(beego.AppConfig.String("login")), http.StatusFound)
	} else {
		beego.ReadFromRequest(&this.Controller)
	}
}
