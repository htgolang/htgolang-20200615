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
func (c *AuthForm) Valid(valid *validation.Validation) {
	c.UserName = strings.TrimSpace(c.UserName)
	c.Password = strings.TrimSpace(c.Password)

	if c.UserName == "" || c.Password == "" {
		_ = valid.SetError("auth", "用户名或密码错误")
	} else {
		user := models.User{Name: c.UserName}
		if err := models.GetUser(&user, "Name"); err != nil {
			_ = valid.SetError("auth", "用户名不存在")
		} else if user.ValidatePassword(c.Password) {
			c.User = &user
		} else {
			_ = valid.SetError("auth", "密码错误")
		}
	}
}
