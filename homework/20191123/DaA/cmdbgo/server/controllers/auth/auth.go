package auth

import (
	"github.com/xxdu521/cmdbgo/server/controllers/base"
	"github.com/xxdu521/cmdbgo/server/models"
)

type LoginRequiredController struct {
	base.BaseController

	User *models.User
}

func (c *LoginRequiredController) Prepare() {
	c.BaseController.Prepare()

	if user := DefaultManager.IsLogin(c) ; user == nil {
		//空值为未登录
		DefaultManager.GoToLoginPage(c)
		c.StopRun()
	} else {
		c.User = user
		c.Data["user"] = user
	}
}

type AuthController struct {
	base.BaseController
}

func (c *AuthController) Login(){
	DefaultManager.Login(c)
}

func (c *AuthController) Logout(){
	DefaultManager.Logout(c)

}