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
	Id           int		`json:"id"`
	Name         string     `orm:"type(varchar);size(256);default();" json:"name"` // 任务名
	Progress     int        `orm:"default(0);" json:"progress"`                        //进度
	Worker       string     `orm:"type(varchar);size(32);default();" json:"worker"`  //执行者（负责人）
	CreateUser   int        `orm:"default(0);" json:"create_user"`                        // 创建人
	Desc         string     `orm:"type(varchar);size(512);default();" json:"desc"` //描述
	Status       int        `orm:"default(0);" json:"status"`                        //状态
	CreateTime   *time.Time `orm:"type(datetime);auto_now_add;" json:"create_time"`       // 创建时间，在创建时自动设置（auto_now_add）
	CompleteTime *time.Time `orm:"type(datetime);null;" json:"complete_time"`               //完成时间，允许为null

	User *User  			`orm:"-" json:"create_user_object"`  //放一个用户表的数据，在任务表的结构体里，但是不存数据库，只有在用的读取的时候，进行关联查询，并输出数据给前端
	StatusText	string 		`orm:"-" json:"status_text"`
}

func (t *Task) CreateUserName() string {
	ormer := orm.NewOrm()
	user := User{Id: t.CreateUser}
	if ormer.Read(&user) == nil {
		return user.Name
	}
	return "未知"
}

func (t *Task) Patch(){
	//拿着创建用户的ID号，去user表里查一下，这个id的名字叫什么
	ormer := orm.NewOrm()
	user := &User{Id: t.CreateUser}
	if ormer.Read(user) == nil {
		t.User = user
	}
	//拿任务的创建状态数字，去映射中找出对应关闭，并赋值给返回数据专用的字段（不存数据库）
	t.StatusText = TaskStatusTexts[t.Status]
}

func init() {
	orm.RegisterModel(&Task{})
}