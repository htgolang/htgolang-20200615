package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/dcosapp/gocmdb/server/utils"
	"time"
)

// 用户的结构体
type User struct {
	Id          int        `orm:"column(id);"json:"id"`
	Name        string     `orm:"column(name);size(32);"json:"name"`
	Password    string     `orm:"column(password);size(1024);"json:"-"`
	Gender      int        `orm:"column(gender);default(0)"json:"gender"`
	Birthday    *time.Time `orm:"column(birthday);null;default(null);"json:"birthday"`
	Tel         string     `orm:"column(tel);size(32);"json:"tel"`
	Email       string     `orm:"column(email);size(64);"json:"email"`
	Addr        string     `orm:"column(addr);size(512);"json:"addr"`
	Remark      string     `orm:"column(remark);size(1024);"json:"remark"`
	IsSuperman  bool       `orm:"column(is_superman);default(false)"json:"is_superman"`
	Status      int        `orm:"column(status);"json:"status"`
	CreatedTime *time.Time `orm:"column(created_time);auto_now_add"json:"created_time"`
	UpdatedTime *time.Time `orm:"column(updated_time);auto_now"json:"updated_time"`
	DeletedTime *time.Time `orm:"column(deleted_time);null;default(null);"json:"-"`

	Token          *Token           `orm:"reverse(one);"json:"-"`
	CloudPlatforms []*CloudPlatform `orm:"reverse(many);"json:"cloud_platforms"`
}

// 设置User密码
func (u *User) SetPassword(password string) {
	u.Password = utils.Md5Salt(password, "")
}

// 检验输入的密码和数据库中的密码是否一致
func (u *User) ValidatePassword(password string) bool {
	salt, _ := utils.SplitMd5Salt(u.Password)
	return utils.Md5Salt(password, salt) == u.Password
}

// 判断用户是否被锁定
func (u *User) IsLock() bool {
	return u.Status == StatusLock
}

// 用户管理的结构体
type UserManager struct {
}

// 生成一个新的UserManager
func NewUserManager() *UserManager {
	return &UserManager{}
}

// 通过id来查询用户
func (m *UserManager) GetById(id int) *User {
	user := &User{Id: id, DeletedTime: nil}
	// 只查询没有被逻辑删除的数据
	err := orm.NewOrm().QueryTable(user).Filter("Id__exact", id).Filter("DeletedTime__isnull", true).One(user)
	if err == nil {
		_, _ = orm.NewOrm().LoadRelated(user, "Token")
		return user
	}
	return nil
}

// 通过Name来查询用户
func (m *UserManager) GetByName(name string) *User {
	user := &User{Name: name, DeletedTime: nil}
	err := orm.NewOrm().QueryTable(user).Filter("Name__exact", name).Filter("DeletedTime__isnull", true).One(user)
	if err == nil {
		return user
	}
	return nil
}

// 查询用户
func (m *UserManager) Query(q string, start int64, length int, user *User) ([]*User, int64, int64) {
	o := orm.NewOrm()
	queryset := o.QueryTable(&User{})
	condition := orm.NewCondition()

	if user.IsSuperman != true {
		condition = condition.And("id__exact", user.Id)
	}

	condition = condition.And("deleted_time__isnull", true)
	total, _ := queryset.SetCond(condition).Count()

	qtotal := total
	if q != "" {
		query := orm.NewCondition()
		query = query.Or("name__icontains", q)
		query = query.Or("tel__icontains", q)
		query = query.Or("addr__icontains", q)
		query = query.Or("email__icontains", q)
		query = query.Or("remark__icontains", q)
		condition = condition.AndCond(query)

		qtotal, _ = queryset.SetCond(condition).Count()
	}
	var users []*User

	_, _ = queryset.SetCond(condition).Limit(length).Offset(start).All(&users)
	return users, total, qtotal
}

// 设置用户状态(0 UnLock, 1 Lock)
func (m *UserManager) SetStatusById(pk, status int) error {
	_, err := orm.NewOrm().QueryTable(&User{}).Filter("id__exact", pk).Update(orm.Params{"status": status})
	return err
}

// 根据ID删除用户(逻辑删除)
func (m *UserManager) DeleteById(pk int) (int64, error) {
	now := time.Now()
	result, err := orm.NewOrm().QueryTable(&User{}).Filter("Id__exact", pk).Update(orm.Params{"DeletedTime": &now})
	return result, err
}

// Token结构体
type Token struct {
	Id          int        `orm:"column(id)"`
	User        *User      `orm:"column(user);rel(one);"`
	AccessKey   string     `orm:"column(access_key);size(1024);"`
	SecretKey   string     `orm:"column(secret_key);size(1024);"`
	CreatedTime *time.Time `orm:"column(created_time);auto_now_add;"`
	UpdateTime  *time.Time `orm:"column(deleted_time);auto_now;"`
}

// Token管理的结构体
type TokenManager struct {
}

// 生成一个新的TokenManager
func NewTokenManager() *TokenManager {
	return &TokenManager{}
}

// 通过accessKey和secretKey来校验token
func (m *TokenManager) GetByKey(accessKey, secretKey string) *Token {
	token := &Token{AccessKey: accessKey, SecretKey: secretKey}
	o := orm.NewOrm()
	if err := o.Read(token, "AccessKey", "SecretKey"); err == nil {
		_, err = o.LoadRelated(token, "User")
		fmt.Println(token)
		return token
	}
	return nil
}

// 为用户生成token
func (m *TokenManager) GenerateByUser(user *User) *Token {
	ormer := orm.NewOrm()
	token := &Token{User: user}
	if ormer.Read(token, "User") == orm.ErrNoRows {
		token.AccessKey = utils.RandString(32)
		token.SecretKey = utils.RandString(32)
		_, _ = ormer.Insert(token)
	} else {
		token.AccessKey = utils.RandString(32)
		token.SecretKey = utils.RandString(32)
		_, _ = ormer.Update(token)
	}
	return token
}

var DefaultUserManager = NewUserManager()
var DefaultTokenManager = NewTokenManager()

// 注册User和Token
func init() {
	orm.RegisterModel(&User{}, &Token{})
}
