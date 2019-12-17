package controllers

import (
	"github.com/astaxie/beego"
	"github.com/xxdu521/cmdbgo/server/controllers/auth"
	"github.com/xxdu521/cmdbgo/server/models"
	"strings"
)

type AgentPageController struct {
	LayoutController
}
func (c *AgentPageController) Index(){
	c.Data["expand"] = "machine_room_management"
	c.Data["menu"] = "agent_management"
	c.TplName = "agent_page/index.html"
	c.LayoutSections["LayoutScript"] = "agent_page/index_script.html"
}

type AgentController struct {
	auth.LoginRequiredController
}
func (c *AgentController) List ()  {
	draw,_ := c.GetInt("draw")
	start,_ := c.GetInt64("start")
	length,_ := c.GetInt("length")

	Max_Query_Length,_ := beego.AppConfig.Int("Max_Query_Length")
	if Max_Query_Length > 10 && length > Max_Query_Length {
		length = Max_Query_Length
	}

	q := strings.TrimSpace(c.GetString("q"))

	result, total, querytotal := models.DefaultAgentManager.Query(q, start, length)

	c.Data["json"] = map[string]interface{}{
		"code": 			200,
		"text": 			"成功",
		"result": 			result,
		"draw": 			draw,
		"recordsTotal": 	total,
		"recordsFiltered": 	querytotal,
	}
	c.ServeJSON()
}
func (c *AgentController) Modify() {}  //agent信息修改，还没写
func (c *AgentController) Delete() {} //删除agent信息，还没写


type ResourcePageController struct {
	LayoutController
}
func (c *ResourcePageController) Index(){
	c.Data["expand"] = "machine_room_management"
	c.Data["menu"] = "resource_management"
	c.TplName = "resource_page/index.html"
	c.LayoutSections["LayoutScript"] = "resource_page/index_script.html"
}

type ResourceController struct {
	auth.LoginRequiredController
}
func (c *ResourceController) List ()  {
	draw,_ := c.GetInt("draw")
	start,_ := c.GetInt64("start")
	length,_ := c.GetInt("length")

	Max_Query_Length,_ := beego.AppConfig.Int("Max_Query_Length")
	if Max_Query_Length > 10 && length > Max_Query_Length {
		length = Max_Query_Length
	}

	q := strings.TrimSpace(c.GetString("q"))

	result, total, querytotal := models.DefaultResourceManager.Query(q, start, length)

	c.Data["json"] = map[string]interface{}{
		"code": 			200,
		"text": 			"成功",
		"result": 			result,
		"draw": 			draw,
		"recordsTotal": 	total,
		"recordsFiltered": 	querytotal,
	}
	c.ServeJSON()
}



