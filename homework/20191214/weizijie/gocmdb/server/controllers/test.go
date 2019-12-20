package controllers

import (
	"fmt"

	"github.com/imsilence/gocmdb/server/controllers/auth"
)

type TestController struct {
	auth.LoginRequiredController
}

func (c *TestController) Test() {
	fmt.Println(c.Ctx.Input.RequestBody)
	c.Data["json"] = 1
	c.ServeJSON()
}
