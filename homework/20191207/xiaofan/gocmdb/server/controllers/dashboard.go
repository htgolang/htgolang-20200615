package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
)

type HomePageController struct {
	LayoutController
}

// homepage/index
func (c *HomePageController) Index() {
	c.Redirect(beego.URLFor(beego.AppConfig.String("home")), http.StatusFound)
}

type DashboardPageController struct {
	LayoutController
}

// dashboardpage/index
func (c *DashboardPageController) Index() {
	c.Data["menu"] = "dashboard"
	c.TplName = "dashboard_page/index.html"
	c.LayoutSections["LayoutScript"] = "dashboard_page/index.script.html"
	c.Data["platformTotal"] = 1
	c.Data["virtualmachineTotal"] = 1
	c.Data["alertTotal"] = 3
	c.Data["userTotal"] = 2
}
