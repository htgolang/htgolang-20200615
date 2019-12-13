package forms

import (
	"strings"

	"github.com/astaxie/beego/validation"
	"github.com/imsilence/gocmdb/server/cloud"
	"github.com/imsilence/gocmdb/server/models"
)

type CloudPlatformCreateForm struct {
	Name       string `form:"name,text,名称"`
	Type       string `form:"type,select,类型"`
	Addr       string `form:"addr,text,地址"`
	Region     string `form:"region,text,地址"`
	AccessKey  string `form:"access_key,text,access_key"`
	SecrectKey string `form:"secrect_key,text,secrect_key"`
	Remark     string `form:"remark,text,备注"`
}

func (f *CloudPlatformCreateForm) Valid(v *validation.Validation) {
	f.Name = strings.TrimSpace(f.Name)
	f.Type = strings.TrimSpace(f.Type)
	f.Region = strings.TrimSpace(f.Region)
	f.Addr = strings.TrimSpace(f.Addr)
	f.Remark = strings.TrimSpace(f.Remark)
	f.AccessKey = strings.TrimSpace(f.AccessKey)
	f.SecrectKey = strings.TrimSpace(f.SecrectKey)

	v.AlphaDash(f.Name, "name.name").Message("用户名只能由数字、英文字母、中划线和下划线组成")
	v.MinSize(f.Name, 5, "name.name").Message("用户名长度必须在%d-%d之内", 5, 32)
	v.MaxSize(f.Name, 32, "name.name").Message("用户名长度必须在%d-%d之内", 5, 32)

	if _, ok := v.ErrorsMap["name"]; !ok && models.DefaultCloudPlatformManager.GetByName(f.Name) != nil {
		v.SetError("name", "名称已存在")
	}

	v.MinSize(f.Addr, 1, "addr.addr").Message("地址不能为空，且长度必须在%d之内", 1024)
	v.MaxSize(f.Addr, 1024, "addr.addr").Message("地址不能为空，且长度必须在%d之内", 1024)

	v.MinSize(f.Region, 1, "region.region").Message("Region不能为空，且长度必须在%d之内", 64)
	v.MaxSize(f.Region, 64, "region.region").Message("Region不能为空，且长度必须在%d之内", 64)

	v.MinSize(f.AccessKey, 1, "access_key.access_key").Message("AccessKey不能为空，且长度必须在%d之内", 1024)
	v.MaxSize(f.AccessKey, 1024, "access_key.access_key").Message("AccessKey不能为空，且长度必须在%d之内", 1024)

	v.MinSize(f.SecrectKey, 1, "secrect_key.secrect_key").Message("SecrectKey不能为空，且长度必须在%d之内", 1024)
	v.MaxSize(f.SecrectKey, 1024, "secrect_key.secrect_key").Message("SecrectKey不能为空，且长度必须在%d之内", 1024)

	v.MaxSize(f.Remark, 1024, "remark.remark").Message("Remark不能为空，且长度必须在%d之内", 1024)

	// 需要验证类型的访问
	if sdk, ok := cloud.DefaultManager.Cloud(f.Type); !ok {
		v.SetError("type", "类型错误")
	} else if !v.HasErrors() {
		sdk.Init(f.Addr, f.Region, f.AccessKey, f.SecrectKey)
		if sdk.TestConnect() != nil {
			v.SetError("type", "配置参数错误")
		}
	}
}

type CloudPlatformModifyForm struct {
	Id         int    `form:"id"`
	Name       string `form:"name,text,名称"`
	Type       string `form:"type,select,类型"`
	Addr       string `form:"addr,text,地址"`
	Region     string `form:"region,text,地址"`
	AccessKey  string `form:"access_key,text,access_key"`
	SecrectKey string `form:"secrect_key,text,secrect_key"`
	Remark     string `form:"remark,text,备注"`
}

func (f *CloudPlatformModifyForm) Valid(v *validation.Validation) {
	f.Name = strings.TrimSpace(f.Name)
	f.Type = strings.TrimSpace(f.Type)
	f.Addr = strings.TrimSpace(f.Addr)
	f.Region = strings.TrimSpace(f.Region)
	f.AccessKey = strings.TrimSpace(f.AccessKey)
	f.SecrectKey = strings.TrimSpace(f.SecrectKey)
	f.Remark = strings.TrimSpace(f.Remark)

	if models.DefaultCloudPlatformManager.GetById(f.Id) == nil {
		v.SetError("error", "操作对象不存在!")
		return
	}

	v.AlphaDash(f.Name, "name.name").Message("用户名只能由数字、英文字母、中划线和下划线组成")
	v.MinSize(f.Name, 5, "name.name").Message("用户名长度必须在%d-%d之内", 5, 32)
	v.MaxSize(f.Name, 32, "name.name").Message("用户名长度必须在%d-%d之内", 5, 32)

	cloudPlatform := &models.CloudPlatform{}
	if _, ok := v.ErrorsMap["name"]; !ok {
		if cloudPlatform = models.DefaultCloudPlatformManager.GetByName(f.Name); cloudPlatform != nil && cloudPlatform.Id != f.Id {
			v.SetError("name", "用户名已存在")
		}
	}

	v.MinSize(f.Addr, 1, "addr.addr").Message("地址不能为空，且长度必须在%d之内", 1024)
	v.MaxSize(f.Addr, 1024, "addr.addr").Message("地址不能为空，且长度必须在%d之内", 1024)

	v.MinSize(f.Region, 1, "region.region").Message("Region不能为空，且长度必须在%d之内", 64)
	v.MaxSize(f.Region, 64, "region.region").Message("Region不能为空，且长度必须在%d之内", 64)

	if f.AccessKey != "" {
		v.MinSize(f.AccessKey, 1, "access_key.access_key").Message("AccessKey不能为空，且长度必须在%d之内", 1024)
		v.MaxSize(f.AccessKey, 1024, "access_key.access_key").Message("AccessKey不能为空，且长度必须在%d之内", 1024)
	}

	if f.SecrectKey != "" {
		v.MinSize(f.SecrectKey, 1, "secrect_key.secrect_key").Message("SecrectKey不能为空，且长度必须在%d之内", 1024)
		v.MaxSize(f.SecrectKey, 1024, "secrect_key.secrect_key").Message("SecrectKey不能为空，且长度必须在%d之内", 1024)
	}

	v.MaxSize(f.Remark, 1024, "remark.remark").Message("Remark不能为空，且长度必须在%d之内", 1024)

	// 需要验证类型的访问
	if sdk, ok := cloud.DefaultManager.Cloud(f.Type); !ok {
		v.SetError("type", "类型错误")
	} else if !v.HasErrors() {
		accessKey := f.AccessKey
		if accessKey == "" {
			accessKey = cloudPlatform.AccessKey
		}
		secrectKey := f.SecrectKey
		if secrectKey == "" {
			secrectKey = cloudPlatform.SecrectKey
		}

		sdk.Init(f.Addr, f.Region, accessKey, secrectKey)
		if sdk.TestConnect() != nil {
			v.SetError("type", "配置参数错误")
		}
	}
}
