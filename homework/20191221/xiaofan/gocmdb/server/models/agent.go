package models

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type Agent struct {
	Id       int        `orm:"column(id);"json:"id"`
	Name     string     `orm:"column(name);size(64);null;"json:"name"`
	UUID     string     `orm:"column(uuid);size(64);"json:"uuid"`
	Hostname string     `orm:"column(hostname);size(64);"json:"hostname"`
	IP       string     `orm:"column(ip);size(4096);"json:"ip"`
	OS       string     `orm:"column(os);size(64);"json:"os"`
	Arch     string     `orm:"column(arch);size(64);"json:"arch"`
	CPU      int        `orm:"column(cpu);"json:"cpu"`
	RAM      int64      `orm:"column(ram);"json:"ram"`
	Disk     string     `orm:"column(disk);size(4096);"json:"disk"`
	Desc     string     `orm:"column(desc);size(1024);null;"json:"desc"`
	BootTime *time.Time `orm:"column(boot_time);null;"json:"boot_time"`
	Time     *time.Time `orm:"column(time);null;"json:"time"`

	HeartbeatTime *time.Time         `orm:"column(heartbeat_time);null;"json:"heartbeat_time"`
	CreatedTime   *time.Time         `orm:"column(created_time);auto_now_add;"json:"created_time"`
	DeletedTime   *time.Time         `orm:"column(deleted_time);null;"json:"deleted_time"`
	IsOnline      bool               `json:"is_online"`
	IPList        []string           `orm:"-";json:"ip_list"`
	Disks         map[string]float64 `orm:"-";json:"disks"`
}

func (a *Agent) Path() {
	if time.Since(*a.HeartbeatTime) < 5*time.Minute {
		a.IsOnline = true
	}
	if a.IP != "" {
		json.Unmarshal([]byte(a.IP), &a.IPList)
	}
	if a.Disk != "" {
		json.Unmarshal([]byte(a.Disk), &a.Disks)
	}

}

type AgentManager struct{}

func NewAgentManager() *AgentManager {
	return &AgentManager{}
}

// 创建或更新agent信息
func (m *AgentManager) CreateOrReplace(agent *Agent) (*Agent, bool, error) {
	now := time.Now()
	ormer := orm.NewOrm()
	orgAgent := &Agent{UUID: agent.UUID}
	if created, _, err := ormer.ReadOrCreate(orgAgent, "UUID"); err != nil {
		fmt.Println(err.Error())
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
		orgAgent.Arch = agent.Arch

		orgAgent.HeartbeatTime = &now
		ormer.Update(orgAgent)
		return orgAgent, created, nil
	}
}

// 更新agent的心跳时间
func (m *AgentManager) Heartbeat(uuid string) {
	ormer := orm.NewOrm()
	ormer.QueryTable(&Agent{}).Filter("UUID__exact", uuid).Update(orm.Params{"HeartbeatTime": time.Now()})
}

// 查询agent
func (m *AgentManager) Query(q string, start int64, length int) ([]*Agent, int64, int64) {
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&Agent{})
	condition := orm.NewCondition()
	condition = condition.And("deleted_time__isnull", true)
	total, _ := queryset.SetCond(condition).Count()

	if q != "" {
		query := orm.NewCondition()
		query = query.Or("hostname__icontains", q)
		query = query.Or("ip__icontains", q)
		condition = condition.AndCond(query)
	}

	var result []*Agent
	qtotal, _ := queryset.SetCond(condition).Count()
	_, _ = queryset.SetCond(condition).Limit(length).Offset(start).All(&result)

	for _, agent := range result {
		agent.Path()
	}
	return result, total, qtotal
}

// 根据id获取agent
func (m *AgentManager) GetById(id int) *Agent {
	cloud := &Agent{Id: id, DeletedTime: nil}
	err := orm.NewOrm().QueryTable(&Agent{}).Filter("Id__exact", id).Filter("deleted_time__isnull", true).One(cloud)
	if err == nil {
		return cloud
	}
	return nil
}

// 根据id获取agent
func (m *AgentManager) GetByUUID(uuid string) *Agent {
	cloud := &Agent{UUID: uuid, DeletedTime: nil}
	err := orm.NewOrm().QueryTable(&Agent{}).Filter("UUID__exact", uuid).Filter("deleted_time__isnull", true).One(cloud)
	if err == nil {
		return cloud
	}
	return nil
}

// 删除agent
func (m *AgentManager) DeleteById(pk int) (int64, error) {
	now := time.Now()
	result, err := orm.NewOrm().QueryTable(&Agent{}).RelatedSel().Filter("id__exact", pk).Update(orm.Params{"DeletedTime": &now})
	return result, err
}

// 修改agent
func (m *AgentManager) Modify(id int, name, desc string) (*Agent, error) {
	agent := &Agent{Id: id}
	ormer := orm.NewOrm()
	err := ormer.Read(agent)
	if err == nil {
		agent.Name = name
		agent.Desc = desc
		if _, err := ormer.Update(agent); err == nil {
			return agent, nil
		}
	}
	return &Agent{}, err
}

var DefaultAgentManager = NewAgentManager()

func init() {
	orm.RegisterModel(&Agent{})
}
