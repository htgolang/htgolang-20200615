package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"net/http"
	"strings"
	"time"
	"todolist/forms"
	"todolist/models"
	"todolist/utils"
)

type UserController struct {
	LoginController
}

//用户管理模块，只允许普通用户改密码
func (c *UserController) Prepare(){
	c.LoginController.Prepare()
	_, action := c.GetControllerAndAction()
	//除了user/password外，其他url只有超管可以操作
	if action != "Password" && !c.User.IsSuper {
		//非管理员并且不是请求user/password的跳到主页面
		c.Redirect(beego.URLFor(beego.AppConfig.String("home")),http.StatusFound)
	}

	c.Layout = "layout/base.html"
	c.Data["nav"] = "user"
}

func (c *UserController) Index(){
	q := strings.TrimSpace(c.GetString("q"))
	var users []models.User

	condition := orm.NewCondition()
	if q != "" {
		condition = condition.Or("name__icontains",q)
		condition = condition.Or("tel__icontains",q)
		condition = condition.Or("addr_icontains",q)
		condition = condition.Or("desc_icontains",q)
	}

	orm.NewOrm().QueryTable(&models.User{}).SetCond(condition).All((&users))

	c.TplName = "user/index.html"
	c.Data["users"] = users
	c.Data["q"] = q
}

func (c *UserController) Create(){
	form := &forms.UserCreateForm{}
	valid := &validation.Validation{}

	if c.Ctx.Input.IsPost() {
		if c.ParseForm(form) == nil {
			if corret, err := valid.Valid(form);err == nil && corret {
				birthday, _ := time.Parse("2006-01-02", form.Birthday)

				user := &models.User{
					Name:       form.Name,
					Birthday:   &birthday,
					Sex:        form.Sex == 1,
					Tel:        form.Tel,
					Addr:       form.Addr,
					Desc:       form.Desc,
				}

				user.SetPassword(form.Password)

				ormer := orm.NewOrm()
				ormer.Insert(user)

				flash := beego.NewFlash()
				flash.Success("添加用户成功")
				flash.Store(&c.Controller)
				c.Redirect(beego.URLFor("UserController.Index"),http.StatusFound)
			}
		}
	}

	c.TplName = "user/create.html"
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["form"] = form
	c.Data["validation"] = valid

}

func (c *UserController) Modify(){
	form := &forms.UserModifyForm{}
	valid := &validation.Validation{}

	if c.Ctx.Input.IsPost() {
		if c.ParseForm(form) == nil {
			if corret,err := valid.Valid(form);err == nil && corret {
				birthday, _ := time.Parse("2006-01-02",form.Birthday)
				form.User.Name = form.Name
				form.User.Birthday = &birthday
				form.User.Sex = form.Sex == 1
				form.User.Tel = form.Tel
				form.User.Addr = form.Addr
				form.User.Desc = form.Desc
			}

			ormer := orm.NewOrm()
			ormer.Update(form.User)

			flash := beego.NewFlash()
			flash.Success("修改用户成功")
			flash.Store(&c.Controller)
			c.Redirect(beego.URLFor("UserController.Index"),http.StatusFound)
		}
	}
	if c.Ctx.Input.IsGet() {
		if id,err := c.GetInt("id");err == nil {
			user := models.User{Id:id}
			if orm.NewOrm().Read(&user) == nil {
				form.Id = user.Id
				form.Name = user.Name
				if user.Sex {
					form.Sex = 1
				}
				if user.Birthday != nil {
					form.Birthday = user.Birthday.Format("2006-01-02")
				}
				form.Addr = user.Addr
				form.Tel = user.Tel
				form.Desc = user.Desc
			}
		}
	}

	c.TplName = "user/modify.html"
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["form"] = form
	c.Data["validation"] = valid
}

func (c *UserController) Delete(){
	if id,err := c.GetInt("id");err == nil {
		orm.NewOrm().Delete(&models.User{Id:id})

		flash := beego.NewFlash()
		flash.Success("删除用户成功")
		flash.Store(&c.Controller)
	}
	c.Redirect(beego.URLFor("UserController.Index"),http.StatusFound)
}

func (c *UserController) Password(){
	form := &forms.ModifyPasswordForm{User: c.User}
	valid := &validation.Validation{}

	if c.Ctx.Input.IsPost() {
		if c.ParseForm(form) == nil {
			if corret,err := valid.Valid(form);err == nil && corret {
				c.User.SetPassword(form.NewPassword)
				ormer := orm.NewOrm()
				ormer.Update(c.User,"Password")

				flash := beego.NewFlash()
				flash.Success("修改密码成功")
				c.Data["flash"] = flash.Data
			}
		}
	}

	c.TplName = "user/password.html"
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["form"] = form
	c.Data["validation"] = valid
}

func (c *UserController) ResetPassword(){
	if id,err := c.GetInt("id");err ==nil && c.User.IsSuper {
		user := models.User{Id:id}
		ormer := orm.NewOrm()
		if err := ormer.Read(&user);err == nil {
			password := utils.RandString(6)
			user.SetPassword(password)
			ormer.Update(&user,"Password")

			flash := beego.NewFlash()
			flash.Success("重用用户<%s>密码为<%s>",user.Name,password)
			flash.Store(&c.Controller)
		}
	}
	c.Redirect(beego.URLFor("UserController.Index"),http.StatusFound)
}


