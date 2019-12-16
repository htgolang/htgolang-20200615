package base

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

//基础控制器，在这里加载全局通用配置，比如xsrf攻击检查
func (c *BaseController) Prepare(){
	c.Data["xsrf_token"] = c.XSRFToken()
}