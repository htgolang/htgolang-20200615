package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"

	"time"
)

type Alarm struct {
	Id int `orm:"column(id)" json:"id"`
	UUID string `orm:"column(uuid);size(64)" json:"uuid"`
	Type int `orm:"column(type)" json:"type"`  // 1 离线，2 CPU, 3 MEM
	Description string `orm:"column(description); null" json:"description"`
	CreatedTime *time.Time `orm:"column(created_time); auto_now_add" json:"created_time"`
	DealedTime *time.Time `orm:"column(dealed_time); null" json:"dealed_time"` // 处理时间
	Reason string `orm:"column(reason); size(4096); null" json:"reason"`
	Notices int `orm:"column(notices)" json:"notices"` // 0 sms, 1 email
	Status int `orm:"column(status)" json:"status"` // 0 未处理，1 正在处理 2 已处理
}

type AlarmManager struct {

}

func NewAlarmManager() *AlarmManager {
	return &AlarmManager{}
}


func (m *AlarmManager) Query(q string, start int64, length int) ([]*Alarm, int64, int64){
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&Alarm{})

	condition := orm.NewCondition()

	total, _ := queryset.SetCond(condition).Count()
	qtotal := total

	if q != "" {
		query := orm.NewCondition()
		query = query.Or("created_time__icontains", q)
		query = query.Or("uuid__iconitains", q)
		query = query.Or("reason__icontains",q)


		condition = condition.AndCond(query)

		qtotal, _ = queryset.SetCond(condition).Count()

	}

	var result []*Alarm
	queryset.SetCond(condition).Limit(length).Offset(start).All(&result)

	return result, total, qtotal


}

func (m *AlarmManager) Create(typ int, uuid []orm.Params, reason string) {

	ormer := orm.NewOrm()
	now := time.Now()

	for _, id := range uuid {
		fmt.Println(id)
		ormer.Insert(&Alarm{
			UUID: fmt.Sprintf("%s", id["uuid"]),
			Type:typ,
			Reason:reason,
			Status:0,
			CreatedTime:&now,
		})
	}

}

func (m *AlarmManager) DeleteById(pk int) error {

	orm.NewOrm().QueryTable(&Alarm{}).Filter("id__exact", pk).Delete()
	return nil
}


var DefaultAlarmManager  = NewAlarmManager()


func init(){
	orm.RegisterModel(new(Alarm))
}