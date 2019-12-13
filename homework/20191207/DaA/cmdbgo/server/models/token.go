package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/xxdu521/cmdbgo/server/utils"
	"time"
)

//token模型
type Token struct {
	Id			int			`orm:"type(int);"`
	AccessKey	string		`orm:"type(varchar);size(1024);default();"`
	SecrectKey	string		`orm:"type(varchar);size(1024);default();"`
	User		*User 		`orm:"rel(one)"`
	CreatedTime *time.Time	`orm:"type(datetime);auto_now_add;"`
	UpdatedTime *time.Time	`orm:"type(datetime);auto_now;"`
	DeletedTime *time.Time	`orm:"type(datetime);null;default(null);"`
}
//token模型管理器
type TokenManager struct {}
func NewTokenManager() *TokenManager {
	return &TokenManager{}
}			//初始化token模型管理器对象
func (m *TokenManager) GetByKey(accessKey,secrectKey string) *Token {
	token := &Token{ AccessKey : accessKey, SecrectKey : secrectKey}
	ormer := orm.NewOrm()
	if err := ormer.Read(token, "AccessKey","SecrectKey");err == nil {
		ormer.LoadRelated(token,"User")
		return token
	}
	return nil
}	//通过key获取token模型数据
func (m *TokenManager) GenerateByUser(user *User) *Token {
	ormer := orm.NewOrm()
	token := &Token{User: user}
	if ormer.Read(token,"User") == orm.ErrNoRows {
		token.SecrectKey = utils.RandString(32)
		token.AccessKey = utils.RandString(32)
		ormer.Insert(token)
	} else {
		token.SecrectKey = utils.RandString(32)
		token.AccessKey = utils.RandString(32)
		ormer.Update(token)
	}
	return token
}				//通过用户创建token模型数据


var DefaultTokenManager = NewTokenManager()

func init(){
	orm.RegisterModel(&Token{})
}