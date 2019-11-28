package routers

import (
	"github.com/astaxie/beego"
	"gocmdb/controllers"
	"gocmdb/controllers/auth"
)

func init() {
	// login/
	beego.AutoRouter(&auth.AuthController{})
	// user/
	beego.AutoRouter(&controllers.UserController{})
	// userpage/
	beego.AutoRouter(&controllers.UserPageController{})
	// token/
	beego.AutoRouter(&controllers.TokenController{})
}
