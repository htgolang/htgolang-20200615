package forms

import (
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"

	"todolist/models"
)

// 用户创建表单
type UserCreateForm struct {
	Name           string `form:"name,text,名称"`
	Password       string `form:"password,password,密码"`
	PasswordVerify string `form:"passwordVerify,password,再次输入密码"`
	Sex            int    `form:"sex,radio,性别"`
	Birthday       string `form:"birthday,date,出生日期"`
	Tel            string `form:"tel,text,电话"`
	Addr           string `form:"addr,text,住址"`
	Desc           string `form:"desc,text,备注"`
}

// 用户创建表单 验证接口（由validation.Valid调用）
func (this *UserCreateForm) Valid(v *validation.Validation) {
	this.Name = strings.TrimSpace(this.Name)
	this.Password = strings.TrimSpace(this.Password)
	this.PasswordVerify = strings.TrimSpace(this.PasswordVerify)
	this.Birthday = strings.TrimSpace(this.Birthday)
	this.Tel = strings.TrimSpace(this.Tel)
	this.Addr = strings.TrimSpace(this.Addr)
	this.Desc = strings.TrimSpace(this.Desc)

	// 使用beego validation提供的验证器验证最小和最大长度
	v.MinSize(this.Name, 2, "name.name").Message("用户名长度必须在%d到%d之间", 2, 16)
	v.MaxSize(this.Name, 16, "name.name").Message("用户名长度必须在%d到%d之间", 2, 16)

	if _, ok := v.ErrorsMap["name"]; !ok {
		// 验证用户名是否存在
		ormer := orm.NewOrm()
		user := models.User{Name: this.Name}
		if ormer.Read(&user, "Name") == nil {
			v.SetError("name", "用户名已存在")
		}
	}

	// 使用beego validation提供的验证器验证最小和最大长度
	v.MinSize(this.Password, 6, "password.password").Message("密码最小长度位%d位", 6)

	// 验证两次密码是否一致
	if this.Password != this.PasswordVerify {
		v.SetError("passwordVerify", "两次输入密码不一致")
	}

	// 使用beego validation提供的验证状态值
	v.Range(this.Sex, 0, 1, "sex.sex").Message("性别选择不正确")

	// 验证时间格式
	if _, err := time.Parse("2006-01-02", this.Birthday); err != nil {
		v.SetError("birthday", "出生日期不正确")
	}

	// 使用beego validation提供的电话号码格式
	v.Phone(this.Tel, "tel.tel").Message("电话不正确")

	// 使用beego validation提供的验证器验证最小和最大长度
	v.MaxSize(this.Addr, 128, "addr.addr").Message("住址长度必须在%d之内", 128)
	v.MaxSize(this.Desc, 128, "desc.desc").Message("备注长度必须在%d之内", 128)
}


// 用户修改表单
type UserModifyForm struct {
	Id       int    `form:"id,hidden,ID"`
	Name     string `form:"name,text,名称"`
	Sex      int    `form:"sex,radio,性别"`
	Birthday string `form:"birthday,date,出生日期"`
	Tel      string `form:"tel,text,电话"`
	Addr     string `form:"addr,text,住址"`
	Desc     string `form:"desc,text,备注"`

	User *models.User
}

// 用户修改表单 验证接口（由validation.Valid调用）
func (this *UserModifyForm) Valid(v *validation.Validation) {
	// 去除用户输入前后空白字符
	this.Name = strings.TrimSpace(this.Name)
	this.Birthday = strings.TrimSpace(this.Birthday)
	this.Tel = strings.TrimSpace(this.Tel)
	this.Addr = strings.TrimSpace(this.Addr)
	this.Desc = strings.TrimSpace(this.Desc)

	// 验证用户是否存在
	user := models.User{Id: this.Id}
	if orm.NewOrm().Read(&user) != nil {
		v.SetError("name", "对象不存在")
		return
	} else {
		this.User = &user
	}

	// 使用beego validation提供的验证器验证最小和最大长度
	v.MinSize(this.Name, 2, "name.name").Message("用户名长度必须在%d到%d之间", 2, 16)
	v.MaxSize(this.Name, 16, "name.name").Message("用户名长度必须在%d到%d之间", 2, 16)

	// 验证用户名是否存在（排除掉自己）
	if _, ok := v.ErrorsMap["name"]; !ok {
		ormer := orm.NewOrm()
		user := models.User{Name: this.Name}
		if ormer.Read(&user, "Name") == nil && user.Id != this.Id {
			v.SetError("name", "用户名已存在")
		}
	}

	// 使用beego validation提供的验证状态值
	v.Range(this.Sex, 0, 1, "sex.sex").Message("性别选择不正确")

	// 验证时间格式
	if _, err := time.Parse("2006-01-02", this.Birthday); err != nil {
		v.SetError("birthday", "出生日期不正确")
	}

	// 使用beego validation提供的电话号码格式
	v.Phone(this.Tel, "tel.tel").Message("电话不正确")

	// 使用beego validation提供的验证器验证最小和最大长度
	v.MaxSize(this.Addr, 128, "addr.addr").Message("住址长度必须在%d之内", 128)
	v.MaxSize(this.Desc, 128, "desc.desc").Message("备注长度必须在%d之内", 128)
}

// 密码修改表单
type ModifyPasswordForm struct {
	OldPassword    string `form:"oldPassword,password,旧密码"`
	NewPassword    string `form:"newPassword,password,新密码"`
	PasswordVerify string `form:"passwordVerify,password,再次输入密码"`

	User *models.User
}

// 密码修改表单 验证接口（由validation.Valid调用）
func (f *ModifyPasswordForm) Valid(v *validation.Validation) {
	//  去除用户输入前后空白字符
	f.OldPassword = strings.TrimSpace(f.OldPassword)
	f.NewPassword = strings.TrimSpace(f.NewPassword)
	f.PasswordVerify = strings.TrimSpace(f.PasswordVerify)

	// 验证旧密码是否正确
	if !f.User.ValidatePassword(f.OldPassword) {
		v.SetError("oldPassword", "密码错误")
	}

	// 使用beego validation提供的验证器验证最小和最大长度
	v.MinSize(f.NewPassword, 6, "newPassword.newPassword").Message("密码最小长度位%d位", 6)

	// 验证两次密码是否一致
	if f.NewPassword != f.PasswordVerify {
		v.SetError("passwordVerify", "两次输入密码不一致")
	}
}
