package routers

import (
	"github.com/xlotz/gocmdb/server/controllers"
	"github.com/xlotz/gocmdb/server/controllers/auth"
	"github.com/astaxie/beego"

)

func init() {

	beego.AutoRouter(&auth.AuthController{})
	//beego.AutoRouter(&controllers.TestController{})
	//beego.AutoRouter(&controllers.TestPageController{})
    beego.AutoRouter(&controllers.UserPageController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.TokenController{})
}
