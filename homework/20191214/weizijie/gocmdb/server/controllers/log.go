package controllers

import (
	"strings"

	"github.com/imsilence/gocmdb/server/controllers/auth"
	"github.com/imsilence/gocmdb/server/models"
)

type LogPageController struct {
	LayoutController
}

func (c *LogPageController) Index() {
	c.Data["menu"] = "log_management"
	c.Data["expand"] = "terminal_management"
	c.TplName = "log_page/index.html"
	c.LayoutSections["LayoutScript"] = "log_page/index.script.html"
}

type LogController struct {
	auth.LoginRequiredController
}

func (c *LogController) List() {
	//draw,start, length, q
	draw, _ := c.GetInt("draw")
	start, _ := c.GetInt64("start")
	length, _ := c.GetInt("length")
	startTime := c.GetString("startTime")
	endTime := c.GetString("endTime")

	q := strings.TrimSpace(c.GetString("q"))
	result, total, queryTotal := models.DefaultResourceManager.Query(q, start, length, startTime, endTime)
	// []*User, total, queryTotall

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

func (c *LogController) Delete() {
	c.Data["json"] = map[string]interface{}{
		"code":   400,
		"text":   "提交数据错误",
		"result": nil,
	}
	if c.Ctx.Input.IsPost() {
		pk, _ := c.GetInt("pk")
		models.DefaultResourceManager.DeleteById(pk)

		c.Data["json"] = map[string]interface{}{
			"code":   200,
			"text":   "删除成功",
			"result": nil,
		}
	}

	c.ServeJSON()
}
