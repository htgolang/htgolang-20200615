package controllers

import "github.com/xxdu521/cmdbgo/server/controllers/auth"

type LayoutController struct {
	auth.LoginRequiredController
}

//layout控制器，是负责加载页面和js的，就是我们每个项目的GET:Index的需求，一定要引入loginrequired控制器的验证过程

//我的逻辑，称为页面控制器。
func (c *LayoutController) Prepare(){
	c.LoginRequiredController.Prepare()

	c.Layout = "layouts/base.html"
	c.LayoutSections = map[string]string{
		"LayoutStyle":"",
		"LayoutScript":"",
	}
	c.Data["menu"] = ""
	c.Data["expand"] = ""
}