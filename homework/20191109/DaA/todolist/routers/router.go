package routers

import (
	"todolist/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/",&controllers.HomeController{}) //定义主页

    beego.AutoRouter(&controllers.AuthController{})

	beego.AutoRouter(&controllers.TaskController{})
    beego.AutoRouter(&controllers.UserController{})

	beego.AutoRouter(&controllers.LoginController{})
}
