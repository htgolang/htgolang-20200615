package routers

import (
	"github.com/astaxie/beego"
	"github.com/dcosapp/gocmdb/server/controllers"
	v1 "github.com/dcosapp/gocmdb/server/controllers/api/v1"
	v2 "github.com/dcosapp/gocmdb/server/controllers/api/v2"
	"github.com/dcosapp/gocmdb/server/controllers/auth"
)

func init() {
	// /
	beego.Router("/", &controllers.HomePageController{})
	// homepage/
	beego.AutoRouter(&controllers.HomePageController{})
	// /dashboardpage/
	beego.AutoRouter(&controllers.DashboardPageController{})
	// login/
	beego.AutoRouter(&auth.AuthController{})
	// user/
	beego.AutoRouter(&controllers.UserController{})
	// userpage/
	beego.AutoRouter(&controllers.UserPageController{})
	// token/
	beego.AutoRouter(&controllers.TokenController{})
	// cloudplatform/
	beego.AutoRouter(&controllers.CloudPlatformController{})
	// cloudplatformpage/
	beego.AutoRouter(&controllers.CloudPlatformPageController{})
	// virtualmachinepage/
	beego.AutoRouter(&controllers.VirtualMachinePageController{})
	// virtualmachine/
	beego.AutoRouter(&controllers.VirtualMachineController{})
	// logpage/
	beego.AutoRouter(&controllers.LogPageController{})
	// log/
	beego.AutoRouter(&controllers.LogController{})
	// resourcepage/
	beego.AutoRouter(&controllers.ResourcePageController{})
	// resource/
	beego.AutoRouter(&controllers.ResourceController{})
	// alarmpage/
	beego.AutoRouter(&controllers.AlarmPageController{})
	// alarm/
	beego.AutoRouter(&controllers.AlarmController{})
	// alarmsettingpage/
	beego.AutoRouter(&controllers.AlarmSettingPageController{})
	// alarmsetting/
	beego.AutoRouter(&controllers.AlarmSettingController{})

	// v1/api/{type}/{uuid}
	v1Namespace := beego.NewNamespace("/v1",
		beego.NSRouter("api/heartbeat/:uuid/", &v1.APIController{}, "*:Heartbeat"),
		beego.NSRouter("api/register/:uuid/", &v1.APIController{}, "*:Register"),
		beego.NSRouter("api/log/:uuid/", &v1.APIController{}, "*:Log"),
	)

	// v2/api/{type}/{uuid}
	v2Namespace := beego.NewNamespace("/v2",
		beego.NSRouter("api/heartbeat/:uuid/", &v2.APIController{}, "*:Heartbeat"),
		beego.NSRouter("api/register/:uuid/", &v2.APIController{}, "*:Register"),
		beego.NSRouter("api/log/:uuid/", &v2.APIController{}, "*:Log"),
	)
	beego.AddNamespace(v1Namespace, v2Namespace)
}
