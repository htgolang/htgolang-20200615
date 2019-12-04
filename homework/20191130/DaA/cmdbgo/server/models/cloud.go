package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)


//云平台模型
type CloudPlatform struct {
	Id          int			`orm:"column(id);" json:"id"`
	Name        string		`orm:"column(name);size(64);" json:"name"`
	Type    	string		`orm:"column(type);size(32);" json:"type"`
	Addr		string		`orm:"column(addr);size(1024);" json:"addr"`
	Region 		string		`orm:"column(region);size(64);" json:"region"`
	AccessKey   string		`orm:"column(access_key);size(1024);null;" json:"-"`
	SecrectKey  string		`orm:"column(secret_key);size(1024);null;" json:"-"`
	Remark      string		`orm:"column(remark);size(1024);null;" json:"remark"`

	CreatedTime *time.Time	`orm:"column(created_time);auto_now_add;" json:"created_time"`
	UpdatedTime *time.Time	`orm:"column(updated_time);auto_now;" json:"updated_time"`
	DeletedTime *time.Time	`orm:"column(deleted_time);null;" json:"-"`

	SyncTime	*time.Time	`orm:"column(sync_time);null;" json:"sync_time"`
	Status      int			`orm:"column(status);" json:"status"`

	User		*User		`orm:"column(user);rel(fk);null;" json:"user"`
	VirtualMachines []*VirtualMachine `orm:"reverse(many);" json:"virtual_machines"`
}
func (p *CloudPlatform) IsEnable() bool {
	return p.Status == 0
}
//云平台模型管理器
type CloudPlatformManager struct {}
func NewCloudPlatformManager() *CloudPlatformManager {
	return &CloudPlatformManager{}
}
func (m *CloudPlatformManager) GetByName(name string) *CloudPlatform {
	ormer := orm.NewOrm()
	result := &CloudPlatform{}

	err := ormer.QueryTable(&CloudPlatform{}).Filter("deleted_time__isnull",true).Filter("name__iexact",name).One(result)
	if err == nil {
		return result
	}
	return nil

}		//通过name获取云平台模型数据
func (m *CloudPlatformManager) GetById(id int) *CloudPlatform {
	ormer := orm.NewOrm()
	result := &CloudPlatform{}

	err := ormer.QueryTable(&CloudPlatform{}).Filter("deleted_time__isnull",true).Filter("id__exact",id).One(result)
	if err == nil {
		return result
	}
	return nil
}				//通过ID获取云平台模型数据
func (m *CloudPlatformManager) SetStatusById(id int, status int) error {
	_,err := orm.NewOrm().QueryTable(&CloudPlatform{}).Filter("id__exact", id).Update(orm.Params{"status": status})
	return err
}	//通过ID设置云平台模型状态
func (m *CloudPlatformManager) DeleteById(id int) error {
	orm.NewOrm().QueryTable(&CloudPlatform{}).Filter("id__exact",id).Update(orm.Params{"deleted_time":time.Now()})
	return nil
}					//通过ID删除云平台模型数据
func (m *CloudPlatformManager) Query(q string, start int64, length int) ([]*CloudPlatform , int64, int64) {
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&CloudPlatform{})

	condition := orm.NewCondition()
	condition = condition.And("deleted_time__isnull",true)

	total,_ := queryset.SetCond(condition).Count()

	qtotal := total
	if q != "" {
		query := orm.NewCondition()
		query = query.Or("name__icontains",q)
		query = query.Or("addr__icontains",q)
		query = query.Or("remark__icontains",q)
		query = query.Or("region__icontains",q)
		condition = condition.AndCond(query)

		qtotal, _ = queryset.SetCond(condition).Count()
	}
	var result []*CloudPlatform

	queryset.SetCond(condition).Limit(length).Offset(start).All(&result)
	return result,total,qtotal
}							//查询
func (m *CloudPlatformManager) Create(name,typ,addr,region,accessKey,secrectKey,remark string,user *User) (*CloudPlatform,error) {
	ormer := orm.NewOrm()
	result := &CloudPlatform{
		Name:name,
		Type:typ,
		Addr:addr,
		Region:region,
		AccessKey:accessKey,
		SecrectKey:secrectKey,
		Remark:remark,
		User:user,
	}

	if _,err := ormer.Insert(result);err != nil {
		return nil,err
	}
	return result,nil
}	//创建
func (m *CloudPlatformManager) Modify(id int,name,typ,addr,region,accessKey,secrectKey,remark string,user *User) (*CloudPlatform,error) {
	ormer := orm.NewOrm()
	cloudPlatform := m.GetById(id)

	if cloudPlatform != nil  {
		if  accessKey == "" && secrectKey == "" {
			cloudPlatform.Name = name
			cloudPlatform.Type = typ
			cloudPlatform.Addr = addr
			cloudPlatform.Region = region
			cloudPlatform.Remark = remark
			if _,err := ormer.Update(cloudPlatform); err != nil {
				return nil,err
			}
			fmt.Println("空")
			return cloudPlatform,nil
		} else {
			cloudPlatform.Name = name
			cloudPlatform.Type = typ
			cloudPlatform.Addr = addr
			cloudPlatform.Region = region
			cloudPlatform.AccessKey = accessKey
			cloudPlatform.SecrectKey = secrectKey
			cloudPlatform.Remark = remark
			fmt.Println("不为空")
			if _,err := ormer.Update(cloudPlatform);err != nil {
				return nil,err
			}
			return cloudPlatform,nil
		}
	}

	fmt.Println(cloudPlatform)
	fmt.Println("对象不存在")
	return nil, fmt.Errorf("操作对象不存在或者")
}	//修改，还没写


//云主机模型
type VirtualMachine struct {
	Id				int				`orm:"column(id);" json:"id"`
	Platform		*CloudPlatform	`orm:"column(platform);rel(fk);" json:"platform"`
	UUID			string			`orm:"column(uuid);size(128);" json:"uuid"`
	Name			string			`orm:"column(name);size(64);" json:"name"`
	CPU				int				`orm:"column(cpu);" json:"cpu"`
	Mem				int64			`orm:"column(mem);" json:"mem"`
	OS				string			`orm:"column(os);size(128);" json:"os"`
	PrivateAddrs	string			`orm:"column(private_addrs);size(1024);" json:"private_addrs"`
	PublicAddrs		string			`orm:"column(public_addrs);size(1024);" json:"public_addrs"`
	Status			string			`orm:"column(status);size(32);" json:"status"`
	VmCreatedTime	string			`orm:"column(vm_created_time);" json:"vm_created_time"`
	VmExpiredTime	string			`orm:"column(vm_expired_time);" json:"vm_expired_time"`
	CreatedTime		*time.Time		`orm:"column(created_time);auto_now_add;" json:"created_time"`
	UpdatedTime		*time.Time		`orm:"column(updated_time);auto_now;" json:"updated_time"`
	DeletedTime		*time.Time		`orm:"column(deleted_time);null;" json:"-"`
}
//云主机模型管理器
type VirtualMachineManager struct {}
func NewVirtualMachineManager() *VirtualMachineManager {
	return &VirtualMachineManager{}
}					//初始化云主机模型管理器对象
func (m *VirtualMachineManager) Query(q string, start int64, length int) ([]*VirtualMachine , int64, int64) {
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&VirtualMachine{})

	condition := orm.NewCondition()
	condition = condition.And("deleted_time__isnull",true)

	total,_ := queryset.SetCond(condition).Count()

	qtotal := total
	if q != "" {
		query := orm.NewCondition()
		query = query.Or("name__icontains",q)
		query = query.Or("public_addrs__icontains",q)
		query = query.Or("private_addrs__icontains",q)
		query = query.Or("os__icontains",q)
		condition = condition.AndCond(query)

		qtotal, _ = queryset.SetCond(condition).Count()
	}
	var result []*VirtualMachine

	queryset.SetCond(condition).Limit(length).Offset(start).All(&result)
	return result,total,qtotal
}	//查询
func (m *VirtualMachineManager) Create() {} 	//创建，还没写，因为这个数据获取与以往form获取不同，需要从定时器获取，是不是应该用新的获取逻辑去做？
func (m *VirtualMachineManager) Modify() {}		//修改，还没写，同上，创建和修改以及删除的逻辑，是不是应该放在自动获取里，重新组织关系。
func (m *VirtualMachineManager) Delete() {}		//删除，还没写，同上。


var DefaultCloudPlatformManager = NewCloudPlatformManager()
var DefaultVirtualMachineManager = NewVirtualMachineManager()

func init(){
	orm.RegisterModel(&CloudPlatform{},&VirtualMachine{})
}

