package forms

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/xxdu521/cmdbgo/server/models"
	"strings"
	"time"
)

//用户管理创建表单验证
type UserCreateForm struct {
	Name 			string	`form:"name,text,名称"`
	Password 		string	`form:"password,password,密码"`
	PasswordVerify 	string	`form:"passwordVerify,password,再次输入密码"`
	Gender 			int		`form:"gender,radio,性别"`
	Birthday 		string	`form:"birthday,date,出生日期"`
	Tel 			string	`form:"tel,text,电话"`
	Email 			string	`form:"email,text,邮箱"`
	Addr 			string	`form:"addr,text,住址"`
	Remark 			string	`form:"remark,text,备注"`
}
func (f *UserCreateForm) Valid(v *validation.Validation) {
	f.Name = strings.TrimSpace(f.Name)
	f.Password = strings.TrimSpace(f.Password)
	f.PasswordVerify = strings.TrimSpace(f.PasswordVerify)
	f.Birthday = strings.TrimSpace(f.Birthday)
	f.Tel = strings.TrimSpace(f.Tel)
	f.Email = strings.TrimSpace(f.Email)
	f.Addr = strings.TrimSpace(f.Addr)
	f.Remark = strings.TrimSpace(f.Remark)

	v.MinSize(f.Name,5,"name.name").Message("用户名长度必须在%d到%d之间",5,32)
	v.MaxSize(f.Name,32,"name.name").Message("用户名长度必须在%d到%d之间",5,32)

	if _, ok := v.ErrorsMap["name"]; !ok {
		ormer := orm.NewOrm()
		user := models.User{Name: f.Name}
		if ormer.Read(&user,"Name") == nil {
			v.SetError("name","用户名已存在")
		}
	}

	v.MinSize(f.Password,6,"password.password").Message("密码长度最少%d位",6)

	if f.Password != f.PasswordVerify {
		v.SetError("passwordVerify","密码不一致")
	}

	v.Range(f.Gender,0,1,"sex.sex").Message("性别选择错误")

	fmt.Println(f.Birthday,f.Name)
	if _, err := time.Parse("01/02/2006",f.Birthday); err != nil {
		v.SetError("birthday","出生日期不正确")
	}

	v.Phone(f.Tel,"tel.tel").Message("手机号码不正确")

	v.MaxSize(f.Addr,128,"addr.addr").Message("住址长度必须在%d之间",128)
	v.MaxSize(f.Remark,128,"desc.desc").Message("备注长度必须在%d之间",128)
}

//用户管理修改表单验证
type UserModifyForm struct {
	Id       		int		`form:"id,hidden,ID"`
	Name 			string	`form:"name,text,名称"`
	Gender 			int		`form:"gender,radio,性别"`
	Birthday 		string	`form:"birthday,date,出生日期"`
	Tel 			string	`form:"tel,text,电话"`
	Email 			string	`form:"email,text,邮箱"`
	Addr 			string	`form:"addr,text,住址"`
	Remark 			string	`form:"remark,text,备注"`

	User			*models.User
}
func (f *UserModifyForm) Valid(v *validation.Validation) {
	// 去除用户输入前后空白字符
	f.Name = strings.TrimSpace(f.Name)
	f.Birthday = strings.TrimSpace(f.Birthday)
	f.Tel = strings.TrimSpace(f.Tel)
	f.Email = strings.TrimSpace(f.Email)
	f.Addr = strings.TrimSpace(f.Addr)
	f.Remark = strings.TrimSpace(f.Remark)

	// 验证用户是否存在
	user := models.User{Id: f.Id}
	if orm.NewOrm().Read(&user) != nil {
		v.SetError("name","用户名不存在")
		return
	} else {
		f.User = &user
	}

	v.MinSize(f.Name,5,"name.name").Message("用户名长度必须在%d到%d之间",5,32)
	v.MaxSize(f.Name,32,"name.name").Message("用户名长度必须在%d到%d之间",5,32)

	//验证用户名是否存在，（去掉自己）
	if _,ok := v.ErrorsMap["name"]; !ok {
		ormer := orm.NewOrm()
		user := models.User{Name: f.Name}
		if ormer.Read(&user,"Name") == nil && user.Id != f.Id {
			v.SetError("name","用户名已经存在")
		}
	}

	// 使用beego validation提供的验证状态值
	v.Range(f.Gender,0,1,"sex.sex").Message("性别选择错误")

	// 验证时间格式
	fmt.Println(f.Birthday)
	if aaa, err := time.Parse("01/02/2006",f.Birthday); err != nil {
		v.SetError("birthday", "出生日期不正确")
		fmt.Println(aaa,err)
	}
	
	// 使用beego validation提供的电话号码格式
	v.Phone(f.Tel,"tel.tel").Message("手机号码不正确")

	//还要验证邮件

	// 使用beego validation提供的验证器验证最小和最大长度
	v.MaxSize(f.Addr,128,"addr.addr").Message("住址长度必须在%d之间",128)
	v.MaxSize(f.Remark,128,"desc.desc").Message("备注长度必须在%d之间",128)
}

//用户管理密码表单验证
type UserSetPasswordForm struct {
	OldPassword		string `form:"oldPassword,password,旧密码"`
	NewPassword		string `form:"newPassword,password,新密码"`
	PasswordVerify	string `form:"passwordVerify,password,密码确认"`

	User			*models.User
}
func (f *UserSetPasswordForm) Valid(v *validation.Validation) {
	//  去除用户输入前后空白字符
	f.OldPassword = strings.TrimSpace(f.OldPassword)
	f.NewPassword = strings.TrimSpace(f.NewPassword)
	f.PasswordVerify = strings.TrimSpace(f.PasswordVerify)

	if !f.User.ValidatePassword(f.OldPassword) {
		v.SetError("oldPassword","密码错误")
	}

	v.MinSize(f.NewPassword,6,"newPassword,newPassword").Message("密码最小长度为%d位",6)

	if f.NewPassword != f.PasswordVerify {
		v.SetError("passwordVerify","两次输入的密码不一致")
	}
}
