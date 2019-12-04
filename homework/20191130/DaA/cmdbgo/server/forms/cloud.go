package forms

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/xxdu521/cmdbgo/server/cloud"
	"github.com/xxdu521/cmdbgo/server/models"
	"strings"
)

//云平台新建form表单验证
type CloudPlatformCreateForm struct {
	Name        string		`form:"name,text,名称"`
	Type    	string		`form:"type,select,类型"`
	Addr		string		`form:"addr,text,地址"`
	Region 		string		`form:"region,text,区域"`
	AccessKey   string		`form:"access_key,text,access_key"`
	SecrectKey  string		`form:"secrect_key,text,secrect_key"`
	Remark      string		`form:"remark,textarea,备注"`
}
func (f *CloudPlatformCreateForm) Valid(v *validation.Validation) {
	f.Name 			= 	strings.TrimSpace(f.Name)
	f.Type			= 	strings.TrimSpace(f.Type)
	f.Addr 			= 	strings.TrimSpace(f.Addr)
	f.Region 		= 	strings.TrimSpace(f.Region)
	f.AccessKey 	= 	strings.TrimSpace(f.AccessKey)
	f.SecrectKey 	= 	strings.TrimSpace(f.SecrectKey)
	f.Remark 		= 	strings.TrimSpace(f.Remark)

	//调试
	//fmt.Println(f.AccessKey,f.SecrectKey)

	//验证名字格式
	v.AlphaDash(f.Name,"name.name").Message("名字只能由大小写英文、数字、下划线和中划线组成")
	v.MinSize(f.Name,5,"name.name").Message("名字长度必须在%d-%d之内", 5, 32)
	v.MaxSize(f.Name,32,"name.name").Message("名字长度必须在%d-%d之内", 5, 32)
	//验证名字是否重复
	if _, ok := v.ErrorsMap["name"]; !ok && models.DefaultCloudPlatformManager.GetByName(f.Name) != nil {
		v.SetError("name","用户名已存在")
	}

	//addr,region,accesskey,secrectkey无需验证，因为只有信息正确，才能创建成功。所以不设任何限制。

	//其他常规信息验证
	v.MaxSize(f.Remark,1024,"remark.remark").Message("备注长度必须在%d之内", 1024)

	//验证云插件类型
	if sdk, ok := cloud.DefaultManager.Cloud(f.Type); !ok {
		v.SetError("type","类型信息错误")
	} else if !v.HasErrors() {
		//通过TestConnect方法，验证用户输入的云信息是否正确
		sdk.Init(f.Addr, f.Region, f.AccessKey, f.SecrectKey)
		if sdk.TestConnect() != nil {
			v.SetError("type","配置信息错误")
		}
	}
}

//云平台修改form表单验证
type CloudPlatformModifyForm struct {
	Id 			int		`form:"id,hidden,Id"`
	Name		string	`form:"name,text,名称"`
	Type		string	`form:"type,select,类型"`
	Addr		string	`form:"addr,text,地址"`
	Region		string	`form:"region,text,区域"`
	AccessKey   string	`form:"access_key,text,access_key"`
	SecrectKey  string	`form:"secrect_key,text,secrect_key"`
	Remark		string	`form:"remark,textarea,备注"`
}

func (f *CloudPlatformModifyForm) Valid(v *validation.Validation) {
	f.Name 			= 	strings.TrimSpace(f.Name)
	f.Type			= 	strings.TrimSpace(f.Type)
	f.Addr 			= 	strings.TrimSpace(f.Addr)
	f.Region 		= 	strings.TrimSpace(f.Region)
	f.AccessKey 	= 	strings.TrimSpace(f.AccessKey)
	f.SecrectKey 	= 	strings.TrimSpace(f.SecrectKey)
	f.Remark 		= 	strings.TrimSpace(f.Remark)

	//验证操作ID是否存在
	if models.DefaultCloudPlatformManager.GetById(f.Id) == nil {
		v.SetError("error","操作的对象不存在")
	}

	//验证名字格式
	v.AlphaDash(f.Name,"name.name").Message("名字只能由大小写英文、数字、下划线和中划线组成")
	v.MinSize(f.Name,5,"name.name").Message("名字长度必须在%d-%d之内", 5, 32)
	v.MaxSize(f.Name,32,"name.name").Message("名字长度必须在%d-%d之内", 5, 32)
	//验证名字是否重复
	if _, ok := v.ErrorsMap["name"]; !ok {
		if cloulPlatform := models.DefaultCloudPlatformManager.GetByName(f.Name); cloulPlatform != nil && cloulPlatform.Id != f.Id {
			v.SetError("name","用户名已存在")
		}
	}

	//addr,region,accesskey,secrectkey无需验证，因为只有信息正确，才能创建成功。所以不设任何限制。

	//其他常规信息验证
	v.MaxSize(f.Remark,1024,"remark.remark").Message("备注长度必须在%d之内", 1024)


	//为f.AccessKey f.SecrectKey为空时，表示不修改原有key，从数据里获取key信息
	if f.AccessKey == "" && f.SecrectKey == "" {
		if cloudPlatform := models.DefaultCloudPlatformManager.GetById(f.Id); cloudPlatform == nil {
			f.AccessKey = cloudPlatform.AccessKey
			f.SecrectKey = cloudPlatform.SecrectKey
			fmt.Println(cloudPlatform.AccessKey,cloudPlatform.SecrectKey)
			fmt.Println(f.AccessKey,f.SecrectKey)
		}
	}

	//验证云插件类型

	if sdk, ok := cloud.DefaultManager.Cloud(f.Type);!ok {
		v.SetError("type","类型信息错误")
	} else if !v.HasErrors() {
		//通过TestConnect方法，验证用户输入的云信息是否正确
		sdk.Init(f.Addr, f.Region, f.AccessKey, f.SecrectKey)
		if sdk.TestConnect() != nil {
			v.SetError("type","配置信息错误")
		}
	}
}
