package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/xxdu521/cmdbgo/server/cloud"
	_ "github.com/xxdu521/cmdbgo/server/cloud/plugins"
	"github.com/xxdu521/cmdbgo/server/controllers/auth"
	"github.com/xxdu521/cmdbgo/server/forms"
	"github.com/xxdu521/cmdbgo/server/models"
	"strings"
)

//云平台页面
type CloudPlatformPageController struct {
	LayoutController
}
func (c *CloudPlatformPageController) Index(){
	c.Data["expand"] = "cloud_management"
	c.Data["menu"] = "cloud_platform_management"
	c.TplName = "cloud_platform_page/index.html"
	c.LayoutSections["LayoutScript"] = "cloud_platform_page/index_script.html"
}
//云平台数据
type CloudPlatformController struct {
	auth.LoginRequiredController
}
func (c *CloudPlatformController) List(){
	draw,_ := c.GetInt("draw")
	start,_ := c.GetInt64("start")
	length,_ := c.GetInt("length")

	Max_Query_Length,_ := beego.AppConfig.Int("Max_Query_Length")
	if Max_Query_Length > 10 && length > Max_Query_Length {
		length = Max_Query_Length
	}

	q := strings.TrimSpace(c.GetString("q"))

	result, total, querytotal := models.DefaultCloudPlatformManager.Query(q, start, length)

	c.Data["json"] = map[string]interface{}{
		"code": 			200,
		"text": 			"成功",
		"result": 			result,
		"draw": 			draw,
		"recordsTotal": 	total,
		"recordsFiltered": 	querytotal,
	}
	c.ServeJSON()
}
func (c *CloudPlatformController) Create(){
	if c.Ctx.Input.IsPost(){
		json := map[string]interface{}{
			"code": 400,
			"text": "提交数据错误",
			"result": nil,
		}

		form := &forms.CloudPlatformCreateForm{}
		valid := &validation.Validation{}

		if err := c.ParseForm(form); err != nil {
			valid.SetError("error", err.Error())
			json["result"] = valid.Errors
		} else {
			if corret, err := valid.Valid(form); err != nil {
				valid.SetError("error", err.Error())
				json["result"] = valid.Errors
			} else if !corret {
				json["result"] = valid.Errors
			} else {
				result,err := models.DefaultCloudPlatformManager.Create(
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
					json["code"], json["text"], json["result"] = 200, "创建成功", result
				} else {
					json["code"], json["text"], json["result"] = 500, "服务器错误", err.Error()
				}
			}
		}
		c.Data["json"] = json
		c.ServeJSON()
	} else {
		c.TplName = "cloud_platform/create.html"
		c.Data["types"] = cloud.DefaultManager.Plugins
	}
}
func (c *CloudPlatformController) Modify(){
	if c.Ctx.Input.IsPost(){
		json := map[string]interface{}{
			"code": 400,
			"text": "请求数据错误",
			"result": nil,
		}

		form := &forms.CloudPlatformModifyForm{}
		valid := &validation.Validation{}

		if err := c.ParseForm(form);err != nil {
			valid.SetError("error", err.Error())
			json["result"] = valid.Errors
		} else {
			if corret, err := valid.Valid(form); err != nil {
				valid.SetError("error", err.Error())
				json["result"] = valid.Errors
			} else if !corret {
				json["result"] = valid.Errors
			} else {
				result, err := models.DefaultCloudPlatformManager.Modify(
					form.Id,
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
					json["code"], json["text"], json["result"] = 200, "成功", result
				} else {
					json["code"], json["text"], json["result"] = 500, "服务端错误", nil
				}
			}
		}
		c.Data["json"] = json
		c.ServeJSON()
	} else {
		id,_ := c.GetInt("id")
		c.Data["object"] = models.DefaultCloudPlatformManager.GetById(id)
		c.Data["types"] = cloud.DefaultManager.Plugins
		c.TplName = "cloud_platform/modify.html"
	}


}
func (c *CloudPlatformController) Lock(){
	json := map[string]interface{}{
		"code": 405,
		"text": "请求方法错误",
		"result": nil,
	}

	id, _ := c.GetInt("id")
	if c.Ctx.Input.IsPost() {
		models.DefaultCloudPlatformManager.SetStatusById(id, 1)
		json["code"], json["text"], json["result"] = 200, "锁定成功", nil
	}

	c.Data["json"] = json
	c.ServeJSON()
}
func (c *CloudPlatformController) UnLock(){
	json := map[string]interface{}{
		"code":   405,
		"text":   "请求方法错误",
		"result": nil,
	}

	id, _ := c.GetInt("id")
	if c.Ctx.Input.IsPost() {
		models.DefaultCloudPlatformManager.SetStatusById(id, 0)
		json["code"], json["text"], json["result"] = 200, "启用成功", nil
	}

	c.Data["json"] = json
	c.ServeJSON()
}
func (c *CloudPlatformController) Delete(){
	json := map[string]interface{}{
		"code":   405,
		"text":   "请求方法错误",
		"result": nil,
	}

	id,_ := c.GetInt("id")

	if c.Ctx.Input.IsPost() {
		models.DefaultCloudPlatformManager.DeleteById(id)
		json["code"], json["text"], json["result"] = 200, "删除成功", nil
	}

	c.Data["json"] = json
	c.ServeJSON()
}

//云主机页面
type VirtualMachinePageController struct {
	LayoutController
}
func (c *VirtualMachinePageController) Index(){
	c.Data["expand"] = "cloud_management"
	c.Data["menu"] = "virtual_machine_management"
	c.TplName = "virtual_machine_page/index.html"
	c.LayoutSections["LayoutScript"] = "virtual_machine_page/index_script.html"
}

//云主机数据
type VirtualMachineController struct {
	auth.LoginRequiredController
}
func (c *VirtualMachineController) List(){
	draw,_ := c.GetInt("draw")
	start,_ := c.GetInt64("start")
	length,_ := c.GetInt("length")

	Max_Query_Length,_ := beego.AppConfig.Int("Max_Query_Length")
	if Max_Query_Length > 10 && length > Max_Query_Length {
		length = Max_Query_Length
	}

	q := strings.TrimSpace(c.GetString("q"))

	result, total, querytotal := models.DefaultVirtualMachineManager.Query(q, start, length)

	c.Data["json"] = map[string]interface{}{
		"code": 			200,
		"text": 			"成功",
		"result": 			result,
		"draw": 			draw,
		"recordsTotal": 	total,
		"recordsFiltered": 	querytotal,
	}
	c.ServeJSON()
}
