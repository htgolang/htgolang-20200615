package forms

import (
	"github.com/astaxie/beego/validation"
	"github.com/xxdu521/cmdbgo/server/models"
	"strings"
)

//用户登录表单验证
type LoginForm struct {
	Name string `form:"name"`
	Password string `form:"password"`

	User *models.User
}
func (f *LoginForm) Valid(v *validation.Validation) {
	f.Name = strings.TrimSpace(f.Name)
	f.Password = strings.TrimSpace(f.Password)

	if f.Name == "" || f.Password == "" {
		v.SetError("error","请输入用户名和密码")
	} else {
		if user := models.DefaultUserManager.GetByName(f.Name); user == nil || !user.ValidatePassword(f.Password) {
			v.SetError("error","用户名或密码错误")
		} else if user.IsLock() {
			v.SetError("error","用户异常")
		} else {
			f.User = user
		}
	}
}
