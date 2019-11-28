package controllers

import (
	"fmt"
	"github.com/xxdu521/cmdbgo/server/controllers/auth"
	"github.com/xxdu521/cmdbgo/server/models"
)

type TokenController struct {
	auth.LoginRequiredController
}

func (c *TokenController) Generate(){
	id,_ := c.GetInt("id")
	if c.Ctx.Input.IsPost(){
		fmt.Println(id)
		token := models.DefaultTokenManager.GenerateByUser(models.DefaultUserManager.GetById(id))
		json := map[string]interface{}{
			"code": 200,
			"text": "生成token成功",
			"result": token,
		}
		c.Data["json"] = json
		c.ServeJSON()
	} else {
		c.Data["object"] = models.DefaultUserManager.GetById(id)
		c.TplName = "token/index.html"
	}
}