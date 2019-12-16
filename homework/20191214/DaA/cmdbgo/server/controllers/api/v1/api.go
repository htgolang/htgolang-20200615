package v1

import (
	"encoding/json"
	"github.com/xxdu521/cmdbgo/server/controllers/api"
	"github.com/xxdu521/cmdbgo/server/models"
)

type APIController struct {
	api.BaseController
}

func (c *APIController) Heartbeat(){
	models.DefaultAgentManager.Heartbeat(c.Ctx.Input.Param(":uuid"))
	c.Data["json"] = map[string]interface{}{
		"code": 200,
		"text": "心跳发送成功",
		"result": nil,
	}
	c.ServeJSON()
}

func (c *APIController) Register(){
	rt := map[string]interface{}{
		"code": 400,
		"text": "请求失败",
		"result": nil,
	}

	agent := &models.Agent{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, agent);err == nil {
		agent, create, err := models.DefaultAgentManager.CreateOrReplace(agent)
		if  err == nil {
			rt["code"],rt["text"],rt["result"] = 200,"注册信息成功",map[string]interface{}{
				"created": create,
				"agent": agent,
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

func (c *APIController) Log(){
	rt := map[string]interface{}{
		"code": 400,
		"text": "请求失败",
		"result": nil,
	}

	log := &models.Log{}
	//fmt.Println(string(c.Ctx.Input.RequestBody))
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, log); err == nil {
		if err := models.DefaultLogManager.Create(log); err == nil {
			rt["code"], rt["text"], rt["result"] = 200,"上传日志成功",nil
		} else {
			rt["text"] = err.Error()
		}
	}
	c.Data["json"] = rt
	c.ServeJSON()
}