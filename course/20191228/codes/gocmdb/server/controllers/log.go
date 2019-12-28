package controllers

import (
	"strings"

	"github.com/imsilence/gocmdb/server/controllers/auth"
	"github.com/imsilence/gocmdb/server/models"
)

type ResourcePageController struct {
	LayoutController
}

func (c *ResourcePageController) Index() {
	c.Data["menu"] = "resource"
	c.Data["expand"] = "log_management"

	c.TplName = "resource_page/index.html"
	c.LayoutSections["LayoutScript"] = "resource_page/index.script.html"
}

type ResourceController struct {
	auth.LoginRequiredController
}

func (c *ResourceController) List() {
	//draw,start, length, q
	draw, _ := c.GetInt("draw")
	start, _ := c.GetInt64("start")
	length, _ := c.GetInt("length")
	q := strings.TrimSpace(c.GetString("q"))

	result, total, queryTotal := models.DefaultResourceManager.Query(q, start, length)

	c.Data["json"] = map[string]interface{}{
		"code":            200,
		"text":            "获取成功",
		"result":          result,
		"draw":            draw,
		"recordsTotal":    total,
		"recordsFiltered": queryTotal,
	}
	c.ServeJSON()
}


func (c *ResourceController) Trend() {
	uuid := strings.TrimSpace(c.GetString("uuid"))

	result := models.DefaultResourceManager.Trend(uuid)

	c.Data["json"] = map[string]interface{}{
		"code":            200,
		"text":            "获取成功",
		"result":          result,
	}
	c.ServeJSON()
}