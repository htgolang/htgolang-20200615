package controllers

import (
	"fmt"
	//"strings"
	"github.com/xlotz/gocmdb/server/controllers/auth"
	"strings"

	"github.com/xlotz/gocmdb/server/models"

)

type AlarmPageController struct {
	LayoutController
}

func (c *AlarmPageController) Index() {
	c.Data["menu"] = "alarm_management"
	c.Data["expand"] = "terminal_management"
	c.TplName = "alarm/index.html"
	c.LayoutSections["LayoutScript"] = "alarm/index.script.html"
}

type AlarmController struct {
	auth.LoginRequiredController
}

func (c *AlarmController) List() {

	draw, _ := c.GetInt("draw")
	start, _ := c.GetInt64("start")
	length, _ := c.GetInt("length")
	q := strings.TrimSpace(c.GetString("q"))

	result, total, queryTotal := models.DefaultAlarmManager.Query(q, start, length)
	fmt.Println(result)

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

