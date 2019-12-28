package controllers

import (
	"fmt"
	"strings"
	"github.com/xlotz/gocmdb/server/controllers/auth"

	"github.com/xlotz/gocmdb/server/models"

)

type AgentPageController struct {
	LayoutController
}

func (c *AgentPageController) Index() {
	c.Data["menu"] = "agent_management"
	c.Data["expand"] = "terminal_management"
	c.TplName = "agent/index.html"
	c.LayoutSections["LayoutScript"] = "agent/index.script.html"
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

	// []*User, total, queryTotal
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




func (c *AgentController) Delete(){
	if c.Ctx.Input.IsPost(){

		pk, _ := c.GetInt("pk")

		models.DefaultAgentManager.DeleteById(pk)

		c.Data["json"] = map[string]interface{}{
			"code": 200,
			"text": "删除成功",
			"result": nil,
		}

	}else {
		c.Data["json"] = map[string]interface{}{
			"code": 400,
			"text": "删除失败",
			"result": nil,
		}
	}

	c.ServeJSON()

}

// 日志管理

type LogPageController struct {
	LayoutController
}

func (c *LogPageController) Index() {
	c.Data["menu"] = "log_management"
	c.Data["expand"] = "terminal_management"
	c.TplName = "log/index.html"
	c.LayoutSections["LayoutScript"] = "log/index.script.html"
}

type LogController struct {
	auth.LoginRequiredController
}

func (c *LogController) List() {
	//draw,start, length, q
	draw, _ := c.GetInt("draw")
	start, _ := c.GetInt64("start")
	length, _ := c.GetInt("length")
	q := strings.TrimSpace(c.GetString("q"))


	// []*User, total, queryTotal
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


func (c *LogController) Trend() {

	uuid:= strings.TrimSpace(c.GetString("uuid"))

	result:= models.DefaultResourceManager.Trend(uuid)

	fmt.Println(result)

	c.Data["json"] = map[string]interface{}{
		"code":            200,
		"text":            "获取成功",
		"result":          result,
	}
	c.ServeJSON()
}

func (c *LogController) Delete(){
	if c.Ctx.Input.IsPost(){

		pk, _ := c.GetInt("pk")

		models.DefaultResourceManager.DeleteById(pk)

		c.Data["json"] = map[string]interface{}{
			"code": 200,
			"text": "删除成功",
			"result": nil,
		}

	}else {
		c.Data["json"] = map[string]interface{}{
			"code": 400,
			"text": "删除失败",
			"result": nil,
		}
	}

	c.ServeJSON()

}

