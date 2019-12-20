package routers

import (
	"github.com/astaxie/beego"

	"github.com/imsilence/gocmdb/server/controllers"
	v1 "github.com/imsilence/gocmdb/server/controllers/api/v1"
	"github.com/imsilence/gocmdb/server/controllers/auth"
)

func init() {
	// 默认路由
	beego.Router("/", &controllers.IndexController{}, "get:Index")

	// 测试
	beego.AutoRouter(&controllers.TestController{})

	beego.AutoRouter(&auth.AuthController{})
	beego.AutoRouter(&controllers.UserPageController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.TokenController{})
	// 云平台管理
	beego.AutoRouter(&controllers.CloudPlatformPageController{})
	beego.AutoRouter(&controllers.CloudPlatformController{})

	beego.AutoRouter(&controllers.VirtualMachinePageController{})
	beego.AutoRouter(&controllers.VirtualMachineController{})

	beego.AutoRouter(&controllers.LogPageController{})
	beego.AutoRouter(&controllers.LogController{})

	beego.AutoRouter(&controllers.AgentPageController{})
	beego.AutoRouter(&controllers.AgentController{})

	v1Namespace := beego.NewNamespace("/v1",
		beego.NSRouter("api/heartbeat/:uuid/", &v1.APIController{}, "*:Heartbeat"),
		beego.NSRouter("api/register/:uuid/", &v1.APIController{}, "*:Register"),
		beego.NSRouter("api/log/:uuid/", &v1.APIController{}, "*:Log"),
	)
	beego.AddNamespace(v1Namespace)
}
