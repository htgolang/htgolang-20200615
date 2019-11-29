package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/xxdu521/cmdbgo/server/utils"
	"time"
)


type Token struct {
	Id			int			`orm:"type(int);"`
	AccessKey	string		`orm:"type(varchar);size(1024);default();"`
	SecrectKey	string		`orm:"type(varchar);size(1024);default();"`
	User		*User 		`orm:"rel(one)"`
	CreatedTime *time.Time	`orm:"type(datetime);auto_now_add;"`
	UpdatedTime *time.Time	`orm:"type(datetime);auto_now;"`
	DeletedTime *time.Time	`orm:"type(datetime);null;default(null);"`
}

type TokenManager struct {}

func NewTokenManager() *TokenManager {
	return &TokenManager{}
}

func (m *TokenManager) GetByKey(accessKey,secrectKey string) *Token {
	token := &Token{ AccessKey : accessKey, SecrectKey : secrectKey}
	ormer := orm.NewOrm()
	if err := ormer.Read(token, "AccessKey","SecrectKey");err == nil {
		ormer.LoadRelated(token,"User")
		return token
	}
	return nil
}

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
}

var DefaultTokenManager = NewTokenManager()

func init(){
	orm.RegisterModel(&Token{})
}