package forms

import (
	"github.com/astaxie/beego/validation"
	"github.com/xlotz/gocmdb/server/models"
	"strings"
)

type LoginForm struct {
	Name string `form:"name"`
	Passwrod string `form:"password"`
	User *models.User
}

func (f *LoginForm) Valid(v *validation.Validation){
	f.Name = strings.TrimSpace(f.Name)
	f.Passwrod = strings.TrimSpace(f.Passwrod)

		if f.Name == "" || f.Passwrod == "" {
			v.SetError("error", " 用户名或密码为空")

		}else {

			if user := models.DefaultUserManager.GetByName(f.Name); user == nil {
				v.SetError("error", "用户名错误")
			}else if !user.ValidatePassword(f.Passwrod){
				v.SetError("error", "密码错误")

			} else if user.IsLock(){
				v.SetError("error", "用户被锁定")
			}else {
				f.User = user
			}

		}

}
