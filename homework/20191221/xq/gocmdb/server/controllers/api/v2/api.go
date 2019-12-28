package v2

import (
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/xlotz/gocmdb/server/controllers/api"
	"github.com/xlotz/gocmdb/server/models"
	"encoding/json"
)


type APIController struct {

	api.BaseController

}

func (c *APIController) Prepare(){

	//c.EnableXSRF = false
	c.BaseController.Prepare()

	if beego.AppConfig.String("agent::token") != c.Ctx.Input.Header("Token"){
		c.Data["json"] = map[string]interface{}{
			"code": 400,
			"text": "Token认证失败",
			"result": nil,
		}
		c.ServeJSON()
		c.StopRun()
	}


}

func (c *APIController) Heartbeat() {


	//fmt.Println(c.Ctx.Input.Param(":uuid"))
	// 获取传的数据
	//fmt.Println(string(c.Ctx.Input.RequestBody))

	models.DefaultAgentManager.Heartbeat(c.Ctx.Input.Param(":uuid"))

	json := map[string]interface{}{
		"code": 200,
		"text": "成功",
		"result": nil,
	}

	c.Data["json"] = json
	c.ServeJSON()
}

func (c *APIController) Register() {
	rt := map[string]interface{}{
		"code": 200,
		"text": "注册成功",
		"result": nil,
	}

	//fmt.Println(c.Ctx.Input.Param(":uuid"))
	// 获取传的数据
	//fmt.Println(string(c.Ctx.Input.RequestBody))

	agent := &models.Agent{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, agent); err == nil {
		agent, created, err := models.DefaultAgentManager.CreateOrReplace(agent)
		if err == nil {
			rt = map[string]interface{}{
				"code": 200,
				"text": "成功",
				"result": map[string]interface{}{
					"created": created,
					"agent": agent,
				},
			}
		}else {
			rt["text"] = err.Error()
		}
	}else {
		rt["text"] = err.Error()
	}

	c.Data["json"] = rt
	c.ServeJSON()
}

func (c *APIController) Log() {
	rt := map[string]interface{}{
		"code":400,
		"text": "请求失败",
		"result": nil,
	}

	log := &models.Log{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, log); err == nil {

		models.DefaultLogManager.Create(log)

		rt["code"] = 200
		rt["text"] = "日志上传成功"
		rt["result"] = log

	}else {
		rt["result"] = err.Error()

	}


	c.Data["json"] = rt
	c.ServeJSON()
}