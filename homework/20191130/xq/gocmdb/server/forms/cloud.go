package forms

import (
	"github.com/astaxie/beego/validation"
	"github.com/xlotz/gocmdb/server/models"
	"github.com/xlotz/gocmdb/server/cloud"
	"strings"

)

type CloudPlatCreateForm struct {
	Name           string    `form:"name"`
	Type      		string    `form:"type"`
	Region         string      `form:"region"`
	Addr           string    `form:"addr"`
	Remark         string    `form:"remark"`
	AccessKey string `form:"access_key"`
	SecrectKey string `form:"secrect_key"`

}

func (f *CloudPlatCreateForm) Valid(v *validation.Validation) {
	f.Name = strings.TrimSpace(f.Name)
	f.Region = strings.TrimSpace(f.Region)
	f.Type = strings.TrimSpace(f.Type)
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
	v.MaxSize(f.Region, 64, "region.region").Message("长度必须在64个字符之内")


	//
	v.MaxSize(f.Addr, 1024, "addr.addr").Message("住址长度必须在1024个字符之内")
	v.MaxSize(f.Remark, 1024, "remark.remark").Message("备注长度必须在1024个字符之内")
	v.MaxSize(f.AccessKey, 1024, "access_key.access_key").Message("长度必须在1024个字符之内")
	v.MaxSize(f.SecrectKey, 1024, "secrect_key.secrect_key").Message("长度必须在1024个字符之内")

	// 类型验证
	if sdk, ok := cloud.DefaultManager.Cloud(f.Type); !ok {
		v.SetError("type", "类型错误")
	}else if v.HasErrors() {
		sdk.Init(f.Addr, f.Region, f.AccessKey, f.SecrectKey)
		if sdk.TestConnect() != nil {
			_ = v.SetError("type", "参数错误")
		}
	}

}


type CloudPlatModifyForm struct {
	Id         int       `form:"id"`
	Name       string    `form:"name"`
	Type      		string    `form:"type"`
	Region         string      `form:"region"`
	Addr           string    `form:"addr"`
	Remark         string    `form:"remark"`
	AccessKey string `form:"access_key"`
	SecrectKey string `form:"secrect_key"`
}

func (f *CloudPlatModifyForm) Valid(v *validation.Validation) {

	f.Name = strings.TrimSpace(f.Name)
	f.Region = strings.TrimSpace(f.Region)
	f.Type = strings.TrimSpace(f.Type)
	f.Addr = strings.TrimSpace(f.Addr)
	f.Remark = strings.TrimSpace(f.Remark)
	f.AccessKey = strings.TrimSpace(f.AccessKey)
	f.SecrectKey = strings.TrimSpace(f.SecrectKey)

	if models.DefaultCloudPlatformManager.GetById(f.Id) == nil {
		v.SetError("error", "操作对象不存在")
		return
	}

	v.AlphaDash(f.Name, "name.name").Message("用户名只能由数字、英文字母、中划线和下划线组成")
	v.MinSize(f.Name, 5, "name.name").Message("用户名长度必须在%d-%d之内", 5, 32)
	v.MaxSize(f.Name, 32, "name.name").Message("用户名长度必须在%d-%d之内", 5, 32)

	if _, ok := v.ErrorsMap["name"]; !ok {

		if c := models.DefaultCloudPlatformManager.GetByName(f.Name); c != nil && c.Id != f.Id{
			v.SetError("name", "名称已存在")
		}

	}
	v.MaxSize(f.Region, 64, "region.region").Message("长度必须在64个字符之内")

	//
	v.MaxSize(f.Addr, 1024, "addr.addr").Message("住址长度必须在1024个字符之内")
	v.MaxSize(f.Remark, 1024, "remark.remark").Message("备注长度必须在1024个字符之内")
	v.MaxSize(f.AccessKey, 1024, "access_key.access_key").Message("长度必须在1024个字符之内")
	v.MaxSize(f.SecrectKey, 1024, "secrect_key.secrect_key").Message("长度必须在1024个字符之内")

	// 类型验证
	if sdk, ok := cloud.DefaultManager.Cloud(f.Type); !ok {
		v.SetError("type", "类型错误")
	}else if v.HasErrors() {
		sdk.Init(f.Addr, f.Region, f.AccessKey, f.SecrectKey)
		if sdk.TestConnect() != nil {
			_ = v.SetError("type", "参数错误")
		}
	}

}
