package forms

import (
	"github.com/astaxie/beego/validation"
	"github.com/dcosapp/gocmdb/server/models"
	"strings"
)

type LoginForm struct {
	Name     string `form:"name"`
	Password string `form:"password"`

	User *models.User
}

// 登录表单验证
func (f *LoginForm) Valid(v *validation.Validation) {
	f.Name = strings.TrimSpace(f.Name)
	f.Password = strings.TrimSpace(f.Password)

	if f.Name == "" || f.Password == "" {
		_ = v.SetError("error", "用户名或密码错误")
	} else {
		if user := models.DefaultUserManager.GetByName(f.Name); user == nil || !user.ValidatePassword(f.Password) {
			_ = v.SetError("error", "用户名或密码错误")
		} else if user.IsLock() {
			_ = v.SetError("error", "用户已被锁定")
		} else {
			f.User = user
		}
	}
}
