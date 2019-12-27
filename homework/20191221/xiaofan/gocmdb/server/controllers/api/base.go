package api

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

// 所有访问v1/api/接口的都不需要xsrf认证
func (c *BaseController) Prepare() {
	c.EnableXSRF = false
}
