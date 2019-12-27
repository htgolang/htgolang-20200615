package controllers

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/dcosapp/gocmdb/server/controllers/auth"
	"github.com/dcosapp/gocmdb/server/forms"
	"github.com/dcosapp/gocmdb/server/models"
	"strings"
	"time"
)

// userpage/
type UserPageController struct {
	LayoutController
}

// userpage/index
func (c *UserPageController) Index() {
	c.Data["menu"] = "user_management"
	c.Data["expand"] = "system_management"
	c.TplName = "user_page/index.html"
	c.LayoutSections["LayoutScript"] = "user_page/index.script.html"
}

// user/
type UserController struct {
	auth.LoginRequireController
}

// user/list
func (c *UserController) List() {
	// draw, start, length, q
	draw, _ := c.GetInt("draw")
	start, _ := c.GetInt64("start")
	length, _ := c.GetInt("length")
	q := strings.TrimSpace(c.GetString("q"))

	// []*User, total, queryTotal
	users, total, queryTotal := models.DefaultUserManager.Query(q, start, length, c.User)

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

// user/create
func (c *UserController) Create() {
	json := map[string]interface{}{
		"code":   405,
		"text":   "请求方式错误",
		"result": nil,
	}

	if c.Ctx.Input.IsPost() {
		form := &forms.UserCreateForm{}
		valid := &validation.Validation{}

		json["code"], json["text"] = 403, "没有权限"
		if c.User.IsSuperman == true {
			json["code"], json["text"] = 400, "请求数据错误"
			if err := c.ParseForm(form); err != nil {
				json["text"] = err.Error()
			} else {
				if ok, err := valid.Valid(form); err != nil {
					json["text"] = err.Error()
				} else if !ok {
					json["result"] = valid.Errors
				} else {
					birthday, _ := time.Parse("2006-01-02", form.Birthday)

					user := &models.User{
						Name:     form.Name,
						Password: form.Password,
						Gender:   form.Gender,
						Birthday: &birthday,
						Tel:      form.Tel,
						Email:    form.Email,
						Addr:     form.Addr,
						Remark:   form.Remark,
					}

					user.SetPassword(form.Password)
					if result, err := orm.NewOrm().Insert(user); err == nil {
						json["code"], json["text"], json["result"] = 200, "创建成功", result
					} else {
						json["code"], json["text"], json["result"] = 500, "服务器端错误", err.Error()
					}
				}
			}
		}
		c.Data["json"] = json
		c.ServeJSON()
	}
	c.TplName = "user/create.html"

}

// user/modify
func (c *UserController) Modify() {
	json := map[string]interface{}{
		"code":   405,
		"text":   "请求方式错误",
		"result": nil,
	}

	if c.Ctx.Input.IsPost() {
		form := &forms.UserModifyForm{}
		valid := &validation.Validation{}

		json["code"], json["text"] = 403, "没有权限"
		if c.User.IsSuperman == true {
			json["code"], json["text"] = 400, "请求数据错误"
			if err := c.ParseForm(form); err != nil {
				json["text"] = err.Error()
			} else {
				if ok, err := valid.Valid(form); err != nil {
					json["text"] = err.Error()
				} else if !ok {
					json["result"] = valid.Errors
				} else {
					birthday, _ := time.Parse("2006-01-02", form.Birthday)
					form.User.Name = form.Name
					form.User.Birthday = &birthday
					form.User.Gender = form.Gender
					form.User.Tel = form.Tel
					form.User.Email = form.Email
					form.User.Addr = form.Addr
					form.User.Remark = form.Remark

					// 更新数据库
					if result, err := orm.NewOrm().Update(form.User); err == nil {
						json["code"], json["text"], json["result"] = 200, "修改用户成功", result
					} else {
						json["code"], json["text"], json["result"] = 500, "服务器端错误", err.Error()
					}
				}
			}
		}
		c.Data["json"] = json
		c.ServeJSON()
	} else {
		pk, _ := c.GetInt("pk")
		c.TplName = "user/modify.html"
		c.Data["object"] = models.DefaultUserManager.GetById(pk)
	}
}

// user/delete
func (c *UserController) Delete() {
	json := map[string]interface{}{
		"code":   400,
		"text":   "请求数据错误",
		"result": nil,
	}
	if c.User.IsSuperman == true {
		if pk, err := c.GetInt("pk"); err != nil {
			json["text"] = err.Error()
		} else {
			if result, err := models.DefaultUserManager.DeleteById(pk); err == nil {
				json["code"], json["text"], json["result"] = 200, "用户删除成功", result
			} else {
				json["code"], json["text"], json["result"] = 500, "服务器端错误", err.Error()
			}
		}
	} else {
		json["code"], json["text"] = 403, "没有权限"
	}
	c.Data["json"] = json
	c.ServeJSON()
}

// user/lock
func (c *UserController) Lock() {
	json := map[string]interface{}{
		"code":   400,
		"text":   "请求数据错误",
		"result": nil,
	}
	if c.User.IsSuperman == true {
		if pk, err := c.GetInt("pk"); err != nil {
			json["text"] = err.Error()
		} else {
			if err := models.DefaultUserManager.SetStatusById(pk, 1); err == nil {
				json["code"], json["text"], json["result"] = 200, "用户锁定成功", nil
			} else {
				json["code"], json["text"], json["result"] = 500, "服务器端错误", err.Error()
			}
		}
	} else {
		json["code"], json["text"] = 403, "没有权限"
	}
	c.Data["json"] = json
	c.ServeJSON()
}

// user/unlock
func (c *UserController) UnLock() {
	json := map[string]interface{}{
		"code":   400,
		"text":   "请求数据错误",
		"result": nil,
	}
	if c.User.IsSuperman == true {
		if pk, err := c.GetInt("pk"); err != nil {
			json["text"] = err.Error()
		} else {
			if err := models.DefaultUserManager.SetStatusById(pk, 0); err == nil {
				json["code"], json["text"], json["result"] = 200, "用户锁定成功", nil
			} else {
				json["code"], json["text"], json["result"] = 500, "服务器端错误", err.Error()
			}
		}
	} else {
		json["code"], json["text"] = 403, "没有权限"
	}
	c.Data["json"] = json
	c.ServeJSON()
}

func (c *UserController) Password() {
	json := map[string]interface{}{
		"code":   405,
		"text":   "请求方式错误",
		"result": nil,
	}

	if c.Ctx.Input.IsPost() {
		form := &forms.ModifyPasswordForm{User: c.User}
		valid := &validation.Validation{}

		json["code"], json["text"] = 400, "请求数据错误"
		if err := c.ParseForm(form); err != nil {
			json["text"] = err.Error()
		} else {
			if ok, err := valid.Valid(form); err != nil {
			} else if !ok {
				json["result"] = valid.Errors
			} else {
				c.User.SetPassword(form.NewPassword)
				if result, err := orm.NewOrm().Update(c.User, "Password"); err == nil {
					json["code"], json["text"], json["result"] = 200, "密码修改成功", result
				} else {
					json["code"], json["text"], json["result"] = 500, "服务器端错误", err.Error()
				}
			}
		}
		c.Data["json"] = json
		c.ServeJSON()
	}
	c.TplName = "user/password.html"
}

type TokenController struct {
	auth.LoginRequireController
}

// token/generate
func (c *TokenController) Generate() {
	pk, _ := c.GetInt("pk")
	if c.Ctx.Input.IsPost() {
		if c.User.IsSuperman == true || c.User.Id == pk {
			models.DefaultTokenManager.GenerateByUser(models.DefaultUserManager.GetById(pk))
			c.Data["json"] = map[string]interface{}{
				"code":   200,
				"text":   "生成Token成功",
				"result": nil,
			}
		} else {
			c.Data["json"] = map[string]interface{}{
				"code":   403,
				"text":   "没有权限",
				"result": nil,
			}
		}
		c.ServeJSON()
	}
	c.Data["object"] = models.DefaultUserManager.GetById(pk)
	c.TplName = "token/index.html"

}
