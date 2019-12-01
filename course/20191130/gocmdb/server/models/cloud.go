package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type CloudPlatform struct {
	Id int `orm:"column(id);" json:"id"`
    Name string `orm:"column(name);size(64);" json:"name"`
    Type string `orm:"column(type);size(32);" json:"type"`
    Addr string `orm:"column(addr);size(1024);" json:"addr"`
    AccessKey string `orm:"column(access_key);size(1024);" json:"-"`
    SecrectKey string `orm:"column(secrect_key);size(1024);" json:"-"`
    Region string `orm:"column(region);size(64);" json:"region"`
    Remark string `orm:"column(remark);size(1024);" json:"remark"`
    CreatedTime *time.Time `orm:"column(created_time);type(datetime);auto_now_add;" json:"created_time"`
    DeletedTime *time.Time `orm:"column(deleted_time);type(datetime);null;" json:"deleted_time"`
    SyncTime *time.Time `orm:"column(sync_time);type(datetime);null;" json:"sync_time"`
    User *User `orm:"column(user);rel(fk);" json:"user"`
	Status int `orm:"column(status);" json:"status"`

	VirtualMachines []*VirtualMachine `orm:"reverse(many);" json:"virtual_machines"`
}

func (p *CloudPlatform) IsEnable() bool {
	return p.Status == 0
}


type CloudPlatformManager struct {}

func (m *CloudPlatformManager) Query(q string, start int64, length int) ([] *CloudPlatform, int64, int64) {
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&CloudPlatform{})

	condition := orm.NewCondition()
	condition = condition.And("deleted_time__isnull", true)

	total, _ := queryset.SetCond(condition).Count()

	qtotal := total
	if q != "" {
		query := orm.NewCondition()
		query = query.Or("name__icontains", q)
		query = query.Or("addr__icontains", q)
		query = query.Or("remark__icontains", q)
		query = query.Or("region__icontains", q)
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

func (m *CloudPlatformManager) GetByName(name string) *CloudPlatform {
	ormer := orm.NewOrm()
	// var result CloudPlatform
	result := &CloudPlatform{}
	err := ormer.QueryTable(&CloudPlatform{}).Filter("deleted_time__isnull", true).Filter("name__exact", name).One(result)
	if err == nil {
		return result
	}
	return nil
}

func (m *CloudPlatformManager) Create(name, typ, addr, region, accessKey, secrectKey, remark string, user *User) (*CloudPlatform, error) {
	ormer := orm.NewOrm()
	result := &CloudPlatform {
		Name: name,
		Type: typ,
		Addr: addr,
		Region: region,
		AccessKey: accessKey,
		SecrectKey: secrectKey,
		Remark: remark,
		User: user,
		Status: 0,
	}
	if _, err := ormer.Insert(result); err != nil {
		return nil, err
	}
	return result, nil
}

func (m *CloudPlatformManager) DeleteById(id int) error {
	orm.NewOrm().QueryTable(&CloudPlatform{}).Filter("Id__exact", id).Update(orm.Params{"DeletedTime" : time.Now()})
	return nil
}


type VirtualMachine struct {
	Id int `orm:"column(id)" json:"id"`
    Platform *CloudPlatform `orm:"column(platform);rel(fk);" json:"platform"`
    UUID string `orm:"column(uuid);size(128);" json:"uuid"`
    Name string `orm:"column(name);size(64);" json:"name"`
    CPU int `orm:"column(cpu);" json:"cpu"`
    Mem int64 `orm:"column(mem);" json:"mem"`
    OS string `orm:"column(os);size(128);" json:"os"`
    PrivateAddrs string `orm:"column(private_addrs);size(1024);" json:"private_addrs"`
    PublicAddrs string `orm:"column(public_addrs);size(1024);" json:"public_addrs"`
    Status string `orm:"column(status);size(32);" json:"status"`
    VmCreatedTime string `orm:"column(vm_created_time);" json:"vm_created_time"`
    VmExpiredTime string `orm:"column(vm_expired_time);" json:"vm_expired_time"`

    CreatedTime *time.Time `orm:"column(created_time);auto_now_add;type(datetime);" json:"created_time"`
    DeletedTime *time.Time `orm:"column(deleted_time);type(datetime);null" json:"deleted_time"`
    UpdatedTime *time.Time `orm:"column(updated_time);auto_now;type(datetime);" json:"updated_time"`
}

type VirtualMachineManager struct {}

func NewVirtualMachineManager() *VirtualMachineManager {
	return &VirtualMachineManager{}
}

func (m *VirtualMachineManager) Query(q string, start int64, length int) ([] *VirtualMachine, int64, int64) {
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&VirtualMachine{})

	condition := orm.NewCondition()
	condition = condition.And("deleted_time__isnull", true)

	total, _ := queryset.SetCond(condition).Count()

	qtotal := total
	if q != "" {
		query := orm.NewCondition()
		query = query.Or("name__icontains", q)
		query = query.Or("public_addrs__icontains", q)
		query = query.Or("private_addrs__icontains", q)
		query = query.Or("os__icontains", q)
		condition = condition.AndCond(query)

		qtotal, _ = queryset.SetCond(condition).Count()
	}
	var result []*VirtualMachine

	queryset.SetCond(condition).Limit(length).Offset(start).All(&result)
	return result, total, qtotal
}


var DefaultCloudPlatformManager = NewCloudPlatformManager()
var DefaultVirtualMachineManager = NewVirtualMachineManager()

func init() {
	orm.RegisterModel(&CloudPlatform{}, new(VirtualMachine))
}