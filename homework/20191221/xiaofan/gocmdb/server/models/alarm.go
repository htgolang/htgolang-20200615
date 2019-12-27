package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type Alarm struct {
	Id           int        `orm:"column(id);"json:"id"`
	Name         string     `orm:"column(name);size(64);null;"json:"name"`
	UUID         string     `orm:"column(uuid);size(64);null;"json:"uuid"`
	Type         int        `orm:"column(type);"json:"type"`
	Status       int        `orm:"column(status);"json:"status"`
	Reason       string     `orm:"column(reason);size(1024);null;"json:"reason"`
	Desc         string     `orm:"column(desc);size(1024);null;"json:"desc"`
	CreatedTime  *time.Time `orm:"column(created_time);auto_now_add;"json:"created_time"`
	DealedTime   *time.Time `orm:"column(dealed_time);null;"json:"dealed_time"`
	CompleteTime *time.Time `orm:"column(complete_time);null;"json:"complete_time"`

	// created_time 产生时间
	// uuid 终端
	// type 类型（1:离线 2：cpu 3：内存）
	// description 告警描述
	// status 状态(0:未处理, 1:正在处理， 2:已处理)
	// dealed_time 处理时间
	// reason 告警原因说明
	// notices 通知方式(sms email)
	// user 通知给谁
}

type AlarmManager struct {
}

func NewAlarmManager() *AlarmManager {
	return &AlarmManager{}
}

func (m *AlarmManager) CreateAlarm(typ int, uuid []orm.Params, reason string) {
	ormer := orm.NewOrm()
	for _, value := range uuid {
		id := fmt.Sprintf("%s", value["uuid"])
		_, _ = ormer.Insert(&Alarm{
			Name:   DefaultAgentManager.GetByUUID(id).Name,
			UUID:   id,
			Type:   typ,
			Status: 0,
			Reason: reason,
		})
	}
}

func (m *AlarmManager) Query(q string, start int64, length int) ([]*Alarm, int64, int64) {
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&Alarm{})
	condition := orm.NewCondition()
	//condition = condition.And("deleted_time__isnull", true)
	total, _ := queryset.SetCond(condition).Count()

	qtotal := total
	if q != "" {
		query := orm.NewCondition()
		query = query.Or("name__icontains", q)
		query = query.Or("uuid__icontains", q)
		query = query.Or("reason__icontains", q)
		query = query.Or("desc__icontains", q)

		condition = condition.AndCond(query)

		qtotal, _ = queryset.SetCond(condition).Count()
	}
	var result []*Alarm
	_, _ = queryset.SetCond(condition).RelatedSel().Limit(length).Offset(start).OrderBy("-id").All(&result)
	return result, total, qtotal
}

func (m *AlarmManager) SetStatusById(id, status int) error {
	if status == 1 {
		_, err := orm.NewOrm().QueryTable(&Alarm{}).Filter("id__exact", id).Update(orm.Params{"status": status, "dealed_time": time.Now()})
		return err
	}
	_, err := orm.NewOrm().QueryTable(&Alarm{}).Filter("id__exact", id).Update(orm.Params{"status": status, "complete_time": time.Now()})
	return err
}

type AlarmSetting struct {
	Id        int    `orm:"column(id);"json:"id"`
	Name      string `orm:"column(name);size(64);null;"json:"name"`
	Type      int    `orm:"column(type);"json:"type"`
	Time      int    `orm:"column(time);"json:"time"`
	Threshold int    `orm:"column(threshold);"json:"threshold"`
	Counter   int    `orm:"column(counter);"json:"counter"`
}

type AlarmSettingManager struct {
}

func NewAlarmSettingManager() *AlarmSettingManager {
	return &AlarmSettingManager{}
}

func (m *AlarmSettingManager) Query(q string, start int64, length int) ([]*AlarmSetting, int64, int64) {
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&AlarmSetting{})
	condition := orm.NewCondition()
	total, _ := queryset.SetCond(condition).Count()

	qtotal := total
	if q != "" {
		query := orm.NewCondition()
		query = query.Or("name__icontains", q)
		query = query.Or("type__icontains", q)
		condition = condition.AndCond(query)

		qtotal, _ = queryset.SetCond(condition).Count()
	}
	var result []*AlarmSetting
	_, _ = queryset.SetCond(condition).RelatedSel().Limit(length).Offset(start).All(&result)
	return result, total, qtotal
}

func (m *AlarmSettingManager) GetById(id int) *AlarmSetting {
	var set = &AlarmSetting{Id: id}
	_ = orm.NewOrm().QueryTable(set).Filter("Id__exact", id).One(set)
	return set
}

func (m *AlarmSettingManager) Modify(id, time, threshold, counter int) (*AlarmSetting, error) {
	set := &AlarmSetting{Id: id}
	ormer := orm.NewOrm()
	err := ormer.Read(set)
	if err == nil {
		set.Time = time
		set.Threshold = threshold
		set.Counter = counter
		if _, err := ormer.Update(set); err == nil {
			return set, nil
		}
	}
	return &AlarmSetting{}, err
}

var DefaultAlarmManager = NewAlarmManager()
var DefaultAlarmSettingManager = NewAlarmSettingManager()

func init() {
	orm.RegisterModel(&Alarm{}, &AlarmSetting{})
}
