package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// 通过iota定义枚举值（任务状态）
const (
	TastStatusNew = iota
	TastStatusDoing
	TastStatusStop
	TastStatusComplete
)

// 定义任务状态映射
var TaskStatusTexts = map[int]string{
	TastStatusNew:      "新建",
	TastStatusDoing:    "正在进行中",
	TastStatusStop:     "停止",
	TastStatusComplete: "完成",
}

// 定义任务模型
type Task struct {
	Id           int
	Name         string     `orm:"type(varchar);size(256);default();"` // 任务名
	Progress     int        `orm:"default(0);"`                        //进度
	Worker       string     `orm:"type(varchar);size(32);default();"`  //执行者（负责人）
	CreateUser   int        `orm:"default(0);"`                        // 创建人
	Desc         string     `orm:"type(varchar);size(512);default();"` //描述
	Status       int        `orm:"default(0);"`                        //状态
	CreateTime   *time.Time `orm:"type(datetime);auto_now_add;"`       // 创建时间，在创建时自动设置（auto_now_add）
	CompleteTime *time.Time `orm:"type(datetime);null;"`               //完成时间，允许为null
}

// 获取任务状态中文
func (t *Task) StatusText() string {
	return TaskStatusTexts[t.Status]
}

// 获取创建者用户名
func (t *Task) CreateUserName() string {
	ormer := orm.NewOrm()
	user := User{Id: t.CreateUser}
	if ormer.Read(&user) == nil {
		return user.Name
	}
	return "未知"
}

func init() {
	orm.RegisterModel(&Task{}) //注册任务模型
}
