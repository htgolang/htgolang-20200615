package base

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

// 为所有的Controller设置xsrf的token
func (c *BaseController) Prepare() {
	c.Data["xsrf_token"] = c.XSRFToken()
}
