package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Agent struct {
	Id       int        `orm:"column(id);" json:"id"`
	UUID     string     `orm:"column(uuid);size(64);" json:"uuid"`
	Hostname string     `orm:"column(hostname);size(64);" json:"hostname"`
	IP       string     `orm:"column(ip);size(4096);" json:"ip"`
	OS       string     `orm:"column(os);size(64);" json:"os"`
	Arch     string     `orm:"column(arch);size(64);" json:"arch"`
	CPU      int        `orm:"column(cpu);" json:"cpu"`
	RAM      int64      `orm:"column(ram);" json:"ram"` //MB
	Disk     string     `orm:"column(disk);size(4096);" json:"disk"`
	BootTime *time.Time `orm:"column(boot_time);null;" json:"boottime"`
	Time     *time.Time `orm:"column(time);null;" json:"time"`

	HeartbeatTime *time.Time `orm:"column(heartbeat_time);null" json:"heartbeat_time"`
	CreatedTime   *time.Time `orm:"column(created_time);auto_now_add" json:"created_time"`
	DeletedTime   *time.Time `orm:"column(deleted_time);null" json:"deleted_time"`
	IsOnline      bool       `orm:"-"; json:"os_online" `
	Remark        string     `orm:"column(remark);size(4096);" json:"remark"`
}

type AgentManager struct {
}

func NewAgentManager() *AgentManager {
	return &AgentManager{}
}

func (m *AgentManager) CreateOrReplace(agent *Agent) (*Agent, bool, error) {
	ormer := orm.NewOrm()
	now := time.Now()
	orgAgent := &Agent{UUID: agent.UUID}
	if created, _, err := ormer.ReadOrCreate(orgAgent, "UUID"); err != nil {
		return nil, false, err
	} else {

		orgAgent.Hostname = agent.Hostname
		orgAgent.IP = agent.IP
		orgAgent.OS = agent.OS
		orgAgent.CPU = agent.CPU
		orgAgent.RAM = agent.RAM
		orgAgent.Disk = agent.Disk
		orgAgent.BootTime = agent.BootTime
		orgAgent.Time = agent.Time
		orgAgent.DeletedTime = nil
		orgAgent.HeartbeatTime = &now
		ormer.Update(orgAgent)
		return orgAgent, created, nil
	}
}

func (m *AgentManager) Heartbeat(uuid string) {
	orm.NewOrm().QueryTable(&Agent{}).Filter("UUID__exact", uuid).Update(orm.Params{"HeartbeatTime": time.Now()})
}

func (m *AgentManager) Query(q string, start int64, length int) ([]*Agent, int64, int64) {
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&Agent{})

	condition := orm.NewCondition()
	condition = condition.And("deleted_time__isnull", true)

	total, _ := queryset.SetCond(condition).Count()
	qtotal := total

	if q != "" {
		query := orm.NewCondition()
		query = query.Or("hostname__icontains", q)
		query = query.Or("ip__icontains", q)
		condition = condition.AndCond(query)
		qtotal, _ = queryset.SetCond(condition).Count()
	}

	var result []*Agent
	queryset.SetCond(condition).Limit(length).Offset(start).All(&result)
	return result, total, qtotal

}

func (m *AgentManager) DeleteById(pk int) error {
	orm.NewOrm().QueryTable(&Agent{}).Filter("id__exact", pk).Update(orm.Params{"deleted_time": time.Now()})
	return nil
}

func (m *AgentManager) GetById(id int) *Agent {
	result := &Agent{}
	ormer := orm.NewOrm()
	err := ormer.QueryTable(&Agent{}).Filter("id__exact", id).Filter("deleted_time__isnull", true).One(result)
	if err == nil {
		return result
	} else {
		fmt.Println(err)
	}
	return nil
}

// 根据id获取agent
func (m *AgentManager) GetByUUID(uuid string) *Agent {
	agent := &Agent{UUID: uuid, DeletedTime: nil}
	err := orm.NewOrm().QueryTable(&Agent{}).Filter("UUID__exact", uuid).Filter("deleted_time__isnull", true).One(agent)
	if err == nil {
		return agent
	}
	return nil
}

func (m *AgentManager) Modify(id int, remark string) (*Agent, error) {
	ormer := orm.NewOrm()
	if result := m.GetById(id); result != nil {
		result.Remark = remark

		if _, err := ormer.Update(result); err != nil {
			return nil, err
		}
		return result, nil
	} else {
		fmt.Println(result)
	}

	return nil, fmt.Errorf("操作对象不存在")
}

var DefaultAgentManager = NewAgentManager()

func init() {
	orm.RegisterModel(new(Agent))
}
