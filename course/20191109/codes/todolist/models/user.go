package models

import (
	"time"

	"github.com/astaxie/beego/orm"

	"github.com/imsilence/todolist/utils"
)

// 用户模型
type User struct {
	Id         int        `json:"id"`
	Name       string     `orm:"type(varchar);size(32);default();" json:"name"`   //用户名
	Password   string     `orm:"type(varchar);size(1024);default();" json:"-"`    // 密码
	Birthday   *time.Time `orm:"type(date);null;" json:"birthday"`                //出生日期，允许为null
	Sex        bool       `orm:"default(false)" json:"sex"`                       //性别，true：男，false： 女
	Tel        string     `orm:"type(varchar);size(16);default()" json:"tel"`     //电话号码
	Addr       string     `orm:"type(varchar);size(512);default()" json:"addr"`   // 住址
	Desc       string     `orm:"type(text);default()" json:"desc"`                //描述
	IsSuper    bool       `orm:"default(false)" json:"is_super"`                  //是否为超级管理员, true:是，false：非
	CreateTime *time.Time `orm:"type(datetime);auto_now_add;" json:"create_time"` // 创建时间，在创建时自动设置（auto_now_add）
}

// 验证密码函数
func (u *User) ValidatePassword(password string) bool {
	return utils.Md5(password) == u.Password
}

// 更新密码函数
func (u *User) SetPassword(password string) {
	u.Password = utils.Md5(password)
}

func init() {
	orm.RegisterModel(&User{}) // 注册模型
}
