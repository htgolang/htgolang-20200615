package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/imsilence/gocmdb/server/cloud"
)

type CloudPlatform struct {
	Id          int        `orm:"column(id);" json:"id"`
	Name        string     `orm:"column(name);size(64);" json:"name"`
	Type        string     `orm:"column(type);size(32);" json:"type"`
	Addr        string     `orm:"column(addr);size(1024);" json:"addr"`
	AccessKey   string     `orm:"column(accesskey);size(1024);" json:"-"`
	SecrectKey  string     `orm:"column(secrectkey);size(1024);" json:"-"`
	Region      string     `orm:"column(region);size(64);" json:"region"`
	Remark      string     `orm:"column(remark);size(1024);" json:"remark"`
	CreatedTime *time.Time `orm:"column(created_time);type(datetime); auto_now_add;" json:"created_time"`
	DeletedTime *time.Time `orm:"column(deleted_time);type(datetime);null;default(null);" json:"deleted_time"`
	SyncTime    *time.Time `orm:"column(sync_time);auto_now;type(datetime);null;" json:"sync_time"`
	Status      int        `orm:"column(status);" json:"status"`
	Msg         string     `orm:"column(msg);size(1024)" json:"msg"`

	User            *User             `orm:"column(user);rel(fk);" json:"user"`
	VirtualMachines []*VirtualMachine `orm:"reverse(many);" json:"virtual_machines"`
}

type CloudPlatformManager struct {
}

func (m *CloudPlatformManager) Query(q string, start int64, length int) ([]*CloudPlatform, int64, int64) {
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

func (m *CloudPlatformManager) GetById(id int) *CloudPlatform {
	result := &CloudPlatform{}
	ormer := orm.NewOrm()
	err := ormer.QueryTable(&CloudPlatform{}).Filter("id__exact", id).Filter("deleted_time__isnull", true).One(result)
	if err == nil {
		ormer.LoadRelated(result, "User")
		return result
	} else {
		fmt.Println(err)
	}
	return nil

}

func (m *CloudPlatformManager) GetByName(name string) *CloudPlatform {
	ormer := orm.NewOrm()
	//var result CloudPlatform
	result := &CloudPlatform{}
	err := ormer.QueryTable(&CloudPlatform{}).Filter("deleted_time__isnull", true).Filter("name__exact", name).One(result)
	if err == nil {
		return result
	} else {
		return nil
	}
}

func (m *CloudPlatformManager) Create(name, typ, addr, region, accessKey, secrectKey, remark string, user *User) (*CloudPlatform, error) {
	ormer := orm.NewOrm()
	result := &CloudPlatform{
		Name:       name,
		Type:       typ,
		Addr:       addr,
		Region:     region,
		AccessKey:  accessKey,
		SecrectKey: secrectKey,
		Remark:     remark,
		User:       user,
		Status:     0,
	}
	if _, err := ormer.Insert(result); err != nil {
		return nil, err
	}
	return result, nil
}

func (m *CloudPlatformManager) Modify(id int, name, typ, addr, region, accessKey, secrectKey, remark string) (*CloudPlatform, error) {
	ormer := orm.NewOrm()
	if result := m.GetById(id); result != nil {
		result.Name = name
		result.Type = typ
		result.Addr = addr
		result.Region = region
		if accessKey != "" {
			result.AccessKey = accessKey
		}
		if secrectKey != "" {
			result.SecrectKey = secrectKey
		}
		// result.AccessKey = accessKey
		// result.SecrectKey = secrectKey
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

func (m *CloudPlatformManager) DeleteById(id int) error {
	// ormer := orm.NewOrm()
	// obj := &CloudPlatform{Id: id}
	// if err := ormer.Read(obj); err == nil {
	// 	now := time.Now()
	// 	obj.DeleteTime = &now
	// 	ormer.Update(obj)
	// }
	DefaultVirtualMachineManager.DeleteByPlatformId(id)
	orm.NewOrm().QueryTable(&CloudPlatform{}).Filter("Id__exact", id).Update(orm.Params{"deleted_time": time.Now()})
	return nil
}

func (m *VirtualMachineManager) DeleteByPlatformId(platform int) error {
	orm.NewOrm().QueryTable(&VirtualMachine{}).Filter("Platform__exact", platform).Update(orm.Params{"DeletedTime": time.Now()})
	return nil
}

func (m *CloudPlatformManager) SyncInfo(platform *CloudPlatform, now time.Time, msg string) error {
	platform.SyncTime = &now
	platform.Msg = msg
	_, err := orm.NewOrm().Update(platform)
	return err
}

func (m *CloudPlatformManager) DisableById(id int) error {
	orm.NewOrm().QueryTable(&CloudPlatform{}).Filter("Id__exact", id).Update(orm.Params{"status": 1})
	return nil
}

func (m *CloudPlatformManager) EnableById(id int) error {
	orm.NewOrm().QueryTable(&CloudPlatform{}).Filter("Id__exact", id).Update(orm.Params{"status": 0})
	return nil
}

type VirtualMachine struct {
	Id            int            `orm:"column(id)" json:"id"`
	Platform      *CloudPlatform `orm:"column(platform); rel(fk);" json:"platform"`
	UUID          string         `orm:"column(uuid);size(128)" json:"uuid"`
	Name          string         `orm:"column(name);size(64)" json:"name"`
	CPU           int            `orm:"column(cpu);" json:"cpu"`
	Mem           int64          `orm:"column(mem);" json:"mem"`
	OS            string         `orm:"column(os);size(128)" json:"os"`
	PrivateAddrs  string         `orm:"column(private_addrs);size(1024)" json:"private_addrs"`
	PublicAddrs   string         `orm:"column(public_addrs);size(1024)" json:"public_addrs"`
	Status        string         `orm:"column(status); size(32)" json:"status"`
	VmCreatedTime string         `orm:"column(vm_created_time);" json:"vm_created_time"`
	VmExpiredTime string         `orm:"column(vm_expired_time);" json:"vm_expired_time"`
	CreatedTime   *time.Time     `orm:"column(created_time);type(datetime); auto_now_add;" json:"created_time"`
	DeletedTime   *time.Time     `orm:"column(deleted_time);type(datetime);null;" json:"deleted_time"`
	UpdatedTime   *time.Time     `orm:"column(updated_time);type(datetime); auto_now;" json:"updated_time"`
}

func (p *CloudPlatform) IsEnable() bool {
	return p.Status == 0
}

type VirtualMachineManager struct{}

func NewVirtualMachineManager() *VirtualMachineManager {
	return &VirtualMachineManager{}
}

func (m *VirtualMachineManager) Query(q string, platform int, start int64, length int) ([]*VirtualMachine, int64, int64) {
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&VirtualMachine{})

	condition := orm.NewCondition()
	condition = condition.And("deleted_time__isnull", true)

	total, _ := queryset.SetCond(condition).Count()

	if q != "" {
		query := orm.NewCondition()
		query = query.Or("name__icontains", q)
		query = query.Or("public_addrs__icontains", q)
		query = query.Or("private_addrs__icontains", q)
		// query = query.Or("platform__icontains", q)
		query = query.Or("os__icontains", q)
		condition = condition.AndCond(query)
	}
	if platform > 0 {
		condition = condition.And("platform__exact", platform)
	}

	var result []*VirtualMachine
	qtotal, _ := queryset.SetCond(condition).Count()
	queryset.SetCond(condition).RelatedSel().Limit(length).Offset(start).All(&result)
	return result, total, qtotal

}

func (m *VirtualMachineManager) SyncInstance(instance *cloud.Instance, platform *CloudPlatform) {
	ormer := orm.NewOrm()
	// 更新

	// 存在的依据是什么
	vm := &VirtualMachine{UUID: instance.UUID, Platform: platform}
	if _, _, err := ormer.ReadOrCreate(vm, "UUID", "Platform"); err != nil {
		fmt.Println(err)
		return
	}

	// 存在
	vm.Name = instance.Name
	vm.OS = instance.OS
	vm.CPU = instance.CPU
	vm.Mem = instance.Mem
	vm.Status = instance.Status
	vm.VmCreatedTime = instance.CreatedTime
	vm.VmExpiredTime = instance.ExpiredTime
	vm.PublicAddrs = strings.Join(instance.PublicAddrs, ",")
	vm.PrivateAddrs = strings.Join(instance.PrivateAddrs, ",")
	ormer.Update(vm)

	// 如果存在  更新
	// 如果不存在 => 创建 => 创建(uuid, platform) + 更新
}

func (m *VirtualMachineManager) SyncInstanceStatus(now time.Time, platform *CloudPlatform) {
	orm.NewOrm().QueryTable(&VirtualMachine{}).Filter("Platform__exact", platform).Filter("UpdatedTime__lt", now).Update(orm.Params{"DeletedTime": now})
	orm.NewOrm().QueryTable(&VirtualMachine{}).Filter("Platform__exact", platform).Filter("UpdatedTime__gte", now).Update(orm.Params{"DeletedTime": nil})
}

func (m *VirtualMachineManager) GetById(id int) *VirtualMachine {
	vm := &VirtualMachine{}
	if orm.NewOrm().QueryTable(new(VirtualMachine)).RelatedSel().Filter("id__exact", id).Filter("DeletedTime__isnull", true).One(vm) == nil {
		return vm
	}
	return nil
}

var DefaultCloudPlatformManager = NewCloudPlatformManager()
var DefaultVirtualMachineManager = NewVirtualMachineManager()

func init() {
	orm.RegisterModel(&CloudPlatform{}, new(VirtualMachine))
}
