package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"net/http"
	"todolist/models"
)

//定义一个控制器，然后给控制器一个prepare方法（可以没有），再写需要的方法（autorouter模式）

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {
	//暂时什么也不做，留着给未来扩容，做事情
	//基本的seesion的验证，我们放在login控制器内做
}
type LoginController struct {
	BaseController //继承那个留给未来做扩展的base控制器，而不是beego.controller
	User *models.User //导入User数据，用于做登录的判断验证
}

func (c *LoginController) Prepare() {
	c.BaseController.Prepare()
	//检查用户是否已经登录（在session中可以获取正确的用户id）
	if session := c.GetSession("user"); session != nil {
		if id,ok := session.(int);ok {
			user := models.User{Id:id}
			ormer := orm.NewOrm()
			if ormer.Read(&user) == nil {
				beego.ReadFromRequest(&c.Controller) //从reqeust中读取flash cookie数据
				c.User = &user
				c.Data["user"] = &user
			}
		}
	}
	if c.User == nil {
		// 用户未登陆则跳转到登陆控制器，并使用next参数传递请求的url
		if c.Ctx.Input.IsAjax(){
			c.Data["json"] = map[string]interface{}{
				"code": 403,
				"text": "未认证",
			}
			c.ServeJSON()
		} else {
			c.Redirect(beego.URLFor(beego.AppConfig.String("login"),"next",c.Ctx.Input.URL()),http.StatusFound)
		}
	}
}