package controllers

import (
	"github.com/astaxie/beego/validation"
	"github.com/dcosapp/gocmdb/server/controllers/auth"
	"github.com/dcosapp/gocmdb/server/forms"
	"github.com/dcosapp/gocmdb/server/models"
	"strings"
)

// logpage/
type LogPageController struct {
	LayoutController
}

// logpage/index
func (c *LogPageController) Index() {
	c.Data["menu"] = "log_management"
	c.Data["expand"] = "agent_management"
	c.TplName = "log_page/index.html"
	c.LayoutSections["LayoutScript"] = "log_page/index.script.html"
}

// resourcepage/
type ResourcePageController struct {
	LayoutController
}

// resourcepage/index
func (c *ResourcePageController) Index() {
	c.Data["menu"] = "resource_management"
	c.Data["expand"] = "agent_management"
	c.TplName = "resource_page/index.html"
	c.LayoutSections["LayoutScript"] = "resource_page/index.script.html"
}

// resource/
type ResourceController struct {
	auth.LoginRequireController
}

// resource/list
func (c *ResourceController) List() {
	// draw, start, length, q
	draw, _ := c.GetInt("draw")
	start, _ := c.GetInt64("start")
	length, _ := c.GetInt("length")
	q := strings.TrimSpace(c.GetString("q"))

	// []*VirtualMachine, total, queryTotal
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

// resource/modify
func (c *ResourceController) Modify() {
	json := map[string]interface{}{
		"code":   405,
		"text":   "请求方式错误",
		"result": nil,
	}

	if c.Ctx.Input.IsPost() {
		form := &forms.AgentForm{}
		valid := &validation.Validation{}

		json["code"], json["text"] = 400, "请求数据错误"
		if err := c.ParseForm(form); err != nil {
			json["text"] = err.Error()
		} else {
			if ok, err := valid.Valid(form); err != nil {
				json["text"] = err.Error()
			} else if !ok {
				json["result"] = valid.Errors
			} else {
				// 更新数据库
				if result, err := models.DefaultAgentManager.Modify(form.Id, form.Name, form.Desc); err == nil {
					json["code"], json["text"], json["result"] = 200, "修改终端成功", result
				} else {
					json["code"], json["text"], json["result"] = 500, "服务器端错误", err.Error()
				}
			}
		}

		c.Data["json"] = json
		c.ServeJSON()
	}

	pk, _ := c.GetInt("pk")
	c.TplName = "resource/modify.html"
	c.Data["object"] = models.DefaultAgentManager.GetById(pk)
}

// resource/delete
func (c *ResourceController) Delete() {
	json := map[string]interface{}{
		"code":   400,
		"text":   "请求数据错误",
		"result": nil,
	}

	if pk, err := c.GetInt("pk"); err != nil {
		json["text"] = err.Error()
	} else {
		if result, err := models.DefaultAgentManager.DeleteById(pk); err == nil {
			json["code"], json["text"], json["result"] = 200, "客户端删除成功", result
		} else {
			json["code"], json["text"], json["result"] = 500, "服务器端错误", err.Error()
		}
	}

	c.Data["json"] = json
	c.ServeJSON()
}

type LogController struct {
	auth.LoginRequireController
}

// log/list
func (c *LogController) List() {
	// draw, start, length, q
	draw, _ := c.GetInt("draw")
	start, _ := c.GetInt64("start")
	length, _ := c.GetInt("length")
	startTime := c.GetString("startTime")
	endTime := c.GetString("endTime")

	q := strings.TrimSpace(c.GetString("q"))

	// []*VirtualMachine, total, queryTotal
	result, total, queryTotal := models.DefaultResourceManager.Query(q, start, length, startTime, endTime)
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
