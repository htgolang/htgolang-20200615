package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type Alarm struct {
	Id          int        `orm:"column(id);" json:"id"`
	UUID        string     `orm:"column(uuid);size(64);" json:"uuid"`
	Type        int        `orm:"column(type)" json:"type"`
	Content     string     `orm:"column(content);type(text);" json:"content"`
	AlarmedTime *time.Time `orm:"column(alarmed_time);" json:"alarmed_time"`
	Status      int        `orm:"column(status);" json:"status"`
	Reason      string     `orm:"column(reason);type(text);" json:"reason"`

	CreatedTime *time.Time `orm:"column(created_time);auto_now_add;" json:"created_time"`
	DeletedTime *time.Time `orm:"column(deleted_time);null;" json:"deleted_time"`
}

type AlarmManager struct{}

func NewAlarmManager() *AlarmManager {
	return &AlarmManager{}
}

func (m *AlarmManager) Create(uuid string, typ int, content string, alarmedTime time.Time) error {
	alarm := &Alarm{
		UUID:        uuid,
		Type:        typ,
		Content:     content,
		AlarmedTime: &alarmedTime,
		Status:      AlarmStatusNew,
		Reason:      "",
	}

	_, err := orm.NewOrm().Insert(alarm)
	return err
}

func (m *AlarmManager) GetCountByUuidAndType(uuid string, typ int, startTime time.Time) int64 {
	cnt, _ := orm.NewOrm().QueryTable(new(Alarm)).Filter("uuid__exact", uuid).Filter("type__exact", typ).Filter("alarmed_time__gte", startTime).Filter("deleted_time__isnull", true).Count()
	return cnt
}

func (m *AlarmManager) GetNotification(limit int) (int64, []*Alarm) {
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(new(Alarm))

	cnt, _ := queryset.Filter("status__exact", AlarmStatusNew).Filter("deleted_time__isnull", true).Count()

	var result []*Alarm
	queryset.Filter("status__exact", AlarmStatusNew).Filter("deleted_time__isnull", true).OrderBy("-alarmed_time").Limit(limit).All(&result)

	return cnt, result
}

func (m *AlarmManager) GetCountForNoComplete() int64 {
	// total, _ := orm.NewOrm().QueryTable(new(Alarm)).Filter("deleted_time__isnull", true).Exclude("status__exact", AlarmStatusComplete).Count()
	// fmt.Println(total)
	total, _ := orm.NewOrm().QueryTable(new(Alarm)).Filter("deleted_time__isnull", true).Filter("status__in", AlarmStatusNew, AlarmStatusDoing).Count()
	return total
}

func (m *AlarmManager) GetStatForNotComplete() map[string]int64 {

	var lines []orm.Params
	orm.NewOrm().Raw("select type, count(*) as cnt from alarm where deleted_time is null and status in (?, ?) group by type", []int{AlarmStatusDoing, AlarmStatusNew}).Values(&lines)

	result := map[string]int64{}
	for _, line := range lines {
		typ := line["type"].(string)
		cntString := line["cnt"].(string)
		cnt, _ := strconv.ParseInt(cntString, 10, 64)
		result[typ] = cnt
	}
	return result
}

func (m *AlarmManager) GetLastestNStat(day int) ([]string, map[string][]int64) {

	endTime := time.Now()
	startTime := endTime.Add(-24*time.Duration(day-1)*time.Hour - 1)
	var lines []orm.Params

	orm.NewOrm().Raw("select date_format(alarmed_time, '%Y-%m-%d') as day, type, status, count(*) as cnt from alarm where deleted_time is null and alarmed_time >= ? group by day, type, status", startTime).Values(&lines)

	//key type+status = day : cnt
	tempStat := map[string]map[string]int64{}

	for _, line := range lines {
		day, _ := line["day"].(string)
		status, _ := line["status"].(string)
		typ, _ := line["type"].(string)
		cntString, _ := line["cnt"].(string)
		cnt, _ := strconv.ParseInt(cntString, 10, 64)
		key := fmt.Sprintf("%s-%s", typ, status)

		if _, ok := tempStat[key]; !ok {
			tempStat[key] = map[string]int64{}
		}
		tempStat[key][day] = cnt

	}

	// type+status : [1, 2, 3, 3, 4]
	// "int-int"
	days := []string{}
	result := map[string][]int64{}

	for startTime.Before(endTime) {
		day := startTime.Format("2006-01-02")
		days = append(days, day)

		for _, typ := range []int{AlarmTypeOffline, AlarmTypeCPU, AlarmTypeRam} {
			for _, status := range []int{AlarmStatusNew, AlarmStatusDoing, AlarmStatusComplete} {
				key := fmt.Sprintf("%d-%d", typ, status)
				value := int64(0)
				if stat, ok := tempStat[key]; ok {
					value = stat[day]
				} else {
				}
				result[key] = append(result[key], value)
			}
		}
		startTime = startTime.Add(24 * time.Hour)
	}

	return days, result
}

func (m *AlarmManager) Query(q string, start int64, length int) ([]*Alarm, int64, int64) {
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&Alarm{})
	condition := orm.NewCondition()

	condition = condition.And("deleted_time__isnull", true)

	total, _ := queryset.SetCond(condition).Count()

	qtotal := total
	if q != "" {
		query := orm.NewCondition()

		condition = condition.AndCond(query)

		qtotal, _ = queryset.SetCond(condition).Count()
	}
	var result []*Alarm
	queryset.SetCond(condition).OrderBy("-created_time").Limit(length).Offset(start).All(&result)
	return result, total, qtotal
}

var DefaultAlarmManager = NewAlarmManager()

func init() {
	orm.RegisterModel(new(Alarm))
}
