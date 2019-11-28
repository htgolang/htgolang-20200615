package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/xlotz/gocmdb/server/utils"
	"time"
)

// 用户模型
type User struct {
	Id         int		`orm:"column(id)" json:"id"`
	Name       string     `orm:"column(name);size(32);" json:"name"`   //用户名
	Password   string     `orm:"column(password);size(1024);" json:"-"`// 密码
	Birthday   *time.Time `orm:"column(birthday);null;default(null);" json:"birthday"`                    //出生日期，允许为null
	Gender        int       `orm:"column(gender);default(0)" json:"gender"`                      //性别，0：男，1： 女
	Tel        string     `orm:"column(tel);size(16);" json:"tel"`    //电话号码
	Email	   string	`orm:"column(email);size(64);" json:"email"`
	Addr       string     `orm:"column(addr);size(512);" json:"addr"`   // 住址
	Desc       string     `orm:"column(desc);" json:"desc"`                //描述
	IsSuperman    bool       `orm:"column(is_superman);default(false)" json:"is_superman"`                      //是否为超级管理员, true:是，false：非
	Remark     string     `orm:"column(remark);size(1024);" json:"remark"`
	Status     int        `orm:"column(status);default(0);" json:"status"` 		// 0表示正常，1表示锁定
	CreatedTime *time.Time `orm:"column(created_time);auto_now_add;" json:"created_time"`        // 创建时间，在创建时自动设置（auto_now_add）
	UpdatedTime *time.Time `orm:"column(updated_time);auto_now;" json:"updated_time"` // 更新时间, 在更新时更新
	DeletedTime *time.Time `orm:"column(deleted_time);null;default(null)" json:"-"`     //允许为null, null 未删除， !null 已删除

	// 和用户一对一
	Token *Token `orm:"reverse(one);" json:"token"`
}

func (u *User) SetPassword(password string) {
	u.Password = utils.Md5Salt(password, "")
}

func (u *User) ValidatePassword(password string) bool {

	salt, _ := utils.SplitMd5Salt(u.Password)

	return utils.Md5Salt(password, salt) == u.Password
}

func (u *User) IsLock() bool{
	return u.Status == StatusLock
}



type UserManager struct {

}

func NewUserManager() *UserManager{
	return &UserManager{}
}

func (m *UserManager) GetById(id int) *User{

	//user := &User{Id: id, DeletedTime: nil}
	user := &User{}
	ormer := orm.NewOrm()

	err := ormer.QueryTable(user).Filter("Id__exact", id).Filter("DeletedTime__isnull", true).One(user)

	if err == nil {
		ormer.LoadRelated(user, "Token")
		return user
	}

	return nil
}

func (m *UserManager) GetByName(name string) *User{
	user := &User{}
	ormer := orm.NewOrm()

	err := ormer.QueryTable(user).Filter("Name__exact", name).Filter("DeletedTime__isnull", true).One(user)

	if err == nil {
		return user
	}

	return nil


}

func (m *UserManager) Query(q string, start int64, length int) ([]*User, int64, int64){

	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&User{})

	condition := orm.NewCondition()

	condition.And("deleted_time__isnull", true)


	total, _ := queryset.SetCond(condition).Count()

	qtotal := total

	if q != "" {
		query := orm.NewCondition()
		query = query.Or("name__icontains", q)
		query = query.Or("tel__icontains", q)
		query = query.Or("addr__icontains", q)
		query = query.Or("email__icontains", q)
		condition = condition.AndCond(query)
		qtotal, _ = queryset.SetCond(condition).Count()
	}

	var users []*User
	queryset.SetCond(condition).Limit(length).Offset(start).All(&users)


	return users, total, qtotal

}

func (m *UserManager) DeleteById(pk int) error{
	orm.NewOrm().QueryTable(&User{}).Filter("id__exact", pk).Update(orm.Params{"deleted_time": time.Now()})
	return nil

}

func (m *UserManager) SetStatusById(pk int, status int) error {

	fmt.Println(pk, status)

	if _, err :=orm.NewOrm().QueryTable(&User{}).Filter("id__exact", pk).Update(orm.Params{"status": status});err ==nil{
		return nil
	}else {
		return err
	}

}

var DefaultUserManager = NewUserManager()




type Token struct {
	Id int `orm:"column(id);"`
	User *User `orm:"column(user);rel(one);"`
	AccessKey string `orm:"column(access_key);size(1024);"`
	SecrectKey string `orm:"column(secrect_key);size(1024);"`
	CreatedTime *time.Time `orm:"column(created_time);auto_now_add;"`
	UpdatedTime *time.Time `orm:"column(updated_time);auto_now;"`

}

type TokenManager struct {

}

func NewTokenManager() *TokenManager{
	return &TokenManager{}
}

func (m *TokenManager) GetByKey(accesskey, secrectkey string) *Token{
	fmt.Println(accesskey, secrectkey)

	token := &Token{AccessKey:accesskey, SecrectKey: secrectkey}
	ormer := orm.NewOrm()

	if err := ormer.Read(token, "AccessKey", "SecrectKey"); err ==nil {

		ormer.LoadRelated(token, "User")
		return token
	}

	return nil

}

func (m *TokenManager) GenerateByUser(user *User) *Token{

	ormer := orm.NewOrm()
	token := &Token{User: user}
	if ormer.Read(token, "User") == orm.ErrNoRows {
		token.AccessKey = utils.RandString(32)
		token.SecrectKey = utils.RandString(32)
		fmt.Println(token.AccessKey)
		ormer.Insert(token)
	}else {
		token.AccessKey = utils.RandString(32)
		token.SecrectKey = utils.RandString(32)
		fmt.Println(token.AccessKey)
		ormer.Update(token)
	}
	return token
}

var DefaultTokenManager = NewTokenManager()


func init() {
	orm.RegisterModel(new(User), new(Token)) // 注册模型

}