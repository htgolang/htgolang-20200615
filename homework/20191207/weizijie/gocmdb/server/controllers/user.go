package controllers

import (
	"strings"

	"github.com/astaxie/beego/validation"
	"github.com/imsilence/gocmdb/server/controllers/auth"
	"github.com/imsilence/gocmdb/server/forms"
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
		json := map[string]interface{}{
			"code": 400,
			"text": "提交数据错误",
		}

		form := &forms.UserCreateForm{}
		valid := &validation.Validation{}
		if err := c.ParseForm(form); err == nil {
			if ok, err := valid.Valid(form); err != nil {
				valid.SetError("error", err.Error())
				json["result"] = valid.Errors
			} else if ok {
				user, err := models.DefaultUserManager.Create(form.Name, form.Password, form.Gender, form.BirthdayTime, form.Tel, form.Email, form.Addr, form.Remark)
				if err == nil {
					json = map[string]interface{}{
						"code":   200,
						"text":   "创建成功",
						"result": user,
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
		c.TplName = "user/create.html"
	}
}

func (c *UserController) Modify() {
	if c.Ctx.Input.IsPost() {
		json := map[string]interface{}{
			"code": 400,
			"text": "提交数据错误",
		}
		form := &forms.UserModifyForm{}
		valid := &validation.Validation{}
		if err := c.ParseForm(form); err == nil {
			if ok, err := valid.Valid(form); err != nil {
				valid.SetError("error", err.Error())
				json["result"] = valid.Errors
			} else if ok {
				user, err := models.DefaultUserManager.Modify(form.Id, form.Name, form.Gender, form.BirthdayTime, form.Tel, form.Email, form.Addr, form.Remark)

				if err == nil {
					json = map[string]interface{}{
						"code":   200,
						"text":   "更新成功",
						"result": user,
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
		c.TplName = "user/modify.html"
		c.Data["object"] = models.DefaultUserManager.GetById(pk)
	}
}

func (c *UserController) Delete() {
	pk, _ := c.GetInt("pk")
	if pk != c.User.Id {
		models.DefaultUserManager.DeleteById(pk)
		c.Data["json"] = map[string]interface{}{
			"code":   200,
			"text":   "删除成功",
			"result": nil, //可以返回删除的用户
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"code":   400,
			"text":   "当前用户不能删除自己",
			"result": nil, //可以返回删除的用户
		}
	}

	c.ServeJSON()
}

func (c *UserController) Lock() {
	pk, _ := c.GetInt("pk")
	if pk != c.User.Id {
		models.DefaultUserManager.SetStatusById(pk, 1)
		c.Data["json"] = map[string]interface{}{
			"code":   200,
			"text":   "锁定成功",
			"result": nil, //可以返回删除的用户
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"code":   400,
			"text":   "当前用户不能锁定自己",
			"result": nil, //可以返回删除的用户
		}
	}
	c.ServeJSON()
}

func (c *UserController) UnLock() {
	pk, _ := c.GetInt("pk")
	if pk != c.User.Id {
		models.DefaultUserManager.SetStatusById(pk, 0)
		c.Data["json"] = map[string]interface{}{
			"code":   200,
			"text":   "解锁成功",
			"result": nil, //可以返回删除的用户
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"code":   400,
			"text":   "当前用户不能解锁自己",
			"result": nil, //可以返回删除的用户
		}
	}
	c.ServeJSON()
}

func (c *UserController) Password() {
	if c.Ctx.Input.IsPost() {
		json := map[string]interface{}{
			"code": 400,
			"text": "提交数据错误",
		}

		form := &forms.UserPasswordForm{User: c.User}
		valid := &validation.Validation{}
		if err := c.ParseForm(form); err == nil {
			if ok, err := valid.Valid(form); err != nil {
				valid.SetError("error", err.Error())
				json["result"] = valid.Errors
			} else if ok {
				err := models.DefaultUserManager.UpdatePassword(c.User.Id, form.Password)
				if err == nil {
					json = map[string]interface{}{
						"code": 200,
						"text": "修改密码成功",
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
		c.TplName = "user/password.html"
	}
}

type TokenController struct {
	auth.LoginRequiredController
}

func (c *TokenController) Generate() {
	if c.Ctx.Input.IsPost() {
		// pk, _ := c.GetInt("pk")
		// // 方法1： 判断pk是否等于自己
		// if pk == c.User.Id {
		// 	models.DefaultTokenManager.GenerateByUser(models.DefaultUserManager.GetById(pk))
		// 	c.Data["json"] = map[string]interface{}{
		// 		"code":   200,
		// 		"text":   "生成Token成功",
		// 		"result": nil, //可以返回Token
		// 	}
		// } else {
		// 	c.Data["json"] = map[string]interface{}{
		// 		"code":   400,
		// 		"text":   "只能对自己的Token进行生成",
		// 		"result": nil, //可以返回Token
		// 	}
		// }

		// 方法2：获取c.User操作
		models.DefaultTokenManager.GenerateByUser(c.User)
		c.Data["json"] = map[string]interface{}{
			"code":   200,
			"text":   "生成Token成功",
			"result": nil, //可以返回Token
		}
		c.ServeJSON()
	} else {
		// pk, _ := c.GetInt("pk")
		// // 方法1： 判断pk是否等于自己
		// if pk == c.User.Id {
		// 	c.Data["object"] = models.DefaultUserManager.GetById(pk)
		// }
		// 方法2：获取c.User.Id
		c.Data["object"] = models.DefaultUserManager.GetById(c.User.Id)
		c.TplName = "token/index.html"
	}
}
