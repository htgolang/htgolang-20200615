package models

import (
	"github.com/astaxie/beego/orm"
	"time"
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

	IsOnline bool `orm:"-" json:"is_onlne"`
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
	orm.NewOrm().QueryTable(&Agent{}).Filter("UUID__exact", uuid).Update(orm.Params{"HeartbeatTime": time.Now()})
}

var DefaultAgentManager = NewAgentManager()


func init() {
	orm.RegisterModel(new(Agent))
}
