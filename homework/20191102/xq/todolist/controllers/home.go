package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
)

// 主页面控制器
type HomeController struct {
	BaseController
}

func (c *HomeController) Get() {
	c.Redirect(beego.URLFor(beego.AppConfig.String("home")), http.StatusFound)
}
