package controllers

import (
	"strings"

	"github.com/imsilence/gocmdb/server/controllers/auth"
	"github.com/imsilence/gocmdb/server/models"
)

type AlarmPageController struct {
	LayoutController
}

func (c *AlarmPageController) Index() {
	c.Data["menu"] = "alarm_management"
	c.Data["expand"] = "monitoring"

	c.TplName = "alarm_page/index.html"
	c.LayoutSections["LayoutScript"] = "alarm_page/index.script.html"
}

type AlarmController struct {
	auth.LoginRequiredController
}

func (c *AlarmController) List() {
	//draw,start, length, q
	draw, _ := c.GetInt("draw")
	start, _ := c.GetInt64("start")
	length, _ := c.GetInt("length")
	q := strings.TrimSpace(c.GetString("q"))

	result, total, queryTotal := models.DefaultAlarmManager.Query(q, start, length)

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
