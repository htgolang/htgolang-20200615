package controllers

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/xlotz/gocmdb/server/controllers/auth"
	"strings"

	"github.com/xlotz/gocmdb/server/models"
	"github.com/xlotz/gocmdb/server/forms"
	"github.com/xlotz/gocmdb/server/cloud"
	//_ "github.com/xlotz/gocmdb/server/cloud/plugins"
	//"strings"
)

type CloudPlatformPageController struct {
	LayoutController
}

func (c *CloudPlatformPageController) Index() {
	c.Data["menu"] = "platform_management"
	c.Data["expand"] = "cloud_management"
	c.TplName = "cloud/index.html"
	c.LayoutSections["LayoutScript"] = "cloud/index.script.html"
}

type CloudPlatformController struct {
	auth.LoginRequiredController
}

func (c *CloudPlatformController) List() {
	//draw,start, length, q
	draw, _ := c.GetInt("draw")
	start, _ := c.GetInt64("start")
	length, _ := c.GetInt("length")
	q := strings.TrimSpace(c.GetString("q"))

	// []*User, total, queryTotal
	result, total, queryTotal := models.DefaultCloudPlatformManager.Query(q, start, length)


	c.Data["json"] = map[string]interface{}{
		"code":            200,
		"text":            "获取成功",
		"result":          result,
		"draw":            draw,
		"recordsTotal":    total,
		"recordsFiltered": queryTotal,
	}
	c.ServeJSON()
}
//
func (c *CloudPlatformController) Create() {

	if c.Ctx.Input.IsPost() {

		json := map[string]interface{}{
			"code": 400,
			"text": "提交数据错误",
		}

		form := &forms.CloudPlatCreateForm{}
		valid := &validation.Validation{}


		if err := c.ParseForm(form); err != nil {
			valid.SetError("error", err.Error())
			json["result"] = valid.Errors

		}else {
			if ok, err := valid.Valid(form); err != nil {
				valid.SetError("error", err.Error())
				json["result"] = valid.Errors

			}else if ok {

				result, err := models.DefaultCloudPlatformManager.Create(
					form.Name,
					form.Type,
					form.Addr,
					form.Region,
					form.AccessKey,
					form.SecrectKey,
					form.Remark,
					c.User,
					)

				if err == nil {
					json = map[string]interface{}{
						"code": 200,
						"result": result,
						"text": "创建成功",
					}
				} else {
					json = map[string]interface{}{
						"code": 500,
						"text": "服务器错误",
						"result": nil,
					}
				}

			}else {
				json = map[string]interface{}{
					"code": 500,
					"result": nil,
				}

			}
		}

		c.Data["json"] = json
		c.ServeJSON()
	} else {
		//get
		c.TplName = "cloud/create.html"
		c.Data["types"] = cloud.DefaultManager.Plugins
	}
}

func (c *CloudPlatformController) Modify() {
	if c.Ctx.Input.IsPost() {
		json := map[string]interface{}{
			"code": 400,
			"text": "提交数据错误",
		}
		form := &forms.CloudPlatModifyForm{}
		valid := &validation.Validation{}
		if err := c.ParseForm(form); err == nil {
			if ok, err := valid.Valid(form); err != nil {
				valid.SetError("error", err.Error())
				json["result"] = valid.Errors
			} else if ok {


				obj := models.DefaultCloudPlatformManager.GetById(form.Id)


				if form.AccessKey == ""  || form.SecrectKey == "" {
					form.AccessKey = obj.AccessKey
					form.SecrectKey = obj.SecrectKey
				}


				result, err := models.DefaultCloudPlatformManager.Modify(
					form.Id,
					form.Name,
					form.Type,
					form.Addr,
					form.Region,
					form.AccessKey,
					form.SecrectKey,
					form.Remark,

					)


				if err == nil {
					json = map[string]interface{}{
						"code":   200,
						"text":   "更新成功",
						"result": result,
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

		c.TplName = "cloud/modify.html"
		c.Data["object"] = models.DefaultCloudPlatformManager.GetById(pk)
		c.Data["types"] = cloud.DefaultManager.Plugins
	}
}

func (c *CloudPlatformController) Delete(){
	if c.Ctx.Input.IsPost(){

		pk, _ := c.GetInt("pk")

		models.DefaultCloudPlatformManager.DeleteById(pk)

		c.Data["json"] = map[string]interface{}{
			"code": 200,
			"text": "删除成功",
			"result": nil,
		}

	}else {
		c.Data["json"] = map[string]interface{}{
			"code": 400,
			"text": "删除失败",
			"result": nil,
		}
	}

	c.ServeJSON()

}

func (c *CloudPlatformController) Lock() {
	pk, _ := c.GetInt("pk")
	fmt.Println(pk)
	models.DefaultCloudPlatformManager.SetStatusById(pk, 1)
	c.Data["json"] = map[string]interface{}{
		"code":   200,
		"text":   "已禁用",
		"result": nil,
	}

	c.ServeJSON()
}

func (c *CloudPlatformController) UnLock() {
	pk, _ := c.GetInt("pk")

	models.DefaultCloudPlatformManager.SetStatusById(pk, 0)
	c.Data["json"] = map[string]interface{}{
		"code":   200,
		"text":   "已开启",
		"result": nil,
	}

	c.ServeJSON()
}

// 云主机管理

func (c *CloudPlatformPageController) VirtualIndex() {
	c.Data["menu"] = "platform_management"
	c.Data["expand"] = "cloud_management"
	c.TplName = "virtual/index.html"
	c.LayoutSections["LayoutScript"] = "virtual/index.script.html"
}




type VirtualMachineformController struct {
	auth.LoginRequiredController
}

func (c *VirtualMachineformController) List() {
	//draw,start, length, q
	draw, _ := c.GetInt("draw")
	start, _ := c.GetInt64("start")
	length, _ := c.GetInt("length")
	q := strings.TrimSpace(c.GetString("q"))


	result, total, queryTotal := models.DefaultVirtualMachineManager.Query(q, start, length)


	c.Data["json"] = map[string]interface{}{
		"code":            200,
		"text":            "获取成功",
		"result":          result,
		"draw":            draw,
		"recordsTotal":    total,
		"recordsFiltered": queryTotal,
	}
	c.ServeJSON()
}