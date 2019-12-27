package v2

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/dcosapp/gocmdb/server/controllers/api"
	"github.com/dcosapp/gocmdb/server/models"
)

type APIController struct {
	api.BaseController
}

// 重写Prepare
func (c *APIController) Prepare() {
	// 重写之后，需要调用BaseController的Prepare()来取消xsrf验证
	c.BaseController.Prepare()
	// 验证客户端token与配置文件中的是否相同
	if beego.AppConfig.String("agent::token") != c.Ctx.Input.Header("Token") {
		c.Data["json"] = map[string]interface{}{
			"code":   400,
			"text":   "Token不正确",
			"result": nil,
		}
		c.ServeJSON()
		c.StopRun()
	}

}

// v1/api/heartbeat/:uuid
func (c *APIController) Heartbeat() {
	models.DefaultAgentManager.Heartbeat(c.Ctx.Input.Param(":uuid"))
	c.Data["json"] = map[string]interface{}{
		"code":   200,
		"text":   "成功",
		"result": nil,
	}
	c.ServeJSON()
}

// v1/api/register/:uuid
func (c *APIController) Register() {
	rt := map[string]interface{}{
		"code":   400,
		"text":   "请求失败",
		"result": nil,
	}

	agent := &models.Agent{}
	// 将接收到的json结构体反序列化
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, agent); err == nil {
		agent, created, err := models.DefaultAgentManager.CreateOrReplace(agent)
		if err == nil {
			rt["code"], rt["text"], rt["result"] = 200, "成功", map[string]interface{}{"created": created, "agent": agent}
		} else {
			rt["code"], rt["result"] = 500, err.Error()
		}
	} else {
		rt["result"] = err.Error()
	}
	c.Data["json"] = rt
	c.ServeJSON()
}

// v1/api/log/:uuid
func (c *APIController) Log() {
	rt := map[string]interface{}{
		"code":   400,
		"text":   "请求失败",
		"result": nil,
	}
	log := &models.Log{}
	// 将接收到的json结构体反序列化
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, log); err == nil {
		if err := models.DefaultLogManager.Create(log); err == nil {
			rt["code"], rt["text"] = 200, "成功"
		} else {
			rt["code"], rt["result"] = 500, err.Error()
		}
	} else {
		rt["result"] = err.Error()
	}
	c.Data["json"] = rt
	c.ServeJSON()
}
