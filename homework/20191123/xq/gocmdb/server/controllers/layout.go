package controllers

import (

	"github.com/xlotz/gocmdb/server/controllers/auth"
)

type LayoutController struct {
	auth.LoginRequiredController
}

func (c *LayoutController) Prepare(){

	c.LoginRequiredController.Prepare()
	c.Layout = "layouts/base.html"

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["LayoutStyle"] = ""
	c.LayoutSections["LayoutScript"] = ""

	// 用于通过判断一级、二级标签显示菜单状态
	c.Data["menu"] = ""
	c.Data["expand"] = ""

}