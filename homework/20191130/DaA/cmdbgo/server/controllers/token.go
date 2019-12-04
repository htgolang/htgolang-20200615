package controllers

import (
	"fmt"
	"github.com/xxdu521/cmdbgo/server/controllers/auth"
	"github.com/xxdu521/cmdbgo/server/models"
)

//token页面+数据,整合到一起了，这个是特殊形式的需求。
type TokenController struct {
	auth.LoginRequiredController
}
func (c *TokenController) Generate(){
	id,_ := c.GetInt("id")

	json := map[string]interface{}{
		"code": 405,
		"text": "请求方法错误",
		"result": nil,
	}

	if c.Ctx.Input.IsPost(){
		fmt.Println(id)
		token := models.DefaultTokenManager.GenerateByUser(models.DefaultUserManager.GetById(id))
		json = map[string]interface{}{
			"code": 200,
			"text": "生成token成功",
			"result": token,
		}
		c.Data["json"] = json
		c.ServeJSON()
	}

	if c.User.Id == id {
		c.Data["object"] = models.DefaultUserManager.GetById(id)
	} else {
		c.Data["object"] = models.Token{}
	}

	fmt.Println(c.Data["object"])
	c.TplName = "token/index.html"
}