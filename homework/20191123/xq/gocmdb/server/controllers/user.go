package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"strings"
	"github.com/xlotz/gocmdb/server/controllers/auth"
	"github.com/xlotz/gocmdb/server/models"
	"github.com/xlotz/gocmdb/server/forms"
	"time"
)


type UserPageController struct {
	LayoutController
}

func (c *UserPageController) Index(){
	c.Data["menu"] = "user_management"
	c.Data["expand"] = "system_management"

	c.TplName = "user_page/index.html"

	c.LayoutSections["LayoutScript"] = "user_page/index_script.html"

}

type UserController struct {

	auth.LoginRequiredController

}

func (c *UserController) List() {
	draw, _ := c.GetInt("draw")
	start, _ := c.GetInt64("start")
	length, _:= c.GetInt("length")
	q := strings.TrimSpace(c.GetString("q"))

	// 用户的切片，total, querytotal
	users, total, queryTotal := models.DefaultUserManager.Query(q, start, length)

	c.Data["json"] = map[string]interface{}{
		"code": 200,
		"text": "获取数据成功",
		"result": users,
		"draw": draw,
		"recordsTotal": total,
		"recordsFiltered":queryTotal,

	}
	c.ServeJSON()


}

func (c *UserController) Create(){

	json := map[string]interface{}{
		"code": "405",
		"text": "请求错误",
		"result": nil,
	}

	if c.Ctx.Input.IsPost(){


		form := &forms.UserCreateForm{}

		fmt.Println(form.Name)

		valid := &validation.Validation{}

		if err := c.ParseForm(form); err != nil {
			json["code"] = 400
			json["text"] = err.Error()
		}else{
			if corret, err := valid.Valid(form); err != nil {
				json["code"] = 400
				json["text"] = err.Error()

			}else if !corret {

				json["result"] = valid.Errors
			}else {


				// 转换时间
				birthday, _ := time.Parse("1/2/2006", form.Birthday)



				// 创建结构体对象
				user := &models.User{
					Name: form.Name,
					Birthday: &birthday,
					Gender: form.Gender,
					Tel: form.Tel,
					Addr: form.Addr,
					Desc: form.Desc,
					Email: form.Email,
					Remark: form.Remark,
				}

				//设置密码
				user.SetPassword(form.Password)

				ormer := orm.NewOrm()

				if _, err := ormer.Insert(user); err == nil {

					json = map[string]interface{}{
						"code": 200,
						"text": "创建成功",
						"result": nil,// 返回已经创建的用户
					}

				}else {
					json = map[string]interface{}{
						"code": 500,
						"text": "服务端错误",
						"result": err.Error(),// 返回已经创建的用户
					}
					fmt.Println(json)

				}
			}
		}
		c.Data["json"] = json
		c.ServeJSON()

	}else {
		c.TplName = "user/create.html"
	}

}

func (c *UserController) Modify(){

	json := map[string]interface{}{
		"code": 405,
		"text": "请求方式错误",
		"result": nil,
	}

	if c.Ctx.Input.IsPost(){


		form := &forms.UserModifyForm{} // 修改表单
		valid := &validation.Validation{} //验证器

		fmt.Println(form)

		if err := c.ParseForm(form); err != nil {

			json["text"] = err.Error()


		}else {
			// 验证表单数据
			if corret, err := valid.Valid(form); err != nil {
				json["code"] = 400
				json["result"] = err.Error()

			}else if !corret {
				json["code"] = 400

			}else {

				// 更新用户信息
				birthday, _ := time.Parse("1/2/2006", form.Birthday)

				form.User.Name = form.Name
				form.User.Birthday = &birthday
				form.User.Gender = form.Gender
				form.User.Tel = form.Tel
				form.User.Addr = form.Addr
				form.User.Desc = form.Desc


				if _, err := orm.NewOrm().Update(form.User); err == nil {
					json["code"] = 200
					json["text"] = "修改成功"
					json["result"] = form.User

				}else {

					json["code"] = 500
					json["text"] = "服务器错误"
					json["result"] = nil
				}

			}
		}


		c.Data["json"] = json

		c.ServeJSON()

	}else {
		pk, _ := c.GetInt("pk")
		c.TplName = "user/modify.html"

		c.Data["object"] = models.DefaultUserManager.GetById(pk)
	}

}

func (c *UserController) Delete(){


	json := map[string]interface{}{
		"code":   405,
		"text":   "请求方式错误",
		"result": nil,
	}
	if c.Ctx.Input.IsPost() {
		if id, err := c.GetInt("pk"); err != nil {

			json["code"] = 400
			json["text"] = "获取ID失败"
			json["result"] = err.Error()

		}else{
			// 删除用户
			orm.NewOrm().Delete(&models.User{Id: id})

			// 通过flash提示用户操作结果
			json["code"] = 200
			json["text"] = "删除成功"
			json["result"] = nil

		}

	}

	c.Data["json"] = json

	c.ServeJSON()

}

func (c *UserController) Lock(){

	pk, _ := c.GetInt("pk")

	if err := models.DefaultUserManager.SetStatusById(pk,1); err == nil{

		c.Data["json"] = map[string]interface{}{
			"code": 200,
			"text": "锁定成功",
			"result": nil,// 返回已经创建的用户
		}


	}else {

		c.Data["json"] = map[string]interface{}{
			"code": 400,
			"text": "锁定失败",
			"result": nil,// 返回已经创建的用户
		}

	}



	c.ServeJSON()



}

func (c *UserController) UnLock(){

	pk, _ := c.GetInt("pk")

	models.DefaultUserManager.SetStatusById(pk, 0)

	c.Data["json"] = map[string]interface{}{
		"code": 200,
		"text": "解锁成功",
		"result": nil,// 返回已经创建的用户
	}

	c.ServeJSON()


}


// =============== Token

type TokenController struct {
	auth.LoginRequiredController
}

func (c *TokenController) Generate() {
	if c.Ctx.Input.IsPost(){
		pk, _:= c.GetInt("pk")
		fmt.Println(pk)
		models.DefaultTokenManager.GenerateByUser(models.DefaultUserManager.GetById(pk))

		c.Data["json"] = map[string]interface{}{
			"code": 200,
			"text": "生成Token成功",
			"result": nil,// 返回已经创建的用户
		}

		c.ServeJSON()
	}else {
		pk, _ := c.GetInt("pk")
		c.Data["object"] = models.DefaultUserManager.GetById(pk)
		c.TplName = "token/index.html"
	}
}




