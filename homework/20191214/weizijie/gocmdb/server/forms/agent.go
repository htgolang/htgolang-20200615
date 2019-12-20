package forms

import (
	"strings"

	"github.com/astaxie/beego/validation"
	"github.com/imsilence/gocmdb/server/models"
)

type AgentModifyForm struct {
	Id       int    `form:"id"`
	Hostname string `form:"hostname,text,主机名"`
	Remark   string `form:"remark,text,备注"`
}

func (f *AgentModifyForm) Valid(v *validation.Validation) {
	f.Remark = strings.TrimSpace(f.Remark)
	f.Hostname = strings.TrimSpace(f.Hostname)
	if models.DefaultAgentManager.GetById(f.Id) == nil {
		v.SetError("error", "操作对象不存在!")
		return
	}

	v.MaxSize(f.Remark, 1024, "remark.remark").Message("Remark不能为空，且长度必须在%d之内", 1024)

}
