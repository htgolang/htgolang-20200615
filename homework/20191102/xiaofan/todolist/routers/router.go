package routers

import (
	"github.com/astaxie/beego"
	"todolist/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})

	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.TaskController{})
	beego.AutoRouter(&controllers.UserController{})

}
