package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Index() {
	c.Redirect(beego.URLFor(beego.AppConfig.String("home")), http.StatusFound)
}
