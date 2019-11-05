package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

const (
	TaskStatusNew = iota
	TaskStatusDoing
	TaskStatusStop
	TaskStatusComplete
)

var TaskStatusTexts = map[int]string{
	TaskStatusNew:      "新建",
	TaskStatusDoing:    "正在进行中",
	TaskStatusStop:     "停止",
	TaskStatusComplete: "完成",
}

type Task struct {
	Id           int
	Name         string     `orm:"type(varchar);size(256);default();" json:"name"`  // 人物名
	Progress     int        `orm:"default(0);" json:"progress"`                     // 进度
	Worker       string     `orm:"type(varchar);size(32);default();" json:"worker"` // 负责人
	CreateUser   int        `orm:"default(0);" json:"create_user"`                  // 创建人
	Desc         string     `orm:"type(varchar);size(512);default();" json:"desc"`  // 描述
	Status       int        `orm:"default(0);" json:"status"`                       // 状态
	CreateTime   *time.Time `orm:"type(datetime);auto_now_add;" json:"create_time"` // 创建时间，在创建时自动设置(auto_now_add)
	CompleteTime *time.Time `orm:"type(datetime);null;" json:"complete_time"`       // 完成时间，允许为null
}

func (t *Task) StatusText() string {
	return TaskStatusTexts[t.Status]
}

func (t *Task) CreateUserName() string {
	o := orm.NewOrm()
	user := User{Id: t.CreateUser}
	if o.Read(&user) == nil {
		return user.Name
	}
	return "未知"
}

func init() {
	orm.RegisterModel(&Task{})
}
