package auth

import (
	"gocmdb/controllers/base"
	"gocmdb/models"
)

/*
	Controller所有的操作都交给manager.go来进行管理
*/

// 判断是否登录的Controller, 除login页面以外，都需要调用它
type LoginRequireController struct {
	base.BaseController

	User *models.User
}

//
func (c *LoginRequireController) Prepare() {
	c.BaseController.Prepare()

	// 判断是是否已经登录
	if user := DefaultManager.IsLogin(c); user == nil {
		// 未登录, 返回登录页面
		DefaultManager.GoToLoginPage(c)
		c.StopRun()
	} else {
		// 已登录, 将用户信息放入c.User和C.Data["user"]
		c.User = user
		c.Data["user"] = user
	}
}

// 登录验证的Controller
type AuthController struct {
	base.BaseController
}

// 登录
func (c *AuthController) Login() {
	DefaultManager.Login(c)
}

// 登出
func (c *AuthController) Logout() {
	DefaultManager.Logout(c)
}
