package routers

import (
	"github.com/astaxie/beego"
	"github.com/xxdu521/cmdbgo/server/controllers"
	"github.com/xxdu521/cmdbgo/server/controllers/auth"
)

func init(){
	beego.AutoRouter(&auth.AuthController{})

	//用户管理
	beego.AutoRouter(&controllers.UserPageController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.TokenController{})
	//云管平台
	beego.AutoRouter(&controllers.CloudPlatformPageController{})
	beego.AutoRouter(&controllers.CloudPlatformController{})
	beego.AutoRouter(&controllers.VirtualMachinePageController{})
	beego.AutoRouter(&controllers.VirtualMachineController{})
}
