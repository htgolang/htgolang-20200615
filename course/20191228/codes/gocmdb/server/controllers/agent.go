package controllers

import (
	"strings"

	"github.com/imsilence/gocmdb/server/controllers/auth"
	"github.com/imsilence/gocmdb/server/models"
)

type AgentPageController struct {
	LayoutController
}

func (c *AgentPageController) Index() {
	c.Data["menu"] = "agent_management"

	c.TplName = "agent_page/index.html"
	c.LayoutSections["LayoutScript"] = "agent_page/index.script.html"
}

type AgentController struct {
	auth.LoginRequiredController
}

func (c *AgentController) List() {
	//draw,start, length, q
	draw, _ := c.GetInt("draw")
	start, _ := c.GetInt64("start")
	length, _ := c.GetInt("length")
	q := strings.TrimSpace(c.GetString("q"))

	result, total, queryTotal := models.DefaultAgentManager.Query(q, start, length)

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

func (c *AgentController) Delete() {
	if c.Ctx.Input.IsPost() {
		pk, _ := c.GetInt("pk")
		models.DefaultAgentManager.DeleteById(pk)
	}
	c.Data["json"] = map[string]interface{}{
		"code":   200,
		"text":   "删除成功",
		"result": nil,
	}
	c.ServeJSON()
}
