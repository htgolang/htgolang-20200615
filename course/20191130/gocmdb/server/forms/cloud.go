package forms

import (
	"strings"
	"github.com/astaxie/beego/validation"

	"github.com/imsilence/gocmdb/server/models"
	"github.com/imsilence/gocmdb/server/cloud"
)

type CloudPlatformCreateForm struct {
	Name string `form:"name"`
    Type string `form:"type"`
    Addr string `form:"addr"`
    AccessKey string `form:"access_key"`
    SecrectKey string `form:"secrect_key"`
    Region string `form:"region"`
    Remark string `form:"remark"`
}

func (f *CloudPlatformCreateForm) Valid(v *validation.Validation) {
	f.Name = strings.TrimSpace(f.Name)
	f.Type = strings.TrimSpace(f.Type)
	f.Addr = strings.TrimSpace(f.Addr)
	f.AccessKey = strings.TrimSpace(f.AccessKey)
	f.SecrectKey = strings.TrimSpace(f.SecrectKey)
	f.Region = strings.TrimSpace(f.Region)
	f.Remark = strings.TrimSpace(f.Remark)

	v.AlphaDash(f.Name, "name.name").Message("名字只能由大小写英文、数字、下划线和中划线组成")
	v.MinSize(f.Name, 5, "name.name").Message("名字长度必须在%d-%d之内", 5, 32)
	v.MaxSize(f.Name, 32, "name.name").Message("名字长度必须在%d-%d之内", 5, 32)

	if _, ok := v.ErrorsMap["name"]; !ok && models.DefaultCloudPlatformManager.GetByName(f.Name) != nil {
		v.SetError("name", "名称已存在")
	}

	v.MinSize(f.Addr, 1, "addr.addr").Message("地址不能为空且长度必须在%d之内", 1024)
	v.MaxSize(f.Addr, 1024, "addr.addr").Message("地址不能为空且长度必须在%d之内", 1024)

	v.MinSize(f.Region, 1, "region.region").Message("区域不能为空且长度必须在%d之内", 64)
	v.MaxSize(f.Region, 64, "region.region").Message("区域不能为空且长度必须在%d之内", 64)

	v.MinSize(f.AccessKey, 1, "access_key.access_key").Message("AccessKey不能为空且长度必须在%d之内", 1024)
	v.MaxSize(f.AccessKey, 1024, "access_key.access_key").Message("AccessKey不能为空不能为空且长度必须在%d之内", 1024)

	v.MinSize(f.SecrectKey, 1, "secrect_key.secrect_key").Message("SecrectKey不能为空且长度必须在%d之内", 1024)
	v.MaxSize(f.SecrectKey, 1024, "secrect_key.secrect_key").Message("SecrectKey不能为空且长度必须在%d之内", 1024)

	v.MaxSize(f.Remark, 1024, "remark.remark").Message("备注长度必须在%d之内", 1024)

	if sdk, ok := cloud.DefaultManager.Cloud(f.Type); !ok {
		v.SetError("type", "类型错误")
	} else if !v.HasErrors() {
		sdk.Init(f.Addr, f.Region, f.AccessKey, f.SecrectKey)
		if sdk.TestConnect() != nil {
			v.SetError("type", "配置参数错误")
		}
	}
}