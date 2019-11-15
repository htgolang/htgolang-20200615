package routers

import (
	"github.com/astaxie/beego"
	"github.com/imsilence/todolist/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{}) //定义根路径/与HomeController路由

	// 使用自动匹配路径定义登陆认证，任务，用户管理路由
	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.TaskController{})
	beego.AutoRouter(&controllers.UserController{})
}
