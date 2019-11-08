package forms

import (
	"todolist/models"
	"github.com/astaxie/beego/validation"
	"strings"
)

type AuthForm struct {
	UserName string `form:"username, text,用户名"`
	Password string `form:"password, password, 密码"`
	User *models.User

}

func (this *AuthForm) Valid(valid *validation.Validation) {
	this.UserName = strings.TrimSpace(this.UserName)
	this.Password = strings.TrimSpace(this.Password)

	if this.UserName == "" || this.Password == "" {
		valid.SetError("auth", "用户名或密码错误")
	} else {
		user := models.User{Name: this.UserName}
		if err := models.GetUser(&user, "Name"); err != nil {
			valid.SetError("auth", "用户名不存在")
		} else if user.ValidatePassword(this.Password) {
			this.User = &user
		} else {
			valid.SetError("auth", "密码验证错误")
		}
	}
}