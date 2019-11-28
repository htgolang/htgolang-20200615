package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/xxdu521/cmdbgo/server/utils"
)

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