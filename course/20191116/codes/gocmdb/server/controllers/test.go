package controllers

import (
	"time"

	"github.com/imsilence/gocmdb/server/controllers/auth"
)

type TestController struct {
	auth.LoginRequiredController
}

func (c *TestController) Test() {
	c.Data["json"] = map[string]interface{}{"now": time.Now()}
	c.ServeJSON()
}
