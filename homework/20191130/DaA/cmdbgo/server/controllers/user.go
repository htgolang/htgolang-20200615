package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/xxdu521/cmdbgo/server/controllers/auth"
	"github.com/xxdu521/cmdbgo/server/forms"
	"github.com/xxdu521/cmdbgo/server/models"
	"strings"
	"time"
)

//用户管理页面
type UserPageController struct {
	LayoutController
}
func (c *UserPageController) Index(){
	c.Data["expand"] = "system_management"
	c.Data["menu"] = "user_management"
	c.TplName = "user_page/index.html"
	c.LayoutSections["LayoutScript"] = "user_page/index_script.html"
}

//用户数据
type UserController struct {
	auth.LoginRequiredController
}
func (c *UserController) List(){
	//draw,start,length,q
	draw,_ := c.GetInt("draw")
	start,_ := c.GetInt64("start")
	length,_ := c.GetInt("length")

	//给API使用的最大返回数，用配置文件做调整。默认2000.
	Max_Query_Length,_ := beego.AppConfig.Int("Max_Query_Length")
	if Max_Query_Length > 10 && length > Max_Query_Length {
		length = Max_Query_Length
	}

	q := strings.TrimSpace(c.GetString("q"))

	//[]*User,total,querytotal
	users, total, querytotal := models.DefaultUserManager.Query(q, start, length)
	c.Data["json"] = map[string]interface{}{
		"code": 			200,
		"text": 			"成功",
		"result": 			users,
		"draw": 			draw,
		"recordsTotal": 	total,
		"recordsFiltered": 	querytotal,
	}
	c.ServeJSON()
}
func (c *UserController) Create(){
	if c.Ctx.Input.IsPost() {
		json := map[string]interface{}{
			"code": 400,
			"text": "提交数据错误",
			"result": nil,
		}

		form := &forms.UserCreateForm{}
		valid := &validation.Validation{}

		if err := c.ParseForm(form);err != nil {
			json["text"] = err.Error()
		} else {
			if corret, err := valid.Valid(form); err != nil {
				json["text"] = err.Error()
			} else if !corret {
				json["result"] = valid.Errors
			} else {
				birthday, _ := time.Parse("01/02/2006", form.Birthday)
				user := &models.User{
					Name:       form.Name,
					Birthday:   &birthday,
					Gender:     form.Gender,
					Tel:        form.Tel,
					Email:		form.Email,
					Addr:       form.Addr,
					Remark:     form.Remark,

				}
				user.SetPassword(form.Password)
				ormer := orm.NewOrm()
				if result, err := ormer.Insert(user); err == nil {
					json["code"], json["text"], json["result"] = 200, "编辑用户成功", result
				} else {
					json["code"], json["text"], json["result"] = 500, "服务器错误", err.Error()
				}
			}
		}
		c.Data["json"] = json
		c.ServeJSON()
	} else {
		c.TplName = "user/create.html"
	}
}
func (c *UserController) Modify(){
	if c.Ctx.Input.IsPost() {
		json := map[string]interface{}{
			"code": 400,
			"text": "提交数据错误",
			"result": nil,
		}

		form := &forms.UserModifyForm{}
		valid := &validation.Validation{}

		if err := c.ParseForm(form);err != nil {
			json["text"] = err.Error()
		} else {
			if corret,err := valid.Valid(form); err != nil {
				json["text"] = err.Error()
			} else if !corret {
				json["result"] = valid.Errors
			} else {
				birthday, _ := time.Parse("01/02/2006", form.Birthday)

				form.User.Name = form.Name
				form.User.Birthday = &birthday
				form.User.Gender = form.Gender
				form.User.Tel = form.Tel
				form.User.Email = form.Email
				form.User.Addr = form.Addr
				form.User.Remark = form.Remark

				ormer := orm.NewOrm()
				if result,err := ormer.Update(form.User);err == nil {
					json["code"], json["text"], json["result"] = 200, "编辑用户成功", result
				} else {
					json["code"], json["text"], json["result"] = 500, "服务器错误", err.Error()
				}
			}
		}
		c.Data["json"] = json
		c.ServeJSON()
	} else {
		id,_ := c.GetInt("id")
		c.Data["object"] = models.DefaultUserManager.GetById(id)
		c.TplName = "user/modify.html"
	}
}
func (c *UserController) Delete(){
	json := map[string]interface{}{
		"code":   405,
		"text":   "请求方法错误",
		"result": nil,
	}

	fmt.Println(c.User)
	id,_ := c.GetInt("id")

	if c.User.Id != id {
		if c.Ctx.Input.IsPost() {
			models.DefaultUserManager.DeleteById(id)
			json["code"], json["text"], json["result"] = 200, "删除成功", nil
		}
	} else {
		json["code"], json["text"], json["result"] = 403, "不能删除自己", nil
	}

	c.Data["json"] = json
	c.ServeJSON()
}
func (c *UserController) Lock(){
	json := map[string]interface{}{
		"code":   405,
		"text":   "请求方法错误",
		"result": nil,
	}

	id, _ := c.GetInt("id")
	if c.User.Id != id {
		if c.Ctx.Input.IsPost() {

			models.DefaultUserManager.SetStatusById(id, 1)
			json["code"], json["text"], json["result"] = 200, "锁定成功", nil
		}
	} else {
		json["code"], json["text"], json["result"] = 403, "不能对自己操作", nil
	}

	c.Data["json"] = json
	c.ServeJSON()
}
func (c *UserController) UnLock(){
	json := map[string]interface{}{
		"code":   405,
		"text":   "请求方法错误",
		"result": nil,
	}

	id, _ := c.GetInt("id")
	if c.User.Id != id {
		if c.Ctx.Input.IsPost() {
			models.DefaultUserManager.SetStatusById(id, 0)
			json["code"], json["text"], json["result"] = 200, "解锁成功", nil
		}
	} else {
		json["code"], json["text"], json["result"] = 403, "不能对自己操作", nil
	}

	c.Data["json"] = json
	c.ServeJSON()
}
func (c *UserController) SetPassword(){
	json := map[string]interface{}{
		"code":   405,
		"text":   "请求方法错误",
		"result": nil,
	}

	if c.Ctx.Input.IsPost() {
		json["code"], json["text"] = 400, "请求数据错误"

		form := &forms.UserSetPasswordForm{User:c.User}
		valid := &validation.Validation{}

		//判断用户的密码是否正确
		if err := c.ParseForm(form); err != nil {
			json["text"] = err.Error()
		} else {
			if corret, err := valid.Valid(form); err != nil {
				json["text"] = err.Error()
			} else if !corret {
				json["result"] = valid.Errors
				fmt.Println(valid.Errors)
			} else {
				c.User.SetPassword(form.NewPassword)

				ormer := orm.NewOrm()
				if result,err := ormer.Update(c.User,"Password"); err == nil {
					json["code"], json["text"], json["result"] = 200, "密码修改成功", result
				} else {
					json["code"], json["text"], json["result"] = 500, "服务器错误", err.Error()
				}
			}
		}

		c.Data["json"] = json
		c.ServeJSON()
	}
	c.TplName = "user/setpassword.html"
}

