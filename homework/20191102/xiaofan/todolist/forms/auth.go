package forms

import (
	"github.com/astaxie/beego/validation"
	"strings"
	"todolist/models"
)

type AuthForm struct {
	UserName string `form:"username,text,用户名"`
	Password string `form:"password,password,密码"`

	User *models.User
}

// 如果你的 struct 实现了接口 validation.ValidFormer
// 当 StructTag 中的测试都成功时，将会执行 Valid 函数进行自定义验证
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
			valid.SetError("auth", "密码错误")
		}
	}
}
