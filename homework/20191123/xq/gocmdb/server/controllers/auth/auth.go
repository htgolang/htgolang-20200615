package auth

import (
	//"github.com/xlotz/gocmdb/server/controllers"
	"github.com/xlotz/gocmdb/server/controllers/base"
	"github.com/xlotz/gocmdb/server/models"
)

type LoginRequiredController struct {

	base.BaseController

	User *models.User

}

func (c *LoginRequiredController) Prepare(){

	c.BaseController.Prepare()

	//manager := NewManager()

	if user := DefaultManger.IsLogin(c); user == nil {

		DefaultManger.GoToLoginPage(c)  // togo

		c.StopRun()   // 登录后，停止后面的罗杰

	}else {

		c.User = user
		c.Data["user"] = user


	}


}

type AuthController struct {

	base.BaseController
}

func (c *AuthController) Login(){

	DefaultManger.Login(c)

}

func (c *AuthController) Logout(){


	DefaultManger.Logout(c)   // 销毁session , 跳转到登录页面
	//manager.GoToLoginPage(c)
}