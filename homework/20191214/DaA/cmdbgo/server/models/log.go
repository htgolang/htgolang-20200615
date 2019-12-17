package models

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"time"
)

const (
	LOGResource = 0x0001
)

//log通道，作为一个通用通道，把各种信息放入msg里，在服务端，分别存储到各自的模型中。
type Log struct {
	UUID 	string 		`json:"uuid"`
	Type 	int 		`json:"type"`
	Msg 	string 		`json:"msg"`
	Time 	*time.Time 	`json:"time"`
}
type LogManager struct{}
func NewLogManager() *LogManager{
	return &LogManager{}
}
func (m *LogManager) Create(log *Log) error {
	switch log.Type{
	case LOGResource:
		resource := &Resource{}
		if err := json.Unmarshal([]byte(log.Msg), resource); err == nil {
			err = DefaultResourceManager.Create(log, resource)
			return err
		} else {
			return err
		}
	}
	return nil
} //Log通道接收信息，如果是LOGResource的信息，就用Resource模型接收存储并存储

//resource插件，作为log通道的一个分支，使用log通道的msg字段，上报消息。
type Resource struct {
	Id 				int 		`orm:"column(id);" json:"id"`
	UUID 			string 		`orm:"column(uuid);size(64);" json:"uuid"`
	Load 			string 		`orm:"column(load);size(1024);" json:"load"`
	CPUPrecent 		float64 	`orm:"column(cpu_percent);" json:"cpu_precent"`
	RAMPrecent 		float64 	`orm:"column(ram_percent);" json:"ram_precent"`
	DiskPrecent 	string 		`orm:"column(disk_percent);size(4096);" json:"disk_precent"`
	Time 			*time.Time 	`orm:"column(time);" json:"time"`
	CreatedTime 	*time.Time 	`orm:"column(created_time);auto_now_add;" json:"created_time"`
	DeletedTime 	*time.Time 	`orm:"column(deleted_time);null;"json:"deleted_time"`
}
type ResourceManager struct{}
func NewResourceManager() *ResourceManager {
	return &ResourceManager{}
}
func (m *ResourceManager) Create(log *Log, resource *Resource) error {
	resource.UUID = log.UUID
	resource.Time = log.Time
	//fmt.Println(log.Time)
	_,err := orm.NewOrm().Insert(resource)
	return err
} //存储到数据库
func (m *ResourceManager) Query(q string, start int64, length int) ([]*Resource, int64, int64) {
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&Resource{})

	condition := orm.NewCondition()
	condition = condition.And("deleted_time__isnull",true)

	total,_ := queryset.SetCond(condition).Count()

	qtotal := total
	if q != "" {
		query := orm.NewCondition()
		query = query.Or("uuid__icontains",q)
		condition = condition.AndCond(query)

		qtotal, _ = queryset.SetCond(condition).Count()
	}
	var result []*Resource

	queryset.SetCond(condition).Limit(length).Offset(start).All(&result)
	return result,total,qtotal
}


var DefaultLogManager = NewLogManager()
var DefaultResourceManager = NewResourceManager()


func init(){
	orm.RegisterModel(&Resource{})
}
