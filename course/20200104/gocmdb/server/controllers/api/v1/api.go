package v1

import (
	"encoding/json"

	"github.com/imsilence/gocmdb/server/controllers/api"
	"github.com/imsilence/gocmdb/server/models"
)

type APIController struct {
	api.BaseController
}

func (c *APIController) Heartbeat() {
	models.DefaultAgentManager.Heartbeat(c.Ctx.Input.Param(":uuid"))
	c.Data["json"] = map[string]interface{}{
		"code":   200,
		"text":   "成功",
		"result": nil,
	}
	c.ServeJSON()
}

func (c *APIController) Register() {
	rt := map[string]interface{}{
		"code":   400,
		"text":   "",
		"result": nil,
	}
	agent := &models.Agent{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, agent); err == nil {
		agent, created, err := models.DefaultAgentManager.CreateOrReplace(agent)
		if err == nil {
			rt = map[string]interface{}{
				"code": 200,
				"text": "成功",
				"result": map[string]interface{}{
					"created": created,
					"agent":   agent,
				},
			}
		} else {
			rt["text"] = err.Error()
		}
	} else {
		rt["text"] = err.Error()
	}
	c.Data["json"] = rt
	c.ServeJSON()
}

func (c *APIController) Log() {
	log := &models.Log{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, log); err == nil {
		models.DefaultLogManager.Create(log)
	}
	c.Data["json"] = map[string]interface{}{
		"code":   200,
		"text":   "成功",
		"result": nil,
	}
	c.ServeJSON()
}

func (c *APIController) Task() {
	c.Data["json"] = map[string]interface{}{
		"code":   200,
		"text":   "成功",
		"result": models.DefaultTaskManager.GetByUuid(c.Ctx.Input.Param(":uuid")),
	}
	c.ServeJSON()
}

func (c *APIController) TaskResult() {
	rt := map[string]interface{}{
		"code":   400,
		"text":   "",
		"result": nil,
	}

	result := &models.Result{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, result); err == nil {
		if err := models.DefaultTaskManager.Result(c.Ctx.Input.Param(":uuid"), result); err != nil {
			rt["text"] = err.Error()
		} else {
			rt = map[string]interface{}{
				"code":   200,
				"text":   "成功",
				"result": nil,
			}
		}
	} else {
		rt["text"] = err.Error()
	}
	c.Data["json"] = rt
	c.ServeJSON()
}
