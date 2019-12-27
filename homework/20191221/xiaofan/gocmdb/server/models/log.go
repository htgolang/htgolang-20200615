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
	UUID string     `json:"uuid"`
	Type int        `json:"type"`
	Msg  string     `json:"msg"`
	Time *time.Time `json:"time"`
}

type LogManager struct {
}

func NewLogManager() *LogManager {
	return &LogManager{}
}

// 将log接收的数据进行存储
func (m *LogManager) Create(log *Log) error {
	switch log.Type {
	case LOGResource:
		resource := &Resource{}

		if err := json.Unmarshal([]byte(log.Msg), resource); err != nil {
			return err
		}

		if err := DefaultResourceManager.Create(log, resource); err != nil {
			return err
		}
	}
	return nil
}

type Resource struct {
	Id          int        `orm:"column(id);"json:"id"`
	Name        string     `orm:"column(name);"json:"name"`
	UUID        string     `orm:"column(uuid);size(64);"json:"uuid"`
	Load        string     `orm:"column(load);size(1024)"json:"load"`
	CPUPercent  float64    `orm:"column(cpu_percent);"json:"cpu_percent"`
	RAMPercent  float64    `orm:"column(ram_percent);"json:"ram_percent"`
	DiskPercent string     `orm:"column(disk_percent);size(4096)"json:"disk_percent"`
	Time        *time.Time `orm:"column(time);"json:"time"`
	CreatedTime *time.Time `orm:"column(created_time);auto_now_add;"json:"created_time"`
	DeletedTime *time.Time `orm:"column(deleted_time);null;"json:"deleted_time"`
}

type ResourceManager struct {
}

func NewResourceManager() *ResourceManager {
	return &ResourceManager{}
}

// 将log的数据组合后，写入resource表中
func (m *ResourceManager) Create(log *Log, resource *Resource) error {
	resource.UUID = log.UUID
	resource.Time = log.Time
	resource.Name = DefaultAgentManager.GetByUUID(resource.UUID).Name
	_, err := orm.NewOrm().Insert(resource)
	return err
}

// 查询log
func (m *ResourceManager) Query(q string, start int64, length int, startTime, endTime string) ([]*Resource, int64, int64) {
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&Resource{})
	condition := orm.NewCondition()
	condition = condition.And("deleted_time__isnull", true)
	total, _ := queryset.SetCond(condition).Count()

	if startTime != "" && endTime != "" {
		condition = condition.And("time__gt", startTime)
		condition = condition.And("time__lt", endTime)
	}

	if q != "" {
		query := orm.NewCondition()
		query = query.Or("uuid__icontains", q)
		query = query.Or("name__icontains", q)
		condition = condition.AndCond(query)
	}

	var result []*Resource

	qtotal, _ := queryset.SetCond(condition).Count()
	_, _ = queryset.SetCond(condition).RelatedSel().Limit(length).Offset(start).All(&result)
	return result, total, qtotal
}

// 构造监控数据
func (m *ResourceManager) Trend(uuid string) []*Resource {
	// 查询最近一小时的内的监控数据
	endTime := time.Now()
	startTime := time.Now().Add(-1 * time.Hour)

	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&Resource{})
	condition := orm.NewCondition()
	condition = condition.And("deleted_time__isnull", true)
	condition = condition.And("created_time__gte", startTime)
	condition = condition.And("uuid__exact", uuid)

	var items []*Resource
	_, _ = queryset.SetCond(condition).OrderBy("created_time").All(&items)

	// 生成一个key = 时间,value = resource的字典
	var itemMapMap = make(map[string]*Resource)
	for _, item := range items {
		itemMapMap[item.CreatedTime.Format("2016-01-02 15:04")] = item
	}

	// 获取每分钟相对应的数据，如果未获取到，则使用0数据进行填充，数据的创建时间为对应的时间
	var result = make([]*Resource, 0, len(items))
	for startTime.Before(endTime) {
		key := startTime.Format("2016-01-02 15:04")
		if item, ok := itemMapMap[key]; ok {
			result = append(result, item)
		} else {
			createTime := startTime
			result = append(result, &Resource{CreatedTime: &createTime})
		}
		startTime = startTime.Add(time.Minute)
	}

	return result
}

var DefaultLogManager = NewLogManager()
var DefaultResourceManager = NewResourceManager()

func init() {
	orm.RegisterModel(&Resource{})
}
