package models

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"time"
)

const (
	LOGResource = 0x0001
)

type Log struct {
	UUID string `json:"uuid"`
	Type int `json:"type"`
	Msg string `json:"msg"`
	Time *time.Time `json:"time"`
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
}

type Resource struct {
	Id 			int 	`orm:"column(id);" json:"id"`
	UUID 		string 	`orm:"column(uuid);size(64);" json:"uuid"`
	Load 		string 	`orm:"column(load);size(1024);" json:"load"`
	CPUPrecent 	float64 `orm:"column(cpu_percent);" json:"cpu_precent"`
	RAMPrecent 	float64 `orm:"column(ram_percent);" json:"ram_precent"`
	DiskPrecent string 	`orm:"column(disk_percent);size(4096);" json:"disk_precent"`
	Time 			*time.Time 	`orm:"column(time);" json:"time"`
	CreatedTime 	*time.Time 	`orm:"column(created_time);auto_now_add;" json:"created_time"`
	DeletedTime 	*time.Time 	`orm:"column(deleted_time);null;"json:"deleted_time"`
}

type ResourceManager struct{}

func NewResourceManager() *ResourceManager{
	return &ResourceManager{}
}

func (m *ResourceManager) Create(log *Log, resource *Resource) error{
	resource.UUID = log.UUID
	resource.Time = log.Time
	//fmt.Println(log.Time)
	_,err := orm.NewOrm().Insert(resource)
	return err
}

var DefaultLogManager = NewLogManager()
var DefaultResourceManager = NewResourceManager()

func init(){
	orm.RegisterModel(&Resource{})
}
