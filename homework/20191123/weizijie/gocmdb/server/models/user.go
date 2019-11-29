package models

import (
	"time"

	"github.com/astaxie/beego/orm"

	"github.com/imsilence/gocmdb/server/utils"
)

type User struct {
	Id          int        `orm:"column(id);" json:"id"`
	Name        string     `orm:"column(name);size(32);" json:"name"`
	Password    string     `orm:"column(password);size(1024);" json:"-"`
	Gender      int        `orm:"column(gender);default(0);" json:"gender"`
	Birthday    *time.Time `orm:"column(birthday);null;default(null);" json:"birthday"`
	Tel         string     `orm:"column(tel);size(1024);" json:"tel"`
	Email       string     `orm:"column(email);size(1024);" json:"email"`
	Addr        string     `orm:"column(addr);size(1024);" json:"addr"`
	Remark      string     `orm:"column(remark);size(1024);" json:"remark"`
	IsSuperman  bool       `orm:"column(is_superman);default(false);" json:"is_superman"`
	Status      int        `orm:"column(status);" json:"status"`
	CreatedTime *time.Time `orm:"column(created_time);auto_now_add;" json:"created_time"`
	UpdatedTime *time.Time `orm:"column(update_time);auto_now;" json:"updated_time"`
	DeletedTime *time.Time `orm:"column(deleted_time);null;default(null);" json:"-"`

	Token *Token `orm:"reverse(one);" json:"token"`
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
	ormer := orm.NewOrm()
	err := ormer.QueryTable(user).Filter("Id__exact", id).Filter("DeletedTime__isnull", true).One(user)
	if err == nil {
		ormer.LoadRelated(user, "Token")
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

func (m *UserManager) Query(q string, start int64, length int) ([]*User, int64, int64) {
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&User{})
	condition := orm.NewCondition()

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
	queryset.SetCond(condition).Limit(length).Offset(start).All(&users)
	return users, total, qtotal
}

func (m *UserManager) DeleteById(pk int) error {
	orm.NewOrm().QueryTable(&User{}).Filter("id__exact", pk).Update(orm.Params{"deleted_time": time.Now()})
	return nil
}

func (m *UserManager) SetStatusById(pk int, status int) error {
	orm.NewOrm().QueryTable(&User{}).Filter("id__exact", pk).Update(orm.Params{"status": status})
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

func (m *TokenManager) GenerateByUser(user *User) *Token {
	ormer := orm.NewOrm()
	token := &Token{User: user}
	if ormer.Read(token, "User") == orm.ErrNoRows {
		token.AccessKey = utils.RandString(32)
		token.SecrectKey = utils.RandString(32)
		ormer.Insert(token)
	} else {
		token.AccessKey = utils.RandString(32)
		token.SecrectKey = utils.RandString(32)
		ormer.Update(token)
	}
	return token
}

var DefaultUserManager = NewUserManager()
var DefaultTokenManager = NewTokenManager()

func init() {
	orm.RegisterModel(new(User), new(Token))
}
