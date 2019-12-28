package routers

import (
	"github.com/astaxie/beego"
	"github.com/xlotz/gocmdb/server/controllers"
	"github.com/xlotz/gocmdb/server/controllers/auth"
	"github.com/xlotz/gocmdb/server/controllers/api/v2"
	"github.com/xlotz/gocmdb/server/controllers/api/v1"


)

func init() {
	// 认证

	beego.Router("/", &controllers.IndexController{}, "get:Index")
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
	// 终端管理
	beego.AutoRouter(&controllers.AgentPageController{})
	beego.AutoRouter(&controllers.AgentController{})
	beego.AutoRouter(&controllers.LogPageController{})
	beego.AutoRouter(&controllers.LogController{})
	beego.AutoRouter(&controllers.AlarmPageController{})
	beego.AutoRouter(&controllers.AlarmController{})

	v1Namespace := beego.NewNamespace("/v1/",
		beego.NSRouter("api/heartbeat/:uuid/", &v1.APIController{}, "*:Heartbeat"),
		beego.NSRouter("api/register/:uuid/", &v1.APIController{}, "*:Register"),
		beego.NSRouter("api/log/:uuid/", &v1.APIController{}, "*:Log"),
		)
	beego.AddNamespace(v1Namespace)

	v2Namespace := beego.NewNamespace("/v2/",
		beego.NSRouter("api/heartbeat/:uuid/", &v2.APIController{}, "*:Heartbeat"),
		beego.NSRouter("api/register/:uuid/", &v2.APIController{}, "*:Register"),
		beego.NSRouter("api/log/:uuid/", &v2.APIController{}, "*:Log"),

	)
	beego.AddNamespace(v2Namespace)
}
