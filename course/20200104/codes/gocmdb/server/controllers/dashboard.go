package controllers

import (
	"github.com/imsilence/gocmdb/server/controllers/auth"
	"github.com/imsilence/gocmdb/server/models"
)

type DashboardPageController struct {
	LayoutController
}

func (c *DashboardPageController) Index() {
	c.Data["menu"] = "dashboard"
	c.TplName = "dashboard_page/index.html"
	c.LayoutSections["LayoutScript"] = "dashboard_page/index.script.html"
}

type DashboardController struct {
	auth.LoginRequiredController
}

func (c *DashboardController) Stat() {
	onlineCnt, offlineCnt := models.DefaultAgentManager.GetStat()
	alarm_trend_days, alarm_trend_data := models.DefaultAlarmManager.GetLastestNStat(7)

	c.Data["json"] = map[string]interface{}{
		"code": 200,
		"text": "获取成功",
		"result": map[string]interface{}{
			"agent_offline_count":  offlineCnt,
			"agent_online_count":  onlineCnt,
			"alarm_count": models.DefaultAlarmManager.GetCountForNoComplete(),
			"alarm_dist":  models.DefaultAlarmManager.GetStatForNotComplete(),
			"alarm_trend": map[string]interface{} {
				"days" : alarm_trend_days,
				"data" : alarm_trend_data,
			},
		},
	}
	c.ServeJSON()
}
