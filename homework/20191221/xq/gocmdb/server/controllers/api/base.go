package api

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare(){
	// 关闭xsrf 验证
	c.EnableXSRF = false
}

