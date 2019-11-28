package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/xxdu521/cmdbgo/server/controllers/auth"
	"github.com/xxdu521/cmdbgo/server/forms"
	"github.com/xxdu521/cmdbgo/server/models"
	"strings"
	"time"
)

type UserPageController struct {
	LayoutController
}

func (c *UserPageController) Index(){
	c.Data["expand"] = "system_management"
	c.Data["menu"] = "user_management"
	c.TplName = "user_page/index.html"
	c.LayoutSections["LayoutScript"] = "user_page/index_script.html"
}


type UserController struct {
	auth.LoginRequiredController
}

func (c *UserController) List(){
	//draw,start,length,q
	draw,_ := c.GetInt("draw")
	start,_ := c.GetInt64("start")
	length,_ := c.GetInt("length")
	Max_Query_Length,_ := beego.AppConfig.Int("Max_Query_Length")
	if Max_Query_Length > 10 && length > Max_Query_Length {
		length = Max_Query_Length
	}
	q := strings.TrimSpace(c.GetString("q"))

	//[]*User,total,querytotal
	users,total,querytotal := models.DefaultUserManager.Query(q,start,length)
	c.Data["json"] = map[string]interface{}{
		"code": 200,
		"text": "成功",
		"result": users,
		"draw": draw,
		"recordsTotal": total,
		"recordsFiltered": querytotal,
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
				birthday, _ := time.ParseInLocation("01/02/2006", form.Birthday,time.Local)

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
				if _,err := ormer.Insert(user); err == nil {
					json = map[string]interface{}{
						"code":   200,
						"text":   "创建用户成功",
						"result": user,
					}
				} else {
					json = map[string]interface{}{
						"code": 500,
						"text": "服务端错误",
						"result": nil,
					}
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

				birthday, _ := time.ParseInLocation("2006-01-02", form.Birthday,time.Local)

				user := &models.User{
					Name:       form.Name,
					Birthday:   &birthday,
					Gender:     form.Gender,
					Tel:        form.Tel,
					Email:		form.Email,
					Addr:       form.Addr,
					Remark:     form.Remark,
				}

				ormer := orm.NewOrm()
				if _,err := ormer.Update(user);err == nil {
					json = map[string]interface{}{
						"code": 200,
						"text": "编辑用户成功",
						"result": user,
					}
				} else {
					json = map[string]interface{}{
						"code": 500,
						"text": "服务端错误",
						"result": nil,
					}
				}
			}
		}
		c.Data["json"] = json
		c.ServeJSON()
	} else {
		id,_ := c.GetInt("id")
		c.TplName = "user/modify.html"
		user := models.DefaultUserManager.GetById(id)
		c.Data["object"] = user
	}
}

func (c *UserController) Delete(){
	id,_ := c.GetInt("id")
	models.DefaultUserManager.DeleteById(id)

	json := map[string]interface{}{
		"code":   200,
		"text":   "编辑成功",
		"result": nil,
	}
	c.Data["json"] = json
	c.ServeJSON()
}

func (c *UserController) Lock(){
	id,_ := c.GetInt("id")
	models.DefaultUserManager.SetStatusById(id,1)

	json := map[string]interface{}{
		"code":   200,
		"text":   "编辑成功",
		"result": nil,
	}
	c.Data["json"] = json
	c.ServeJSON()
}

func (c *UserController) UnLock(){
	id,_ := c.GetInt("id")
	models.DefaultUserManager.SetStatusById(id,0)

	json := map[string]interface{}{
		"code":   200,
		"text":   "编辑成功",
		"result": nil,
	}
	c.Data["json"] = json
	c.ServeJSON()
}
