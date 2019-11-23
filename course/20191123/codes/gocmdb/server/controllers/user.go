package controllers

import (
	"strings"

	"github.com/imsilence/gocmdb/server/controllers/auth"
	"github.com/imsilence/gocmdb/server/models"
)

type UserPageController struct {
	LayoutController
}

func (c *UserPageController) Index() {
	c.Data["menu"] = "user_management"
	c.Data["expand"] = "system_management"
	c.TplName = "user_page/index.html"
	c.LayoutSections["LayoutScript"] = "user_page/index.script.html"
}

type UserController struct {
	auth.LoginRequiredController
}

func (c *UserController) List() {
	//draw,start, length, q
	draw, _ := c.GetInt("draw")
	start, _ := c.GetInt64("start")
	length, _ := c.GetInt("length")
	q := strings.TrimSpace(c.GetString("q"))

	// []*User, total, queryTotal
	users, total, queryTotal := models.DefaultUserManager.Query(q, start, length)

	c.Data["json"] = map[string]interface{}{
		"code":            200,
		"text":            "获取成功",
		"result":          users,
		"draw":            draw,
		"recordsTotal":    total,
		"recordsFiltered": queryTotal,
	}
	c.ServeJSON()
}

func (c *UserController) Create() {
	if c.Ctx.Input.IsPost() {
		c.Data["json"] = map[string]interface{}{
			"code":   200,
			"text":   "创建成功",
			"result": nil, //可以返回创建的用户
		}
		c.ServeJSON()
	} else {
		//get
		c.TplName = "user/create.html"
	}
}

func (c *UserController) Modify() {
	if c.Ctx.Input.IsPost() {
		c.Data["json"] = map[string]interface{}{
			"code":   200,
			"text":   "编辑成功",
			"result": nil, //可以返回编辑后的用户
		}
		c.ServeJSON()
	} else {
		//get
		pk, _ := c.GetInt("pk")
		c.TplName = "user/modify.html"
		c.Data["object"] = models.DefaultUserManager.GetById(pk)
	}
}

func (c *UserController) Delete() {
	pk, _ := c.GetInt("pk")
	models.DefaultUserManager.DeleteById(pk)
	c.Data["json"] = map[string]interface{}{
		"code":   200,
		"text":   "删除成功",
		"result": nil, //可以返回删除的用户
	}
	c.ServeJSON()
}

func (c *UserController) Lock() {
	pk, _ := c.GetInt("pk")
	models.DefaultUserManager.SetStatusById(pk, 1)
	c.Data["json"] = map[string]interface{}{
		"code":   200,
		"text":   "锁定成功",
		"result": nil, //可以返回删除的用户
	}
	c.ServeJSON()
}

func (c *UserController) UnLock() {
	pk, _ := c.GetInt("pk")
	models.DefaultUserManager.SetStatusById(pk, 0)
	c.Data["json"] = map[string]interface{}{
		"code":   200,
		"text":   "解锁成功",
		"result": nil, //可以返回删除的用户
	}
	c.ServeJSON()
}

type TokenController struct {
	auth.LoginRequiredController
}

func (c *TokenController) Generate() {
	if c.Ctx.Input.IsPost() {
		pk, _ := c.GetInt("pk")
		models.DefaultTokenManager.GenerateByUser(models.DefaultUserManager.GetById(pk))
		c.Data["json"] = map[string]interface{}{
			"code":   200,
			"text":   "生成Token成功",
			"result": nil, //可以返回Token
		}
		c.ServeJSON()
	} else {
		pk, _ := c.GetInt("pk")
		c.Data["object"] = models.DefaultUserManager.GetById(pk)
		c.TplName = "token/index.html"
	}
}
