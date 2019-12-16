package routers

import (
	"github.com/astaxie/beego"
	"github.com/xxdu521/cmdbgo/server/controllers"
	v1 "github.com/xxdu521/cmdbgo/server/controllers/api/v1"
	"github.com/xxdu521/cmdbgo/server/controllers/auth"
)

func init(){
	//默认页跳转
	beego.Router("/", &controllers.IndexController{}, "get:Index")

	//认真（登录/退出）
	beego.AutoRouter(&auth.AuthController{})

	//用户管理
	beego.AutoRouter(&controllers.UserPageController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.TokenController{})  //特殊访问，单独处理，包含在User逻辑下

	//云管平台（云平台/云主机）
	beego.AutoRouter(&controllers.CloudPlatformPageController{})
	beego.AutoRouter(&controllers.CloudPlatformController{})
	beego.AutoRouter(&controllers.VirtualMachinePageController{})
	beego.AutoRouter(&controllers.VirtualMachineController{})

	//api（heartbeat/register/log）
	v1Namespace := beego.NewNamespace("/v1",
		beego.NSRouter("api/heartbeat/:uuid/", &v1.APIController{}, "*:Heartbeat"),
		beego.NSRouter("api/register/:uuid/", &v1.APIController{}, "*:Register"),
		beego.NSRouter("api/log/:uuid/", &v1.APIController{}, "*:Log"),
		)
	beego.AddNamespace(v1Namespace)


}
