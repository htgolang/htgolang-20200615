package controllers

import "gocmdb/controllers/auth"

type LayoutController struct {
	auth.LoginRequireController
}

// 基础Layout配置，所有页面都需要用到
func (c *LayoutController) Prepare() {
	c.LoginRequireController.Prepare()
	c.Layout = "layouts/base.html"
	c.LayoutSections = map[string]string{
		"LayoutStyle":  "",
		"LayoutScript": "",
	}

	c.Data["menu"] = ""
	c.Data["expand"] = ""
}
