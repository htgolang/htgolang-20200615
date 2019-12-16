package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Agent struct {
	Id 			int 	`orm:"column(id);" json:"id"`
	UUID 		string 	`orm:"column(uuid);size(64);" json:"uuid"`
	HostName 	string 	`orm:"column(host_name);size(128);" json:"host_name"`
	IP 			string 	`orm:"column(ip);size(4096);" json:"ip"`
	OS 			string 	`orm:"column(os);size(64);" json:"os"`
	Arch 		string 	`orm:"column(arch);size(64);" json:"arch"`
	CPU 		int		`orm:"column(cpu);" json:"cpu"`
	RAM 		int64 	`orm:"column(ram);" json:"ram"` //MB
	Disk 		string 	`orm:"column(disk);size(4096);" json:"disk"`
	BootTime 	*time.Time `orm:"column(boot_time);null" json:"boot_time"`
	Time 		*time.Time `orm:"column(time);null;" json:"time"`

	HeartbeatTime 	*time.Time 	`orm:"column(heartbeat_time);null;" json:"heartbeat_time"`
	CreatedTime 	*time.Time 	`orm:"column(created_time);auto_now_add;null;" json:"created_time"`
	DeletedTime 	*time.Time 	`orm:"column(deleted_time);null;"json:"deleted_time"`

	//json用信息
	IsOnline 	bool `orm:"-" json:"is_online"`

}

type AgentManager struct {

}

func NewAgentManager() *AgentManager{
	return &AgentManager{}
}

func (m *AgentManager) CreateOrReplace(agent *Agent) (*Agent, bool, error){
	ormer := orm.NewOrm()
	orgAgent := &Agent{UUID: agent.UUID}
	if created, _, err := ormer.ReadOrCreate(orgAgent, "UUID"); err != nil{
		return nil, false, err
	} else {
		now := time.Now()
		orgAgent.HostName = agent.HostName
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
	ormer := orm.NewOrm()
	ormer.QueryTable(&Agent{}).Filter("UUID__exact", uuid).Update(orm.Params{"HeartbeatTime":time.Now()})
}


var DefaultAgentManager = NewAgentManager()

func init(){
	orm.RegisterModel(&Agent{})
}