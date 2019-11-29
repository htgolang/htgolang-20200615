package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/imsilence/gocmdb/server/controllers/auth"
	"github.com/imsilence/gocmdb/server/forms"
	"github.com/imsilence/gocmdb/server/models"
)

type UserPageController struct {
	LayoutController
}

func (c *UserPageController) Index() {
	c.Data["menu"] = ""
	c.Data["expand"] = "system_management"
	c.TplName = "user_page/index.html"
	c.LayoutSections["LayoutScript"] = "user_page/index.script.html"
}

type UserController struct {
	auth.LoginRequiredController
}

func (c *UserController) List() {
	// draw,start,length,q
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
			"code":   400,
			"text":   "提交数据错误",
			"result": nil,
		}
		form := &forms.UserCreateForm{}
		valid := &validation.Validation{}

		if err := c.ParseForm(form); err != nil {
			json["text"] = err.Error()
		} else {
			// 表单验证
			if corret, err := valid.Valid(form); err != nil {
				json["text"] = err.Error()
			} else if !corret {
				json["result"] = valid.Errors
			} else {
				// 创建用户
				birthday, _ := time.Parse("2006-01-02", form.Birthday)

				user := &models.User{
					Name:     form.Name,
					Birthday: &birthday,
					Gender:   form.Gender,
					Tel:      form.Tel,
					Addr:     form.Addr,
					Remark:   form.Remark,
					Email:    form.Email,
				}
				user.SetPassword(form.Password)

				ormer := orm.NewOrm()
				if _, err := ormer.Insert(user); err == nil {
					json = map[string]interface{}{
						"code":   200,
						"text":   "创建成功",
						"result": user,
					}
				} else {
					json = map[string]interface{}{
						"code":   500,
						"text":   "服务端错误",
						"result": nil,
					}
				}
			}
		}
		c.Data["json"] = json
		c.ServeJSON()
	} else {
		//get
		c.TplName = "user/create.html"
	}
}

func (c *UserController) Modify() {
	form := &forms.UserModifyForm{} // 任务修改表单
	valid := &validation.Validation{}

	if c.Ctx.Input.IsPost() {
		json := map[string]interface{}{
			"code":   400,
			"text":   "请求数据错误",
			"result": nil,
		}

		// 解析请求参数到form中(根据form标签)
		if err := c.ParseForm(form); err != nil {
			json["text"] = err.Error()
		} else {
			if corret, err := valid.Valid(form); err != nil {
				json["text"] = err.Error()
			} else if !corret {
				json["result"] = valid.Errors
			} else {
				// 验证任务信息（超级管理员可修改任意任务信息，普通管理员只能修改自己创建的任务信息）
				birthday, _ := time.Parse("2006-01-02", form.Birthday)
				form.User.Name = form.Name
				form.User.Gender = form.Gender
				form.User.Birthday = &birthday
				//form.User.Birthday = form.Birthday
				form.User.Tel = form.Tel
				form.User.Email = form.Email
				form.User.Addr = form.Addr
				form.User.Remark = form.Remark

				// 更新数据库
				if _, err := orm.NewOrm().Update(form.User); err == nil {
					json["code"] = 200
					json["text"] = "修改成功"
					json["result"] = form.User

				} else {
					fmt.Println(err)
					json["code"] = 500
					json["text"] = "服务器错误"
					json["result"] = nil
				}
			}
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

func (c *UserController) ModifyPassword() {
	form := &forms.ModifyPasswordForm{User: c.User}
	valid := &validation.Validation{}

	if c.Ctx.Input.IsPost() {
		json := map[string]interface{}{
			"code":   400,
			"text":   "请求数据错误",
			"result": nil,
		}
		if err := c.ParseForm(form); err != nil {
			json["text"] = err.Error()
		} else {
			if corret, err := valid.Valid(form); err != nil {
				json["text"] = err.Error()
			} else if !corret {
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
	} else {
		pk, _ := c.GetInt("pk")
		c.TplName = "user/modifypassword.html"
		c.Data["object"] = models.DefaultUserManager.GetById(pk)
	}

}

func (c *UserController) ResetPassword() {
	form := &forms.ResetPasswordForm{User: c.User}
	valid := &validation.Validation{}

	if c.Ctx.Input.IsPost() {
		json := map[string]interface{}{
			"code":   400,
			"text":   "请求数据错误",
			"result": nil,
		}
		if err := c.ParseForm(form); err != nil {
			json["text"] = err.Error()
		} else {
			if corret, err := valid.Valid(form); err != nil {
				json["text"] = err.Error()
			} else if !corret {
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
	} else {
		pk, _ := c.GetInt("pk")
		c.TplName = "user/resetpassword.html"
		c.Data["object"] = models.DefaultUserManager.GetById(pk)
	}
}

type TokenController struct {
	auth.LoginRequiredController
}

func (c *TokenController) Generate() {
	if c.Ctx.Input.IsPost() {
		pk, _ := c.GetInt("pk")
		token := models.DefaultTokenManager.GenerateByUser(models.DefaultUserManager.GetById(pk))
		//models.DefaultTokenManager.GenerateByUser(models.DefaultUserManager.GetById(pk))
		// c.Data["json"] = map[string]interface{}{
		// 	"code":   200,
		// 	"text":   "生成token成功",
		// 	"result": token, //可以返回当前Token
		// }
		json := map[string]interface{}{
			"code":   200,
			"text":   "生成token成功",
			"result": token, //可以返回当前Token
		}
		c.Data["json"] = json
		c.ServeJSON()
	} else {
		pk, _ := c.GetInt("pk")
		c.Data["object"] = models.DefaultUserManager.GetById(pk)
		c.TplName = "token/index.html"
	}
}
