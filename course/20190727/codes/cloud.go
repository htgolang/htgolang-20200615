package main

import "fmt"

type Host struct {
	Id   string
	Name string
}

type Cloud interface {
	GetList() []Host
	Start(Id string) error
	Stop(Id string) error
	Detail(Id string) Host
}

type Tenant struct {
	API    string
	key    string
	secret string
}

func NewTenant(api, key, secret string) Tenant {
	return Tenant{api, key, secret}
}

func (t Tenant) GetList() []Host {
	fmt.Println("Tenant.GetList")
	return []Host{}
}

func (t Tenant) Start(Id string) error {
	fmt.Println("Tenant.Start")
	return nil
}

func (t Tenant) Stop(Id string) error {
	fmt.Println("Tenant.Stop")
	return nil
}

func (t Tenant) Detail(Id string) Host {
	fmt.Println("Tenant.Detail")
	return Host{}
}

type Aliyun struct {
	API    string
	key    string
	secret string
}

func NewAliyun(api, key, secret string) Aliyun {
	return Aliyun{api, key, secret}
}

func (t Aliyun) GetList() []Host {
	fmt.Println("Aliyun.GetList")
	return []Host{}
}

func (t Aliyun) Start(Id string) error {
	fmt.Println("Aliyun.Start")
	return nil
}

func (t Aliyun) Stop(Id string) error {
	fmt.Println("Aliyun.Stop")
	return nil
}

func (t Aliyun) Detail(Id string) Host {
	fmt.Println("Aliyun.Detail")
	return Host{}
}

func NewCloud(cloudType string) Cloud {
	var cloud Cloud
	if cloudType == "tenant" {
		cloud = NewTenant("tenant", "tenant", "tenant")
	} else if cloudType == "aliyun" {
		cloud = NewAliyun("aliyun", "aliyun", "aliyun")
	} else {
		fmt.Println("error")
	}
	return cloud
}

func main() {

	var cloudType string // tenant/aliyun

	cloudType = "tenant"

	cloud := NewCloud(cloudType)
	if cloud == nil {
		return
	}

	var operate string
EOF:
	for {
		fmt.Print("请输入操作:")
		fmt.Scan(&operate)
		switch operate {
		case "1":
			cloud.GetList()
			fmt.Println("同步主机列表")
			fmt.Println("插入数据库")
			fmt.Println("通知新添加主机")
		case "2":
			cloud.Start("1")
			fmt.Println("启动主机")
		case "3":
			cloud.Stop("2")
			fmt.Println("停止主机")
		case "4":
			cloud.Detail("3")
			fmt.Println("获取主机详情")
		case "5":
			break EOF
		}
	}
}
