package models

import (
	"time"

	"github.com/astaxie/beego/orm"

	"github.com/imsilence/gocmdb/server/utils"
)

type User struct {
	Id          int        `orm:"column(id);"`
	Name        string     `orm:"column(name);size(32);"`
	Password    string     `orm:"column(password);size(1024);"`
	Gender      int        `orm:"column(gender);default(0);"`
	Birthday    *time.Time `orm:"column(birthday);null;default(null);"`
	Tel         string     `orm:"column(tel);size(1024);"`
	Email       string     `orm:"column(email);size(1024);"`
	Addr        string     `orm:"column(addr);size(1024);"`
	Remark      string     `orm:"column(remark);size(1024);"`
	IsSuperman  bool       `orm:"column(is_superman);default(false);"`
	Status      int        `orm:"column(status);"`
	CreatedTime *time.Time `orm:"column(created_time);auto_now_add;"`
	UpdatedTime *time.Time `orm:"column(update_time);auto_now;"`
	DeletedTime *time.Time `orm:"column(deleted_time);null;default(null);"`

	Token *Token `orm:"reverse(one);"`
}

func (u *User) SetPassword(password string) {
	u.Password = utils.Md5Salt(password, "")
}

func (u *User) ValidatePassword(password string) bool {
	salt, _ := utils.SplitMd5Salt(u.Password)
	return utils.Md5Salt(password, salt) == u.Password
}

func (u *User) IsLock() bool {
	return u.Status == StatusLock
}

type UserManager struct{}

func NewUserManager() *UserManager {
	return &UserManager{}
}

func (m *UserManager) GetById(id int) *User {
	user := &User{}
	err := orm.NewOrm().QueryTable(user).Filter("Id__exact", id).Filter("DeletedTime__isnull", true).One(user)
	if err == nil {
		return user
	}
	return nil
}

func (m *UserManager) GetByName(name string) *User {
	user := &User{}
	err := orm.NewOrm().QueryTable(user).Filter("Name__exact", name).Filter("DeletedTime__isnull", true).One(user)
	if err == nil {
		return user
	}

	return nil
}

type Token struct {
	Id          int        `orm:"column(id);"`
	User        *User      `orm:"column(user);rel(one);"`
	AccessKey   string     `orm:"column(access_key);size(1024);"`
	SecrectKey  string     `orm:"column(secrect_key);size(1024);"`
	CreatedTime *time.Time `orm:"column(created_time);auto_now_add;"`
	UpdateTime  *time.Time `orm:"column(updated_time);auto_now;"`
}

type TokenManager struct {
}

func NewTokenManager() *TokenManager {
	return &TokenManager{}
}

func (m *TokenManager) GetByKey(accessKey, secrectKey string) *Token {
	token := &Token{AccessKey: accessKey, SecrectKey: secrectKey}
	ormer := orm.NewOrm()
	if err := ormer.Read(token, "AccessKey", "SecrectKey"); err == nil {
		ormer.LoadRelated(token, "User")
		return token
	}
	return nil
}

var DefaultUserManager = NewUserManager()
var DefaultTokenManager = NewTokenManager()

func init() {
	orm.RegisterModel(new(User), new(Token))
}
