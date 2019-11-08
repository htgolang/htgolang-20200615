package routers

import (
	"github.com/astaxie/beego"
	"todolist/controllers"
	)

func init() {
	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.TaskController{})
	beego.Router("/", &controllers.TaskController{}, "get:Index")
	beego.AutoRouter(&controllers.UserController{})
}
