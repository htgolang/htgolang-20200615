package forms

import "github.com/astaxie/beego/validation"

type AlarmSettingForms struct {
	Id        int `form:"id,text,id"`
	Time      int `form:"time,text,时间"`
	Threshold int `form:"threshold,text,使用率"`
	Counter   int `form:"counter,text,次数"`
}

func (c *AlarmSettingForms) Valid(v *validation.Validation) {
	v.Range(c.Time, 1, 60, "time.time").Message("离线时间必须在%d到d%分钟之间", 1, 60)
	v.Range(c.Counter, 1, 10, "counter.counter").Message("告警次数必须在%d到d%次之间", 1, 10)
	v.Range(c.Threshold, 1, 100, "threshold.threshold").Message("使用率必须在%d到d%%之间", 1, 100)
}
