package routers

import (
	"github.com/astaxie/beego"

	"github.com/imsilence/gocmdb/server/controllers/auth"
	"github.com/imsilence/gocmdb/server/controllers"
)

func init() {
	// 认证
	beego.AutoRouter(&auth.AuthController{})

	// 用户页面
	beego.AutoRouter(&controllers.UserPageController{})

	// 用户
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.TokenController{})


	// 云平台页面
	beego.AutoRouter(&controllers.CloudPlatformPageController{})

	// 云平台
	beego.AutoRouter(&controllers.CloudPlatformController{})

	// 云主机页面
	beego.AutoRouter(&controllers.VirtualMachinePageController{})
	// 云主机
	beego.AutoRouter(&controllers.VirtualMachineController{})
}
