package auth

import (
	"github.com/xxdu521/cmdbgo/server/controllers/base"
	"github.com/xxdu521/cmdbgo/server/models"
)

//实际上，整个登录过程，分成两个部分，一部分是auth控制器，负责真正的登录，退出的访问处理。

//另一个部分是loginrequired控制器部分，这个控制器，是给数据获取准备的。

type LoginRequiredController struct {
	base.BaseController

	User *models.User	//这个控制器是内部验证用的，所以带有一个User的实例，用来存储当前登录的用户信息
}
//用户状态判断层，loginrequired负责处理用户登录状态，用户如果没有登录，就强制用户登录，用户登录的话，就添加用户信息

//loginrequired控制器只有Prepare方法，负责做用户登录状态的验证。还有一个layout控制器，负责处理html和js代码。

//我的逻辑，称为数据控制器。
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

//登录页的入口，auth/login，调用管理器的login方法
func (c *AuthController) Login(){
	DefaultManager.Login(c)
}
//退出的入口，auth/logout，调用管理器的logout方法
func (c *AuthController) Logout(){
	DefaultManager.Logout(c)
}