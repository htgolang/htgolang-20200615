package forms

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/dcosapp/gocmdb/server/models"
	"strings"
)

// agent增加属性表单
type AgentForm struct {
	Id   int    `form:"id,text,id"`
	Name string `form:"name,text,名称"`
	Desc string `form:"desc,textarea,备注"`
}

// 用户创建表单 验证接口（由validation.Valid调用）
func (c *AgentForm) Valid(v *validation.Validation) {
	// 去除用户输入前后空白字符
	c.Name = strings.TrimSpace(c.Name)
	c.Desc = strings.TrimSpace(c.Desc)

	agent := models.DefaultAgentManager.GetById(c.Id)
	if agent == nil {
		v.SetError("id", "该终端不存在")
	}

	v.AlphaDash(c.Name, "name.name").Message("只能由大小写英文、数字、下划线和中划线组成")

	v.MinSize(c.Name, 5, "name.name").Message("终端名称长度必须在%d到%d之间", 5, 32)
	v.MaxSize(c.Name, 32, "name.name").Message("终端名称长度必须在%d到%d之间", 5, 32)

	if _, ok := v.ErrorsMap["name"]; !ok {
		// 验证终端名称是否存在
		ormer := orm.NewOrm()
		agent := models.Agent{Name: c.Name}
		if ormer.Read(&agent, "Name") == nil {
			if agent.Id != c.Id {
				_ = v.SetError("name", "终端名称已存在")
			}
		}
	}

	v.MaxSize(c.Desc, 512, "desc.desc").Message("描述长度必须在%d之内", 512)
}
