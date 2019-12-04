package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type CloudPlatform struct {
	Id          int        `orm:"column(id);" json:"id"`
	Name        string     `orm:"column(name);size(64);" json:"name"`
	Type    string     `orm:"column(type);size(32);" json:"type"`
	Addr      string        `orm:"column(addr);size(1024);" json:"addr"`
	AccessKey    string `orm:"column(access_key);size(1024);" json:"-"`
	SecrectKey         string    `orm:"column(secrect_key);size(1024);" json:"-"`
	Region       string     `orm:"column(region);size(64);" json:"region"`
	Remark      string     `orm:"column(remark);size(1024);" json:"remark"`
	Status      int        `orm:"column(status);" json:"status"`
	CreatedTime *time.Time `orm:"column(created_time); type(datetime);auto_now_add;" json:"created_time"`
	SyncTime *time.Time `orm:"column(sync_time); type(datetime);null;" json:"sync_time"`
	DeletedTime *time.Time `orm:"column(deleted_time);type(datetime);null;" json:"deleted_time"`

	User *User `orm:"column(user);rel(fk);" json:"user"`

	VirtualMachine []*VirtualMachine `orm:"reverse(many);" json:"virtual_machine"`
}

func (p *CloudPlatform) IsEnable() bool {
	return p.Status == 0
}

type CloudPlatformManager struct {

}

func (m *CloudPlatformManager) Query(q string, start int64, length int) ([] *CloudPlatform, int64, int64){
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&CloudPlatform{})

	condition := orm.NewCondition()

	condition = condition.And("deleted_time__isnull", true)

	total, _ := queryset.SetCond(condition).Count()
	qtotal := total

	if q != "" {
		query := orm.NewCondition()
		query = query.Or("name__icontain", q)
		query = query.Or("addr__icontain", q)
		query = query.Or("remark__icontain", q)
		query = query.Or("name__icontain", q)
		query = query.Or("region__icontain", q)
		condition = condition.AndCond(query)

		qtotal, _ = queryset.SetCond(condition).Count()

	}

	var result []*CloudPlatform
	queryset.SetCond(condition).Limit(length).Offset(start).All(&result)

	return result, total, qtotal


}

func NewCloudPlatformManager() *CloudPlatformManager {
	return &CloudPlatformManager{}
}
func (m *CloudPlatformManager) GetById(pk int) *CloudPlatform {
	result := CloudPlatform{}
	err := orm.NewOrm().QueryTable(&CloudPlatform{}).Filter("id__exact", pk).Filter("DeletedTime__isnull", true).One(&result)
	if err == nil {
		return &result
	}

	return nil
}
//
func (m *CloudPlatformManager) GetByName(name string) *CloudPlatform {
	result := CloudPlatform{}
	err := orm.NewOrm().QueryTable(&CloudPlatform{}).Filter("name__exact", name).Filter("DeletedTime__isnull", true).One(&result)
	if err == nil {
		return &result
	}

	return nil
}
//

func (m *CloudPlatformManager) Create(name, typ, addr, region, accesskey, secrectkey, remark string, user *User) (*CloudPlatform, error) {

	ormer := orm.NewOrm()

	result := &CloudPlatform{
		Name:name,
		Type:typ,
		Addr:addr,
		Region:region,
		AccessKey:accesskey,
		SecrectKey:accesskey,
		Remark:remark,
		User:user,
		Status: 0,
	}
	if _, err := ormer.Insert(result); err != nil {

		return nil, err
	}

	return result, nil
}

func (m *CloudPlatformManager) DeleteById(pk int) error {

	orm.NewOrm().QueryTable(&CloudPlatform{}).Filter("id__exact", pk).Update(orm.Params{"deleted_time": time.Now()})
	return nil
}

func (m *CloudPlatformManager) SetStatusById(pk int, status int) error {

	orm.NewOrm().QueryTable(&CloudPlatform{}).Filter("id__exact", pk).Update(orm.Params{"status": status})
	return nil
}

func (m *CloudPlatformManager) Modify(id int, name, typ, addr, region, accesskey, secrectkey, remark string) (*CloudPlatform, error) {
	ormer := orm.NewOrm()
	if c := m.GetById(id); c != nil {
		c.Name = name
		c.Type = typ
		c.Addr = addr
		c.Region = region
		c.AccessKey = accesskey
		c.SecrectKey = secrectkey
		c.Remark = remark
		if _, err := ormer.Update(c); err != nil {
			return nil, err
		}
		return c, nil
	}

	return nil, fmt.Errorf("操作对象不存在")
}

type VirtualMachine struct {
	Id          int        `orm:"column(id);" json:"id"`
	Platform 	*CloudPlatform `orm:"column(platform); rel(fk);" json:"platform"`
	UUID 		string 	`orm:"column(uuid);size(128);" json:"uuid"`
	Name        string     `orm:"column(name);size(64);" json:"name"`
	CPU    int     `orm:"column(cpu);" json:"cpu"`
	Mem      int64        `orm:"column(mem);" json:"mem"`
	OS    string `orm:"column(os);size(1024);" json:"os"`
	PrivateAddrs         string    `orm:"column(private_addrs);size(1024);" json:"private_addrs"`
	PublicAddrs       string     `orm:"column(public_addrs);size(1024);" json:"public_addrs"`
	Status      int        `orm:"column(status);size(32);" json:"status"`
	VmCreatedTime string `orm:"column(vm_created_time)"; json:"vm_created_time"`
	VmExpiredTime string `orm:"column(vm_expired_time)"; json:"vm_expired_time"`

	CreatedTime *time.Time `orm:"column(created_time); type(datetime);auto_now_add;" json:"created_time"`
	UpdatedTime *time.Time `orm:"column(updated_time); type(datetime);auto_now;" json:"sync_time"`
	DeletedTime *time.Time `orm:"column(deleted_time);type(datetime);null;" json:"deleted_time"`

	//User *User `orm:"column(user);rel(fk);" json:"user"`
}
type VirtualMachineManager struct {}

func (m *VirtualMachineManager) Query(q string, start int64, length int) ([] *VirtualMachine, int64, int64){
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&VirtualMachine{})

	condition := orm.NewCondition()

	condition = condition.And("deleted_time__isnull", true)

	total, _ := queryset.SetCond(condition).Count()
	qtotal := total

	if q != "" {
		query := orm.NewCondition()
		query = query.Or("name__icontain", q)
		query = query.Or("addr__icontain", q)
		query = query.Or("remark__icontain", q)
		query = query.Or("name__icontain", q)
		query = query.Or("region__icontain", q)
		condition = condition.AndCond(query)

		qtotal, _ = queryset.SetCond(condition).Count()

	}

	var result []*VirtualMachine
	queryset.SetCond(condition).Limit(length).Offset(start).All(&result)


	return result, total, qtotal


}

func NewVirtualMachineManager() *VirtualMachineManager {
	return &VirtualMachineManager{}
}

//
func (m *VirtualMachineManager) GetByName(name string) *VirtualMachine {
	result := VirtualMachine{}
	err := orm.NewOrm().QueryTable(&VirtualMachine{}).Filter("name__exact", name).Filter("DeletedTime__isnull", true).One(&result)
	if err == nil {
		return &result
	}

	return nil
}
//

func (m *VirtualMachineManager) DeleteById(pk int) error {
	fmt.Println(pk)
	orm.NewOrm().QueryTable(&CloudPlatform{}).Filter("id__exact", pk).Update(orm.Params{"deleted_time": time.Now()})
	return nil
}




var DefaultCloudPlatformManager = NewCloudPlatformManager()
var DefaultVirtualMachineManager = NewVirtualMachineManager()
//
func init() {
	orm.RegisterModel(new(CloudPlatform), new(VirtualMachine))
}
