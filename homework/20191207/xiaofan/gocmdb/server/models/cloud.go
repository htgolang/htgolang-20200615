package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/dcosapp/gocmdb/server/cloud"
	"strings"
	"time"
)

// 云平台model
type CloudPlatform struct {
	Id          int        `orm:"column(id);"json:"id"`
	Name        string     `orm:"column(name);size(64);"json:"name"`
	Type        string     `orm:"column(type);size(32);"json:"type"`
	Addr        string     `orm:"column(addr);size(1024);"json:"addr"`
	AccessKey   string     `orm:"column(access_key);size(1024);"json:"-"`
	SecretKey   string     `orm:"column(secret_key);size(1024);"json:"-"`
	Region      string     `orm:"column(region);size(64);"json:"region"`
	Remark      string     `orm:"column(remark);size(1024);"json:"remark"`
	CreatedTime *time.Time `orm:"column(created_time);type(datetime);auto_now_add"json:"created_time"`
	SyncedTime  *time.Time `orm:"column(synced_time);type(datetime);null"json:"synced_time"`
	DeletedTime *time.Time `orm:"column(deleted_time);type(datetime);null;default(null);"json:"-"`
	User        *User      `orm:"column(user);rel(fk);"json:"user"`
	Status      int        `orm:"column(status);"json:"status"`
	Msg         string     `orm:"column(msg);size(1024);"json:"msg"`

	VirtualMachines []*VirtualMachine `orm:"reverse(many);"json:"virtual_machines"`
}

func (p *CloudPlatform) IsEnable() bool {
	return p.Status == 0
}

type CloudPlatformManager struct {
}

// 查询云平台
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
	_, _ = queryset.SetCond(condition).RelatedSel().Limit(length).Offset(start).All(&result)
	return result, total, qtotal
}

// 根据名称获取云平台
func (m *CloudPlatformManager) GetByName(name string) *CloudPlatform {
	cloud := &CloudPlatform{Name: name, DeletedTime: nil}
	err := orm.NewOrm().QueryTable(&CloudPlatform{}).Filter("Name__exact", name).Filter("deleted_time__isnull", true).One(cloud)
	if err == nil {
		return cloud
	}
	return nil
}

// 更改同步时间与同步信息 (成功||失败)
func (m *CloudPlatformManager) SyncInfo(platform *CloudPlatform, now time.Time, msg string) error {
	platform.SyncedTime = &now
	platform.Msg = msg
	_, err := orm.NewOrm().Update(platform)
	return err
}

// 根据id获取云平台
func (m *CloudPlatformManager) GetById(id int) *CloudPlatform {
	cloud := &CloudPlatform{Id: id, DeletedTime: nil}
	err := orm.NewOrm().QueryTable(&CloudPlatform{}).Filter("Id__exact", id).Filter("deleted_time__isnull", true).One(cloud)
	if err == nil {
		return cloud
	}
	return nil
}

// 创建云平台
func (m *CloudPlatformManager) Create(name, typ, addr, accessKey, secretKey, region, remark string, user *User) (*CloudPlatform, error) {
	ormer := orm.NewOrm()
	result := &CloudPlatform{
		Name:      name,
		Type:      typ,
		Addr:      addr,
		AccessKey: accessKey,
		SecretKey: secretKey,
		Region:    region,
		Remark:    remark,
		User:      user,
	}
	if _, err := ormer.Insert(result); err != nil {
		return nil, err
	}
	return result, nil
}

// 修改云平台
func (m *CloudPlatformManager) Modify(id int, name, typ, addr, region, accessKey, secretKey, remark string) (*CloudPlatform, error) {
	platform := &CloudPlatform{Id: id}
	ormer := orm.NewOrm()
	err := ormer.Read(platform)
	if err == nil {
		platform.Name = name
		platform.Type = typ
		platform.Addr = addr
		platform.Region = region
		platform.Remark = remark
		if accessKey != "" {
			platform.AccessKey = accessKey
		}
		if secretKey != "" {
			platform.SecretKey = secretKey
		}
		if _, err := ormer.Update(platform); err == nil {
			return platform, nil
		}
	}
	return &CloudPlatform{}, err
}

// 删除云平台(逻辑删除)
func (m *CloudPlatformManager) DeleteById(pk int) (int64, error) {
	now := time.Now()
	result, err := orm.NewOrm().QueryTable(&CloudPlatform{}).Filter("Id__exact", pk).Update(orm.Params{"DeletedTime": &now})
	return result, err
}

// 设置是否锁定(0 Enable, 1 Disable)
func (m *CloudPlatformManager) SetStatusById(pk, status int) error {
	_, err := orm.NewOrm().QueryTable(&CloudPlatform{}).Filter("id__exact", pk).Update(orm.Params{"status": status})
	return err
}

func NewCloudPlatformManager() *CloudPlatformManager {
	return &CloudPlatformManager{}
}

// 虚拟机model
type VirtualMachine struct {
	Id            int            `orm:"column(id);"json:"id"`
	Platform      *CloudPlatform `orm:"column(Platform);rel(fk);"json:"platform"`
	UUID          string         `orm:"column(uuid);size(128);"json:"uuid"`
	Name          string         `orm:"column(name);size(64);"json:"name"`
	CPU           int            `orm:"column(cpu);"json:"cpu"`
	Mem           int64          `orm:"column(mem);"json:"mem"`
	OS            string         `orm:"column(os);size(128);"json:"os"`
	PrivateAddrs  string         `orm:"column(private_addrs);size(1024);"json:"private_addrs"`
	PublicAddrs   string         `orm:"column(public_addrs);size(1024);"json:"public_addrs"`
	Status        string         `orm:"column(status);size(32);"json:"status"`
	VmCreatedTime string         `orm:"column(vm_created_time);"json:"vm_created_time"`
	VmExpiredTime string         `orm:"column(vm_expired_time);"json:"vm_expired_time"`

	CreatedTime *time.Time `orm:"column(created_time);auto_now_add;type(datetime);"json:"created_time"`
	DeletedTime *time.Time `orm:"column(deleted_time);null;type(datetime);"json:"deleted_time"`
	UpdatedTime *time.Time `orm:"column(updated_time);auto_now;type(datetime);"json:"updated_time"`
}

type VirtualMachineManager struct {
}

// 查询虚拟机
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
		query = query.Or("os__icontains", q)
		condition = condition.AndCond(query)
	}

	if platform > 0 {
		condition = condition.And("platform__exact", platform)
	}
	var result []*VirtualMachine

	qtotal, _ := queryset.SetCond(condition).Count()
	_, _ = queryset.SetCond(condition).RelatedSel().Limit(length).Offset(start).All(&result)
	return result, total, qtotal
}

// 根据名称获取虚拟机
func (m *VirtualMachineManager) GetByName(name string) *VirtualMachine {
	vm := &VirtualMachine{}
	err := orm.NewOrm().QueryTable(VirtualMachine{}).Filter("Name__exact", name).Filter("deleted_time__isnull", true).One(vm)
	if err == nil {
		return vm
	}
	return nil
}

// 根据ID获取虚拟机
func (m *VirtualMachineManager) GetById(id int) *VirtualMachine {
	vm := &VirtualMachine{}
	err := orm.NewOrm().QueryTable(VirtualMachine{}).RelatedSel().Filter("Id__exact", id).Filter("deleted_time__isnull", true).One(vm)
	if err == nil {
		return vm
	}
	return nil
}

// 当云平台被删除或禁用后，数据库中逻辑删除虚拟机
func (m *VirtualMachineManager) DeleteById(pk int) (int64, error) {
	now := time.Now()
	result, err := orm.NewOrm().QueryTable(&VirtualMachine{}).RelatedSel().Filter("Platform__exact", pk).Update(orm.Params{"DeletedTime": &now})
	return result, err
}

// 虚拟机信息写入数据库
func (m *VirtualMachineManager) SyncInstance(instance *cloud.Instance, platform *CloudPlatform) {
	ormer := orm.NewOrm()
	vm := &VirtualMachine{UUID: instance.UUID, Platform: platform}

	// 是否创建, id, err
	if _, _, err := ormer.ReadOrCreate(vm, "UUID", "Platform"); err != nil {
		return
	}
	vm.Name = instance.Name
	vm.OS = instance.OS
	vm.CPU = instance.CPU
	vm.Mem = instance.Memory
	vm.Status = instance.Status
	vm.VmCreatedTime = instance.CreatedTime
	vm.VmExpiredTime = instance.ExpiredTime
	vm.PublicAddrs = strings.Join(instance.PublicAddrs, ",")
	vm.PrivateAddrs = strings.Join(instance.PrivateAddrs, ",")
	_, _ = ormer.Update(vm)
}

// 虚拟机状态写入数据库
func (m *VirtualMachineManager) SyncInstacneStatus(now time.Time, platform *CloudPlatform) {
	orm.NewOrm().QueryTable(&VirtualMachine{}).Filter("Platform__exact", platform).Filter("UpdatedTime__lt",
		now).Update(orm.Params{"DeletedTime": now})
	orm.NewOrm().QueryTable(&VirtualMachine{}).Filter("Platform__exact", platform).Filter("UpdatedTime__gt",
		now).Update(orm.Params{"DeletedTime": nil})
}

func NewVirtualMachineManager() *VirtualMachineManager {
	return &VirtualMachineManager{}
}

var DefaultCloudPlatformManager = NewCloudPlatformManager()
var DefaultVirtualMachineManager = NewVirtualMachineManager()

func init() {
	orm.RegisterModel(&CloudPlatform{}, &VirtualMachine{})
}
