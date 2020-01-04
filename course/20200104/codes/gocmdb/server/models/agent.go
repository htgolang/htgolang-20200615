package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"encoding/json"
)

type Agent struct {
	Id       int        `orm:"column(id);" json:"id"`
	UUID     string     `orm:"column(uuid);size(64);" json:"uuid"`
	Hostname string     `orm:"column(hostname);size(64);" json:"hostname"`
	IP       string     `orm:"column(ip);size(4096);" json:"ip"`
	OS       string     `orm:"column(os);size(64);" json:"os"`
	Arch     string     `orm:"column(arch);size(64);" json:"arch"`
	CPU      int        `orm:"column(cpu);" json:"cpu"`
	RAM      int64      `orm:"column(ram);" json:"ram"` // MB
	Disk     string     `orm:"column(disk);size(4096);" json:"disk"`
	BootTime *time.Time `orm:"column(boot_time);null;" json:"boottime"`
	Time     *time.Time `orm:"column(time);null;" json:"time"`

	HeartbeatTime *time.Time `orm:"column(heartbeat_time);null;" json:"heartbeat_time"`
	CreatedTime   *time.Time `orm:"column(created_time);auto_now_add;" json:"created_time"`
	DeletedTime   *time.Time `orm:"column(deleted_time);null;" json:"deleted_time"`

	IsOnline bool `orm:"-" json:"is_online"`
	IPList []string `orm:"-" json:"ip_list"`
	Disks map[string]float64 `orm:"-" json:"disks"`
}

func (a *Agent) Patch() {
	if time.Since(*a.HeartbeatTime) < 5 * time.Minute {
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

func (m *AgentManager) CreateOrReplace(agent *Agent) (*Agent, bool, error) {
	now := time.Now()
	ormer := orm.NewOrm()
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
	orm.NewOrm().QueryTable(&Agent{}).Filter("UUID__exact", uuid).Update(orm.Params{"HeartbeatTime": time.Now(), "DeletedTime": nil})
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
		query = query.Or("os__icontains", q)
		query = query.Or("arch__icontains", q)
		condition = condition.AndCond(query)

		qtotal, _ = queryset.SetCond(condition).Count()
	}
	var result []*Agent

	queryset.SetCond(condition).RelatedSel().Limit(length).Offset(start).All(&result)
	for _, agent := range result {
		agent.Patch()
	}
	return result, total, qtotal
}

func (m *AgentManager) DeleteById(id int) error {
	orm.NewOrm().QueryTable(&Agent{}).Filter("Id__exact", id).Update(orm.Params{"DeletedTime": time.Now()})
	return nil
}

func (m *AgentManager) GetStat() (int64, int64) {
	now := time.Now()
	onlineTime := now.Add(-5 * time.Minute)

	queryset := orm.NewOrm().QueryTable(new(Agent)).Filter("deleted_time__isnull", true)

	onlineCnt, _ := queryset.Filter("heartbeat_time__gte", onlineTime).Count()
	offlineCnt, _ :=queryset.Filter("heartbeat_time__lt", onlineTime).Count()
	return onlineCnt, offlineCnt
}

var DefaultAgentManager = NewAgentManager()


func init() {
	orm.RegisterModel(new(Agent))
}
