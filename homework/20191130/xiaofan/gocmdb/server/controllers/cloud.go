package controllers

import (
	"github.com/astaxie/beego/validation"
	"gocmdb/cloud"
	"gocmdb/controllers/auth"
	"gocmdb/forms"
	"gocmdb/models"
	"strings"
)

// cloudplatformpage/
type CloudPlatformPageController struct {
	LayoutController
}

// cloudplatformpage/index
func (c *CloudPlatformPageController) Index() {
	c.Data["expand"] = "cloud_management"
	c.Data["menu"] = "cloud_platform_management"

	c.TplName = "cloud_platform_page/index.html"
	c.LayoutSections["LayoutScript"] = "cloud_platform_page/index.script.html"
}

// cloudplatform/
type CloudPlatformController struct {
	auth.LoginRequireController
}

// cloudplatform/list
func (c *CloudPlatformController) List() {
	// draw, start, length, q
	draw, _ := c.GetInt("draw")
	start, _ := c.GetInt64("start")
	length, _ := c.GetInt("length")
	q := strings.TrimSpace(c.GetString("q"))

	// []*CloudPlatform, total, queryTotal
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

// cloudplatform/create
func (c *CloudPlatformController) Create() {
	json := map[string]interface{}{
		"code":   405,
		"text":   "请求方式错误",
		"result": nil,
	}

	if c.Ctx.Input.IsPost() {
		form := &forms.CloudPlatformCreateForm{}
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
					if result, err := models.DefaultCloudPlatformManager.Create(form.Name, form.Type, form.Addr, form.AccessKey,
						form.SecretKey, form.Region, form.Remark, c.User); err == nil {
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
	c.TplName = "cloud_platform/create.html"
	c.Data["types"] = cloud.DefaultManager.Plugins // Plugins map[string]ICloud
}

// cloudplatform/modify
func (c *CloudPlatformController) Modify() {
	json := map[string]interface{}{
		"code":   405,
		"text":   "请求方式错误",
		"result": nil,
	}

	if c.Ctx.Input.IsPost() {
		form := &forms.CloudPlatformModifyForm{}
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
					// 更新数据库
					if result, err := models.DefaultCloudPlatformManager.Modify(form.Id, form.Name, form.Type, form.Addr,
						form.Region, form.AccessKey, form.SecretKey, form.Remark); err == nil {
						json["code"], json["text"], json["result"] = 200, "修改用户成功", result
					} else {
						json["code"], json["text"], json["result"] = 500, "服务器端错误", err.Error()
					}
				}
			}
		}
		c.Data["json"] = json
		c.ServeJSON()
	}

	pk, _ := c.GetInt("pk")
	c.TplName = "cloud_platform/modify.html"
	c.Data["object"] = models.DefaultCloudPlatformManager.GetById(pk)
	c.Data["types"] = cloud.DefaultManager.Plugins // Plugins map[string]ICloud
}

// cloudplatform/lock
func (c *CloudPlatformController) Disable() {
	json := map[string]interface{}{
		"code":   400,
		"text":   "请求数据错误",
		"result": nil,
	}
	if c.User.IsSuperman == true {
		if pk, err := c.GetInt("pk"); err != nil {
			json["text"] = err.Error()
		} else {
			if err := models.DefaultCloudPlatformManager.SetStatusById(pk, 1); err == nil {
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

// cloudplatform/unlock
func (c *CloudPlatformController) Enable() {
	json := map[string]interface{}{
		"code":   400,
		"text":   "请求数据错误",
		"result": nil,
	}
	if c.User.IsSuperman == true {
		if pk, err := c.GetInt("pk"); err != nil {
			json["text"] = err.Error()
		} else {
			if err := models.DefaultCloudPlatformManager.SetStatusById(pk, 0); err == nil {
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

// cloudplatform/delete
func (c *CloudPlatformController) Delete() {
	json := map[string]interface{}{
		"code":   400,
		"text":   "请求数据错误",
		"result": nil,
	}
	if c.User.IsSuperman == true {
		if pk, err := c.GetInt("pk"); err != nil {
			json["text"] = err.Error()
		} else {
			if result, err := models.DefaultCloudPlatformManager.DeleteById(pk); err == nil {
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

// virtualmachinepage/
type VirtualMachinePageController struct {
	LayoutController
}

// virtualmachinepage/index
func (c *VirtualMachinePageController) Index() {
	c.Data["expand"] = "cloud_management"
	c.Data["menu"] = "virtual_machine_management"
	c.TplName = "virtual_machine_page/index.html"
	c.LayoutSections["LayoutScript"] = "virtual_machine_page/index.script.html"
}

type VirtualMachineController struct {
	auth.LoginRequireController
}

// virtualmachine/list
func (c *VirtualMachineController) List() {
	// draw, start, length, q
	draw, _ := c.GetInt("draw")
	start, _ := c.GetInt64("start")
	length, _ := c.GetInt("length")
	q := strings.TrimSpace(c.GetString("q"))

	// []*VirtualMachine, total, queryTotal
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

func (c *VirtualMachineController) Stop() {
	c.Data["json"] = map[string]interface{}{
		"code":   200,
		"text":   "关机成功",
		"result": nil,
	}
	c.ServeJSON()
}

func (c *VirtualMachineController) Start() {
	c.Data["json"] = map[string]interface{}{
		"code":   200,
		"text":   "开机成功",
		"result": nil,
	}
	c.ServeJSON()
}

func (c *VirtualMachineController) Restart() {
	c.Data["json"] = map[string]interface{}{
		"code":   200,
		"text":   "重启成功",
		"result": nil,
	}
	c.ServeJSON()
}
