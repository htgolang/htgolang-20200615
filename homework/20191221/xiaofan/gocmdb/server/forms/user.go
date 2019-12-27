package forms

import (
	"github.com/dcosapp/gocmdb/server/models"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

// 用户创建表单
type UserCreateForm struct {
	Name           string `form:"name,text,名称"`
	Password       string `form:"password,password,密码"`
	PasswordVerify string `form:"passwordVerify,password,再次输入密码"`
	Gender         int    `form:"gender,radio,性别"`
	Birthday       string `form:"birthday,date,出生日期"`
	Tel            string `form:"tel,text,联系方式"`
	Email          string `form:"email,text,邮箱"`
	Addr           string `form:"addr,textarea,住址"`
	Remark         string `form:"remark,textarea,备注"`
}

// 用户创建表单 验证接口（由validation.Valid调用）
func (c *UserCreateForm) Valid(v *validation.Validation) {
	// 去除用户输入前后空白字符
	c.Name = strings.TrimSpace(c.Name)
	c.Password = strings.TrimSpace(c.Password)
	c.PasswordVerify = strings.TrimSpace(c.PasswordVerify)
	c.Birthday = strings.TrimSpace(c.Birthday)
	c.Tel = strings.TrimSpace(c.Tel)
	c.Addr = strings.TrimSpace(c.Addr)
	c.Remark = strings.TrimSpace(c.Remark)

	// 使用beego validation提供的验证器验证最小和最大长度
	v.MinSize(c.Name, 2, "name.name").Message("用户名长度必须在%d到%d之间", 2, 16)
	v.MaxSize(c.Name, 16, "name.name").Message("用户名长度必须在%d到%d之间", 2, 16)

	if _, ok := v.ErrorsMap["name"]; !ok {
		// 验证用户名是否存在
		ormer := orm.NewOrm()
		user := models.User{Name: c.Name}
		if ormer.Read(&user, "Name") == nil {
			_ = v.SetError("name", "用户名已存在")
		}
	}

	// 使用beego validation提供的验证器验证最小和最大长度
	v.MinSize(c.Password, 6, "password.password").Message("密码最小长度位%d位", 6)

	// 验证两次密码是否一致
	if c.Password != c.PasswordVerify {
		_ = v.SetError("passwordVerify", "两次输入密码不一致")
	}

	// 使用beego validation提供的验证状态值
	v.Range(c.Gender, 0, 1, "sex.sex").Message("性别选择不正确")

	// 验证时间格式
	if _, err := time.Parse("2006-01-02", c.Birthday); err != nil {
		_ = v.SetError("birthday", "出生日期不正确")
	}

	// 使用beego validation提供的电话号码格式
	v.Phone(c.Tel, "tel.tel").Message("电话不正确")

	// 使用beego validation提供的验证器验证最小和最大长度
	v.MaxSize(c.Addr, 128, "addr.addr").Message("住址长度必须在%d之内", 128)
	v.MaxSize(c.Remark, 128, "desc.desc").Message("备注长度必须在%d之内", 128)
}

// 用户修改表单
type UserModifyForm struct {
	Id       int    `form:"id,hidden,ID"`
	Name     string `form:"name,text,名称"`
	Gender   int    `form:"gender,radio,性别"`
	Birthday string `form:"birthday,date,出生日期"`
	Tel      string `form:"tel,text,电话"`
	Email    string `form:"email,text,邮箱"`
	Addr     string `form:"addr,text,住址"`
	Remark   string `form:"remark,text,备注"`

	User *models.User
}

// 用户修改表单 验证接口（由validation.Valid调用）
func (c *UserModifyForm) Valid(v *validation.Validation) {
	// 去除用户输入前后空白字符
	c.Name = strings.TrimSpace(c.Name)
	c.Birthday = strings.TrimSpace(c.Birthday)
	c.Tel = strings.TrimSpace(c.Tel)
	c.Addr = strings.TrimSpace(c.Addr)
	c.Remark = strings.TrimSpace(c.Remark)

	// 验证用户是否存在
	user := models.User{Id: c.Id}
	if orm.NewOrm().Read(&user) != nil {
		_ = v.SetError("name", "对象不存在")
		return
	} else {
		c.User = &user
	}

	// 使用beego validation提供的验证器验证最小和最大长度
	v.MinSize(c.Name, 2, "name.name").Message("用户名长度必须在%d到%d之间", 2, 16)
	v.MaxSize(c.Name, 16, "name.name").Message("用户名长度必须在%d到%d之间", 2, 16)

	// 验证用户名是否存在（排除掉自己）
	if _, ok := v.ErrorsMap["name"]; !ok {
		ormer := orm.NewOrm()
		user := models.User{Name: c.Name}
		if ormer.Read(&user, "Name") == nil && user.Id != c.Id {
			_ = v.SetError("name", "用户名已存在")
		}
	}

	// 使用beego validation提供的验证状态值
	v.Range(c.Gender, 0, 1, "sex.sex").Message("性别选择不正确")

	// 验证时间格式
	if _, err := time.Parse("2006-01-02", c.Birthday); err != nil {
		_ = v.SetError("birthday", "出生日期不正确")
	}

	// 使用beego validation提供的电话号码格式
	v.Phone(c.Tel, "tel.tel").Message("电话不正确")

	// 使用beego validation提供的验证器验证最小和最大长度
	v.MaxSize(c.Addr, 128, "addr.addr").Message("住址长度必须在%d之内", 128)
	v.MaxSize(c.Remark, 128, "desc.desc").Message("备注长度必须在%d之内", 128)
}

// 密码修改表单
type ModifyPasswordForm struct {
	OldPassword    string `form:"oldPassword,password,旧密码"`
	NewPassword    string `form:"newPassword,password,新密码"`
	PasswordVerify string `form:"passwordVerify,password,再次输入密码"`

	User *models.User
}

// 密码修改表单 验证接口（由validation.Valid调用）
func (c *ModifyPasswordForm) Valid(v *validation.Validation) {
	//  去除用户输入前后空白字符
	c.OldPassword = strings.TrimSpace(c.OldPassword)
	c.NewPassword = strings.TrimSpace(c.NewPassword)
	c.PasswordVerify = strings.TrimSpace(c.PasswordVerify)

	// 验证旧密码是否正确
	if !c.User.ValidatePassword(c.OldPassword) {
		_ = v.SetError("oldPassword", "旧密码错误")
	}

	// 使用beego validation提供的验证器验证最小和最大长度
	v.MinSize(c.NewPassword, 6, "newPassword.newPassword").Message("密码最小长度为%d位", 6)

	// 验证两次密码是否一致
	if c.NewPassword != c.PasswordVerify {
		_ = v.SetError("passwordVerify", "两次输入密码不一致")
	}
}
