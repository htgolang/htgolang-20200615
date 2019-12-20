package controllers

import (
	"strings"

	"github.com/astaxie/beego/validation"
	"github.com/imsilence/gocmdb/server/controllers/auth"
	"github.com/imsilence/gocmdb/server/forms"
	"github.com/imsilence/gocmdb/server/models"
)

type AgentPageController struct {
	LayoutController
}

func (c *AgentPageController) Index() {
	c.Data["menu"] = "agent_management"
	c.Data["expand"] = "terminal_management"
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
func (c *AgentController) Modify() {
	if c.Ctx.Input.IsPost() {
		json := map[string]interface{}{
			"code": 400,
			"text": "提交数据错误",
		}
		form := &forms.AgentModifyForm{}
		valid := &validation.Validation{}
		if err := c.ParseForm(form); err == nil {
			if ok, err := valid.Valid(form); err != nil {
				valid.SetError("error", err.Error())
				json["result"] = valid.Errors
			} else if ok {
				result, err := models.DefaultAgentManager.Modify(
					form.Id,
					form.Remark,
				)

				if err == nil {
					json = map[string]interface{}{
						"code":   200,
						"text":   "更新成功",
						"result": result,
					}
				} else {
					json = map[string]interface{}{
						"code": 500,
						"text": "服务器错误",
					}
				}
			} else {
				json["result"] = valid.Errors
			}
		} else {
			valid.SetError("error", err.Error())
			json["result"] = valid.Errors
		}
		c.Data["json"] = json
		c.ServeJSON()

	} else {
		//get
		pk, _ := c.GetInt("pk")
		c.TplName = "agent/modify.html"
		c.Data["object"] = models.DefaultAgentManager.GetById(pk)
		// c.Data["types"] = cloud.DefaultManager.Plugins
	}
}

func (c *AgentController) Delete() {
	c.Data["json"] = map[string]interface{}{
		"code":   400,
		"text":   "提交数据错误",
		"result": nil,
	}
	if c.Ctx.Input.IsPost() {
		pk, _ := c.GetInt("pk")
		models.DefaultAgentManager.DeleteById(pk)

		c.Data["json"] = map[string]interface{}{
			"code":   200,
			"text":   "删除成功",
			"result": nil,
		}
	}

	c.ServeJSON()
}
