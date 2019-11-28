package controllers

import (
	"gocmdb/controllers/auth"
	"time"
)

type TestController struct {
	auth.LoginRequireController
}

func (c *TestController) List() {
	c.Data["json"] = time.Now()
	c.ServeJSON()
}

type TestPageController struct {
	LayoutController
}

func (c *TestPageController) Index() {
	c.TplName = "test_page/index.html"
}
