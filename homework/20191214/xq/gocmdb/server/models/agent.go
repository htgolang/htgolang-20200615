package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Agent struct {
	Id int `orm:"column(id)" json:"id"`
	UUID string `orm:"column(uuid);size(64)" json:"uuid"`
	Hostname string `orm:"column(hostname);size(64)" json:"hostname"`
	IP string `orm:"column(ip);size(4096)" json:"ip"`
	OS string `orm:"column(os);size(64)" json:"os"`
	Arch string `orm:"column(arch);size(64)" json:"arch"`
	CPU int `orm:"column(cpu)" json:"cpu"`
	MEM int64 `orm:"column(mem)" json:"mem"` // MB
	Disk string `orm:"column(disk);size(4096)" json:"disk"`
	BootTime *time.Time `orm:"column(boot_time); null" json:"boot_time"`
	Time *time.Time `orm:"column(time); null" json:"time"`
	CreatedTime *time.Time `orm:"column(created_time); auto_now_add" json:"created_time"`
	DeletedTime *time.Time `orm:"column(deleted_time); null" json:"deleted_time"`
	HeartbeatTime *time.Time `orm:"column(heartbeat_time);null" json:"heartbeat_time"`
	IsOnline bool `orm:"-" json:"is_online"`
}


type AgentManager struct {

}

func NewAgentManager() *AgentManager {
	return &AgentManager{}
}

func (m *AgentManager) CreateOrReplace(agent *Agent) (*Agent, bool, error) {

	now := time.Now()
	ormer := orm.NewOrm()

	orgAgent := &Agent{
		UUID:agent.UUID,
	}

	if created, _, err := ormer.ReadOrCreate(orgAgent, "UUID"); err != nil {
		return nil,false, err
	}else {


		orgAgent.Hostname = agent.Hostname
		orgAgent.IP = agent.IP
		orgAgent.OS = agent.OS
		orgAgent.CPU = agent.CPU
		orgAgent.MEM = agent.MEM
		orgAgent.Arch = agent.Arch
		orgAgent.Disk = agent.Disk
		orgAgent.BootTime = agent.BootTime
		agent.DeletedTime = nil
		orgAgent.Time = agent.Time
		orgAgent.HeartbeatTime = &now

		ormer.Update(orgAgent)

		return orgAgent, created, nil

	}

}

func (m *AgentManager) Heartbeat(uuid string) {
	ormer := orm.NewOrm()
	ormer.QueryTable(&Agent{}).Filter("UUID__exact", uuid).Update(orm.Params{"HeartbeatTime":time.Now()})
}


func (m *AgentManager) Query(q string, start int64, length int) ([]*Agent, int64, int64){
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

var DefaultAgentManager  = NewAgentManager()


func init(){
	orm.RegisterModel(new(Agent))
}