package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) Get() {
	c.Redirect(beego.URLFor(beego.AppConfig.String("home")), http.StatusFound)
}
