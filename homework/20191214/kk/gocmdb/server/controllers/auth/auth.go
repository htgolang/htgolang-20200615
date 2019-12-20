package auth

import (
	"github.com/imsilence/gocmdb/server/controllers/base"

	"github.com/imsilence/gocmdb/server/models"
)

type LoginRequiredController struct {
	base.BaseController

	User *models.User
}

func (c *LoginRequiredController) Prepare() {
	c.BaseController.Prepare()

	if user := DefaultManger.IsLogin(c); user == nil {
		// 未登陆
		DefaultManger.GoToLoginPage(c) // todo 需要修改参数
		c.StopRun()
	} else {
		// 已登陆
		c.User = user
		c.Data["user"] = user
	}
}

type AuthController struct {
	base.BaseController
}

func (c *AuthController) Login() {
	DefaultManger.Login(c)
}

func (c *AuthController) Logout() {
	DefaultManger.Logout(c)
}
