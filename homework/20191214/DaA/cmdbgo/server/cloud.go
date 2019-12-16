package main

import (
	"flag"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"github.com/xxdu521/cmdbgo/server/cloud"
	_ "github.com/xxdu521/cmdbgo/server/cloud/plugins"
	"github.com/xxdu521/cmdbgo/server/models"
	_ "github.com/xxdu521/cmdbgo/server/routers"
	"os"
	"time"
)

func main(){
	//初始化命令行参数，解析，初始化orm，测试数据库连接，
	h := flag.Bool("h",false,"help")
	help := flag.Bool("help",false,"help")
	verbose := flag.Bool("v",false,"verbose")

	flag.Usage = func() {
		fmt.Println("usage: web -h")
		flag.PrintDefaults()
	}

	// 解析命令行参数
	flag.Parse()

	if *h || *help {
		flag.Usage()
		os.Exit(0)
	}

	//fmt.Println(*init,*syncdb,*force,*verbose)
	//设置日志输出到文件
	beego.SetLogger("file",`{"filename":"logs/web.log","level":7}`,)

	if !*verbose {
		//删除控制台日志
		beego.BeeLogger.DelLogger("console")
	} else {
		//打开debug日志
		orm.Debug = true
	}

	// 初始化orm
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default","mysql",beego.AppConfig.String("dsn"))
	// create database gocmdb default charset utf8mb4;  自己创建数据库

	// 测试数据库连接是否正常
	if db,err := orm.GetDB(); err != nil || db.Ping() != nil {
		beego.Error("数据库连接失败")
		os.Exit(-1)
	}

	for now := range time.Tick(10 * time.Second) {
		fmt.Println(now)

		platforms, _, _ := models.DefaultCloudPlatformManager.Query("",0,0)
		for _,platform := range platforms {
			if !platform.IsEnable(){
				continue
			}
			//fmt.Println(platform)
			if sdk, ok := cloud.DefaultManager.Cloud(platform.Type); !ok {
				fmt.Println("云平台未注册")
			} else {
				sdk.Init(platform.Addr, platform.Region, platform.AccessKey, platform.SecrectKey)
				if err := sdk.TestConnect(); err != nil {
					fmt.Println("测试连接失败", err)
					models.DefaultCloudPlatformManager.SyncInfo(platform, now, fmt.Sprintf("测试连接失败", err.Error()))
				} else {
					fmt.Println("连接测试成功")
					for _, instance := range sdk.GetInstance() {
						models.DefaultVirtualMachineManager.SyncInstance(instance, platform)
						fmt.Println(instance.Name)
					}
					models.DefaultVirtualMachineManager.SyncInstanceStatus(now, platform)
					models.DefaultCloudPlatformManager.SyncInfo(platform, now,"")
				}
			}
			//fmt.Println(platform)
		}
	}
}