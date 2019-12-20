package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Index() {
	c.Redirect(beego.URLFor(beego.AppConfig.String("home")), http.StatusFound)
}
