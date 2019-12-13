package routers

import (
	"github.com/astaxie/beego"
	"github.com/dcosapp/gocmdb/server/controllers"
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
}
