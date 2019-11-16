package forms

import (
	"strings"

	"github.com/astaxie/beego/validation"
	"github.com/imsilence/gocmdb/server/models"
)

type LoginForm struct {
	Name     string `form:"name"`
	Password string `form:"password"`

	User *models.User
}

func (f *LoginForm) Valid(v *validation.Validation) {
	f.Name = strings.TrimSpace(f.Name)
	f.Password = strings.TrimSpace(f.Password)

	if f.Name == "" || f.Password == "" {
		v.SetError("error", "用户名或密码错误")
	} else {
		if user := models.DefaultUserManager.GetByName(f.Name); user == nil || !user.ValidatePassword(f.Password) {
			v.SetError("error", "用户名或密码错误")
		} else if user.IsLock() {
			v.SetError("error", "用户名被锁定")
		} else {
			f.User = user
		}
	}
}
