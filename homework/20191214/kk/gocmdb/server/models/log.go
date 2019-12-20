package models

import (
	"encoding/json"
	"time"

	"github.com/astaxie/beego/orm"
)

const (
	LOGResource = 0X0001
)

type Log struct {
	UUID string     `json:"uuid"`
	Type int        `json:"type"`
	Msg  string     `json:"msg"`
	Time *time.Time `json:"time"`
}

type LogManager struct{}

func NewLogManager() *LogManager {
	return &LogManager{}
}

func (m *LogManager) Create(log *Log) {
	switch log.Type {
	case LOGResource:
		resource := &Resource{}
		if err := json.Unmarshal([]byte(log.Msg), resource); err == nil {
			DefaultResourceManager.Create(log, resource)
		}
	}
}

type Resource struct {
	Id          int        `orm:"column(id);" json:"id"`
	UUID        string     `orm:"column(uuid);size(64);" json:"uuid"`
	Load        string     `orm:"column(load);size(1024);"  json:"load"`
	CPUPrecent  float64    `orm:"column(cpu_percent);" json:"cpu_percent"`
	RAMPrecent  float64    `orm:"column(ram_percent);" json:"ram_percent"`
	DiskPrecent string     `orm:"column(disk_percent);size(4096);" json:"disk_percent"`
	Time        *time.Time `orm:"column(time);" json:"time"`
	CreatedTime *time.Time `orm:"column(created_time);auto_now_add;" json:"created_time"`
	DeletedTime *time.Time `orm:"column(deleted_time);null;" json:"deleted_time"`
}

type ResourceManager struct{}

func NewResourceManager() *ResourceManager {
	return &ResourceManager{}
}

func (m *ResourceManager) Create(log *Log, resource *Resource) {
	resource.UUID = log.UUID
	resource.Time = log.Time
	orm.NewOrm().Insert(resource)
}


func (m *ResourceManager) Query(q string, start int64, length int) ([]*Resource, int64, int64) {
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&Resource{})
	condition := orm.NewCondition()

	condition = condition.And("deleted_time__isnull", true)

	total, _ := queryset.SetCond(condition).Count()

	qtotal := total
	if q != "" {
		query := orm.NewCondition()

		condition = condition.AndCond(query)

		qtotal, _ = queryset.SetCond(condition).Count()
	}
	var result []*Resource
	queryset.SetCond(condition).OrderBy("-created_time").Limit(length).Offset(start).All(&result)
	return result, total, qtotal
}


var DefaultLogManager = NewLogManager()
var DefaultResourceManager = NewResourceManager()

func init() {
	orm.RegisterModel(new(Resource))
}
