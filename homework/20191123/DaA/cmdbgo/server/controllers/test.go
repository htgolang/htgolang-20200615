package controllers

import (
	"github.com/xxdu521/cmdbgo/server/controllers/auth"
	"time"
)

type TestController struct {
	auth.LoginRequiredController
}

func (c *TestController) Test() {
	c.Data["json"] = map[string]interface{}{"now":time.Now()}
	c.ServeJSON()
}


type TestPageController struct {
	LayoutController
}

func (c *TestPageController) Index(){
	c.Data["menu"] = "user_management"
	c.Data["expand"] = "system_management"
	c.TplName = "test_page/index.html"
}