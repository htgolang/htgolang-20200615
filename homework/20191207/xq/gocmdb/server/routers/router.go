package routers

import (
	"github.com/astaxie/beego"

	"github.com/xlotz/gocmdb/server/controllers/auth"
	"github.com/xlotz/gocmdb/server/controllers"
)

func init() {
	// 认证
	beego.AutoRouter(&auth.AuthController{})
	// 用户页面
	beego.AutoRouter(&controllers.UserPageController{})
	// 用户管理
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.TokenController{})
	// 云平台管理
	beego.AutoRouter(&controllers.CloudPlatformPageController{})
	beego.AutoRouter(&controllers.CloudPlatformController{})
	//beego.AutoRouter(&controllers.VirtualMachinePageController())
	beego.AutoRouter(&controllers.VirtualMachineformController{})
}
