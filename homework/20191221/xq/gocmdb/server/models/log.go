package models

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	//"github.com/xlotz/gocmdb/agent/entity/log"
	"time"
)

const (

	// 表示数字，区分日志
	LOGResource = 0X0001
)


type Log struct {
	UUID string `json:"uuid"`
	Type int `json:"type"`
	Msg string `json:"msg"`
	Time *time.Time `json:"time"`

}

type LogManager struct {

}

func NewLogManager() *LogManager {
	return &LogManager{}
}

func (m *LogManager) Create(log *Log){
	switch log.Type {
	case LOGResource:
		resource := &Resource{}
		if err := json.Unmarshal([]byte(log.Msg), resource);err == nil {
			DefaultResourceManager.Create(log, resource)

		}
	}
}



type Resource struct {
	Id int `orm:"column(id)" json:"id"`
	UUID string `orm:"column(uuid);size(64)" json:"uuid"`
	CPUPrecent float64 `orm:"column(cpu_precent)" json:"cpu_precent"`
	MEMPrecent float64 `orm:"column(mem_precent)" json:"mem_precent"`
	DiskPrecent string `orm:"column(disk_precent); size(4096)" json:"disk_precent"`
	Load string `orm:"column(load);size(1024)" json:"load"`
	Time *time.Time `orm:"column(time);null" json:"time"`
	CreatedTime *time.Time `orm:"column(created_time);auto_now_add" json:"created_time"`
	DeletedTime *time.Time `orm:"column(deleted_time); null" json:"deleted_time"`
}

type ResourceManager struct {

}

func NewResourceManager() *ResourceManager {
	return &ResourceManager{}
}

func (m *ResourceManager) Create(log *Log, resource *Resource) {

	resource.UUID = log.UUID
	resource.Time = log.Time
	orm.NewOrm().Insert(resource)
}

func (m *ResourceManager) Query(q string, start int64, length int) ([]*Resource, int64, int64){
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&Resource{})

	condition := orm.NewCondition()

	condition = condition.And("deleted_time__isnull", true)

	total, _ := queryset.SetCond(condition).Count()
	qtotal := total

	if q != "" {
		query := orm.NewCondition()
		query = query.Or("uuid__icontains", q)
		//query = query.Or("created_time__", q)
		condition = condition.AndCond(query)

		qtotal, _ = queryset.SetCond(condition).Count()

	}

	var result []*Resource
	queryset.SetCond(condition).Limit(length).Offset(start).All(&result)

	return result, total, qtotal


}

func (m *ResourceManager) DeleteById(pk int) error {

	orm.NewOrm().QueryTable(&Resource{}).Filter("id__exact", pk).Update(orm.Params{"deleted_time": time.Now()})
	return nil
}

func (m *ResourceManager) Trend(uuid string) []*Resource {
	// 24小时之前
	startTime := time.Now().Add(-24 * time.Hour)
	endTime := time.Now()

	fmt.Println(uuid)

	condition := orm.NewCondition()
	condition = condition.And("deleted_time__isnull", true)
	condition = condition.And("uuid__exact", uuid)
	condition = condition.And("created_time__gte", startTime)

	var items []*Resource

	orm.NewOrm().QueryTable(&Resource{}).SetCond(condition).OrderBy("created_time").All(&items)

	//fmt.Println(result)
	var itemMap map[string]*Resource = make(map[string]*Resource)
	for _, item := range items {
		// 判断时间是否
		itemMap[item.CreatedTime.Format("2006-01-02 15:04")] = item
	}
	var result []*Resource = make([]*Resource, 0, len(items))

	for startTime.Before(endTime) {
		key := startTime.Format("2006-01-02 15:04")
		if item, ok := itemMap[key]; ok {
			result = append(result, item)
		}else {
			result = append(result, &Resource{CreatedTime: &startTime}) // 取引用
		}
		startTime = startTime.Add(time.Minute)
	}

	return result
}




var DefaultLogManager = NewLogManager()
var DefaultResourceManager = NewResourceManager()


func init(){
	orm.RegisterModel(new(Resource))

}