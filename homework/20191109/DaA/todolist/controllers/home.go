package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
)

type HomeController struct {
	beego.Controller
}

//home控制器暂时不写prepare方法

//定义get方法，用于接收主页请求
func (c *HomeController) Get() {
	//主页由配置文件去定义
	c.Redirect(beego.URLFor(beego.AppConfig.String("home")),http.StatusFound)
}
