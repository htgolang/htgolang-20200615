package forms

import (
	"github.com/astaxie/beego/validation"
	"github.com/dcosapp/gocmdb/server/cloud"
	"github.com/dcosapp/gocmdb/server/models"
	"strings"
)

// 云平台创建表单
type CloudPlatformCreateForm struct {
	Name      string `form:"name,text,名称"`
	Type      string `form:"type,select,类型"`
	Addr      string `form:"addr,text,地址"`
	Region    string `form:"region,text,地址"`
	AccessKey string `form:"access_key,text,access_key"`
	SecretKey string `form:"secret_key,text,secret_key"`
	Remark    string `form:"remark,textarea,备注"`
}

// 云平台创建表单, 验证接口（由validation.Valid调用）
func (c *CloudPlatformCreateForm) Valid(v *validation.Validation) {
	// 去除用户输入前后空白字符
	c.Name = strings.TrimSpace(c.Name)
	c.Type = strings.TrimSpace(c.Type)
	c.Region = strings.TrimSpace(c.Region)
	c.AccessKey = strings.TrimSpace(c.AccessKey)
	c.SecretKey = strings.TrimSpace(c.SecretKey)
	c.Addr = strings.TrimSpace(c.Addr)
	c.Remark = strings.TrimSpace(c.Remark)

	v.AlphaDash(c.Name, "name.name").Message("只能由大小写英文、数字、下划线和中划线组成")

	v.MinSize(c.Name, 2, "name.name").Message("名称长度必须在%d到%d之间", 2, 32)
	v.MaxSize(c.Name, 32, "name.name").Message("名称长度必须在%d到%d之间", 2, 32)

	if _, ok := v.ErrorsMap["name"]; !ok && models.DefaultCloudPlatformManager.GetByName(c.Name) != nil {
		_ = v.SetError("name", "名称已存在")
	}

	v.MinSize(c.Addr, 1, "addr.addr").Message("地址不能为空且长度必须在%d到%d之间", 1, 1024)
	v.MaxSize(c.Addr, 1024, "addr.addr").Message("地址长度必须在%d到%d之间", 1, 1024)

	v.MinSize(c.Region, 1, "region.region").Message("区域不能为空且长度必须在%d到%d之间", 1, 64)
	v.MaxSize(c.Region, 64, "region.region").Message("区域长度必须在%d到%d之间", 1, 64)

	v.MinSize(c.SecretKey, 1, "secret_key.secret_key").Message("SecretKey不能为空且长度必须在%d到%d之间", 1, 1024)
	v.MaxSize(c.SecretKey, 1024, "secret_key.secret_key").Message("SecretKey长度必须在%d到%d之间", 1, 1024)

	v.MinSize(c.AccessKey, 1, "access_key.access_key").Message("AccessKey不能为空且长度必须在%d到%d之间", 1, 1024)
	v.MaxSize(c.AccessKey, 1024, "access_key.access_key").Message("AccessKey长度必须在%d到%d之间", 1, 1024)

	v.MaxSize(c.Remark, 1024, "remark.remark").Message("备注长度必须在%d之内", 1024)

	// 判断Cloud Manager是否有这个插件名
	if sdk, ok := cloud.DefaultManager.Cloud(c.Type); !ok {
		_ = v.SetError("type", "类型错误")
	} else if !v.HasErrors() { // 如果前面都验证通过，这里就开始测试连接
		sdk.Init(c.Addr, c.Region, c.AccessKey, c.SecretKey)
		if sdk.TestConnect() != nil {
			_ = v.SetError("type", "配置参数错误")
		}
	}
}

// 云平台修改表单
type CloudPlatformModifyForm struct {
	Id        int    `form:"id,text,id"`
	Name      string `form:"name,text,名称"`
	Type      string `form:"type,select,类型"`
	Addr      string `form:"addr,text,地址"`
	Region    string `form:"region,text,地址"`
	AccessKey string `form:"access_key,text,access_key"`
	SecretKey string `form:"secret_key,text,secret_key"`
	Remark    string `form:"remark,textarea,备注"`
}

// 云平台修改表单, 验证接口（由validation.Valid调用）
func (c *CloudPlatformModifyForm) Valid(v *validation.Validation) {
	// 去除用户输入前后空白字符
	c.Name = strings.TrimSpace(c.Name)
	c.Type = strings.TrimSpace(c.Type)
	c.Region = strings.TrimSpace(c.Region)
	c.AccessKey = strings.TrimSpace(c.AccessKey)
	c.SecretKey = strings.TrimSpace(c.SecretKey)
	c.Addr = strings.TrimSpace(c.Addr)
	c.Remark = strings.TrimSpace(c.Remark)

	platform := models.DefaultCloudPlatformManager.GetById(c.Id)
	if platform == nil {
		_ = v.SetError("id", "该云平台不存在")
		return
	}

	v.AlphaDash(c.Name, "name.name").Message("只能由大小写英文、数字、下划线和中划线组成")

	v.MinSize(c.Name, 5, "name.name").Message("名称长度必须在%d到%d之间", 5, 32)
	v.MaxSize(c.Name, 32, "name.name").Message("名称长度必须在%d到%d之间", 5, 32)

	if c.Name != platform.Name {
		if _, ok := v.ErrorsMap["name"]; !ok && models.DefaultCloudPlatformManager.GetByName(c.Name) != nil {
			_ = v.SetError("name", "名称已存在")
		}
	}

	v.MinSize(c.Addr, 1, "addr.addr").Message("地址不能为空且长度必须在%d到%d之间", 1, 1024)
	v.MaxSize(c.Addr, 1024, "addr.addr").Message("地址长度必须在%d到%d之间", 1, 1024)

	v.MinSize(c.Region, 1, "region.region").Message("区域不能为空且长度必须在%d到%d之间", 1, 64)
	v.MaxSize(c.Region, 64, "region.region").Message("区域长度必须在%d到%d之间", 1, 64)

	v.MinSize(c.SecretKey, 0, "secret_key.secret_key").Message("SecretKey不能为空且长度必须在%d到%d之间", 0, 1024)
	v.MaxSize(c.SecretKey, 1024, "secret_key.secret_key").Message("SecretKey长度必须在%d到%d之间", 0, 1024)

	v.MinSize(c.AccessKey, 0, "access_key.access_key").Message("AccessKey不能为空且长度必须在%d到%d之间", 0, 1024)
	v.MaxSize(c.AccessKey, 1024, "access_key.access_key").Message("AccessKey长度必须在%d到%d之间", 0, 1024)

	v.MaxSize(c.Remark, 1024, "remark.remark").Message("备注长度必须在%d之内", 1024)

	// 判断Cloud Manager是否有这个插件名
	if sdk, ok := cloud.DefaultManager.Cloud(c.Type); !ok {
		_ = v.SetError("type", "类型错误")
	} else if !v.HasErrors() { // 如果前面都验证通过，这里就开始测试连接
		accKey, secKey := c.AccessKey, c.SecretKey
		if accKey == "" {
			accKey = platform.AccessKey
		}
		if secKey == "" {
			secKey = platform.SecretKey
		}
		sdk.Init(c.Addr, c.Region, accKey, secKey)
		if sdk.TestConnect() != nil {
			_ = v.SetError("type", "配置参数错误")
		}
		sdk.GetInstance()
	}
}
