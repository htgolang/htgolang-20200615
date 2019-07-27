package main

import "fmt"

type Host struct {
	Id   string
	Name string
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

func main() {

	var cloudType string // tenant/aliyun

	cloudType = "aliyun"

	var tenant Tenant = NewTenant("tenant", "tenant", "tenant")
	var aliyun Aliyun = NewAliyun("aliyun", "aliyun", "aliyun")

	var operate string
EOF:
	for {
		fmt.Print("请输入操作:")
		fmt.Scan(&operate)
		switch operate {
		case "1":
			if cloudType == "tenant" {
				tenant.GetList()
			} else if cloudType == "aliyun" {
				aliyun.GetList()
			}
			fmt.Println("同步主机列表")
			fmt.Println("插入数据库")
			fmt.Println("通知新添加主机")
		case "2":
			if cloudType == "tenant" {
				tenant.Start("1")
			} else if cloudType == "aliyun" {
				aliyun.Start("1")
			}
			fmt.Println("启动主机")
		case "3":

			if cloudType == "tenant" {
				tenant.Stop("1")
			} else if cloudType == "aliyun" {
				aliyun.Stop("1")
			}
			fmt.Println("停止主机")
		case "4":
			if cloudType == "tenant" {
				tenant.Detail("1")
			} else if cloudType == "aliyun" {
				aliyun.Detail("1")
			}
			fmt.Println("获取主机详情")
		case "5":
			break EOF
		}
	}
}
