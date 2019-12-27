package controllers

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/dcosapp/gocmdb/server/controllers/auth"
	"github.com/dcosapp/gocmdb/server/forms"
	"github.com/dcosapp/gocmdb/server/models"
	"strings"
)

type AlarmPageController struct {
	LayoutController
}

func (c *AlarmPageController) Index() {
	c.Data["expand"] = "alarm_management"
	c.Data["menu"] = "view_alarm"

	c.TplName = "view_alarm_page/index.html"
	c.LayoutSections["LayoutScript"] = "view_alarm_page/index.script.html"
}

type AlarmController struct {
	auth.LoginRequireController
}

func (c *AlarmController) List() {
	// draw, start, length, q
	draw, _ := c.GetInt("draw")
	start, _ := c.GetInt64("start")
	length, _ := c.GetInt("length")
	q := strings.TrimSpace(c.GetString("q"))

	// []*Alarm, total, queryTotal
	result, total, queryTotal := models.DefaultAlarmManager.Query(q, start, length)
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

func (c *AlarmController) Dealed() {
	json := map[string]interface{}{
		"code":   400,
		"text":   "请求数据错误",
		"result": nil,
	}

	if id, err := c.GetInt("pk"); err == nil {
		err := models.DefaultAlarmManager.SetStatusById(id, 1)
		if err == nil {
			json["code"], json["text"], json["result"] = 200, "告警开始处理", nil
		} else {
			json["code"], json["text"], json["result"] = 500, "服务器端错误", err.Error()
		}
	}

	c.Data["json"] = json
	c.ServeJSON()
}

func (c *AlarmController) Complete() {
	json := map[string]interface{}{
		"code":   400,
		"text":   "请求数据错误",
		"result": nil,
	}

	if id, err := c.GetInt("pk"); err == nil {
		err := models.DefaultAlarmManager.SetStatusById(id, 2)
		if err == nil {
			json["code"], json["text"], json["result"] = 200, "告警处理完成", nil
		} else {
			json["code"], json["text"], json["result"] = 500, "服务器端错误", err.Error()
		}
	}

	c.Data["json"] = json
	c.ServeJSON()
}

type AlarmSettingPageController struct {
	LayoutController
}

func (c *AlarmSettingPageController) Index() {
	c.Data["expand"] = "alarm_management"
	c.Data["menu"] = "alarm_setting"

	c.TplName = "alarm_setting/index.html"
	c.LayoutSections["LayoutScript"] = "alarm_setting/index.script.html"
}

type AlarmSettingController struct {
	auth.LoginRequireController
}

func (c *AlarmSettingController) List() {
	// draw, start, length, q
	draw, _ := c.GetInt("draw")
	start, _ := c.GetInt64("start")
	length, _ := c.GetInt("length")
	q := strings.TrimSpace(c.GetString("q"))

	// []*Alarm, total, queryTotal
	result, total, queryTotal := models.DefaultAlarmSettingManager.Query(q, start, length)
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

func (c *AlarmSettingController) Modify() {
	json := map[string]interface{}{
		"code":   405,
		"text":   "请求方式错误",
		"result": nil,
	}

	if c.Ctx.Input.IsPost() {
		form := &forms.AlarmSettingForms{}
		valid := &validation.Validation{}

		json["code"], json["text"] = 400, "请求数据错误"
		if err := c.ParseForm(form); err != nil {
			json["text"] = err.Error()
			fmt.Println(1)
		} else {
			if ok, err := valid.Valid(form); err != nil {
				json["text"] = err.Error()
			} else if !ok {
				json["result"] = valid.Errors
			} else {
				// 更新数据库
				if result, err := models.DefaultAlarmSettingManager.Modify(form.Id, form.Time, form.Threshold, form.Counter); err == nil {
					json["code"], json["text"], json["result"] = 200, "修改告警设置成功", result
				} else {
					json["code"], json["text"], json["result"] = 500, "服务器端错误", err.Error()
				}
			}
		}

		c.Data["json"] = json
		c.ServeJSON()
	}

	pk, _ := c.GetInt("pk")
	c.TplName = "alarm_setting/modify.html"
	c.Data["object"] = models.DefaultAlarmSettingManager.GetById(pk)
}
