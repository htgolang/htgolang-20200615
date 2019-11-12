package forms

import (
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"

	"github.com/imsilence/todolist/models"
)

// 用户认证表单
type AuthForm struct {
	UserName string `form:"username,text,用户名"`
	Password string `form:"password,password,密码"`
	Next     string `form:"next,hidden,跳转URL"`

	User *models.User
}

// 用户任务表单 验证接口（由validation.Valid调用）
func (f *AuthForm) Valid(v *validation.Validation) {
	// 去除用户输入前后空白字符
	f.UserName = strings.TrimSpace(f.UserName)
	f.Password = strings.TrimSpace(f.Password)

	// 验证输入不能为空
	if f.UserName == "" || f.Password == "" {
		v.SetError("auth", "用户名或密码错误") //设置错误信息
	} else {
		// 通过用户名查找用户，并验证密码是否正确
		user := models.User{Name: f.UserName}
		ormer := orm.NewOrm()
		err := ormer.Read(&user, "Name")
		if err != nil || !user.ValidatePassword(f.Password) {
			v.SetError("auth", "用户名或密码错误")
		} else {
			f.User = &user
		}
	}
}
