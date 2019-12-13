package main

import (
	"flag"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/dcosapp/gocmdb/server/cloud"
	_ "github.com/dcosapp/gocmdb/server/cloud/plugins"
	"github.com/dcosapp/gocmdb/server/models"
	_ "github.com/dcosapp/gocmdb/server/routers"
	"github.com/dcosapp/gocmdb/server/utils"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

func sync() {
	for now := range time.Tick(10 * time.Second) {
		// 获取所有云平台
		platforms, _, _ := models.DefaultCloudPlatformManager.Query("", 0, 0)
		for _, platform := range platforms {
			// 云平台是否启用
			if platform.IsEnable() {
				// 云平台是否注册
				if sdk, ok := cloud.DefaultManager.Cloud(platform.Type); !ok {
					beego.Info("云平台未注册: ", platform.Name)
				} else {
					sdk.Init(platform.Addr, platform.Region, platform.AccessKey, platform.SecretKey)

					if err := sdk.TestConnect(); err != nil {
						beego.Error("连接测试失败: ", err.Error())
						// 将失败信息同步至数据库
						_ = models.DefaultCloudPlatformManager.SyncInfo(platform, now, fmt.Sprintf("连接测试失败,%s", err.Error()))
					} else {
						for _, instance := range sdk.GetInstance() {
							models.DefaultVirtualMachineManager.SyncInstance(instance, platform)
							beego.Debug("Platform: ", platform.Name, ";Instance: ", instance)
						}
						models.DefaultVirtualMachineManager.SyncInstacneStatus(now, platform)
						_ = models.DefaultCloudPlatformManager.SyncInfo(platform, now, "同步成功")
					}
				}
			}
		}
	}
}

func main() {
	// 设置命令行参数
	h := flag.Bool("h", false, "help")
	help := flag.Bool("help", false, "help")
	init := flag.Bool("init", false, "init server")
	syncdb := flag.Bool("syncdb", false, "sync db")
	force := flag.Bool("force", false, "force sync db(drop table)")
	verbose := flag.Bool("v", false, "verbose")

	// 设置使用帮助
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

	// 设置日志到文件
	_ = beego.SetLogger("file", `{"filename":"logs/web.log","level":7}`)

	// 初始化数据库
	_ = orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("dsn"))

	if !*verbose {
		// 删除控制台日志
		_ = beego.BeeLogger.DelLogger("console")
	} else {
		//orm.Debug = true
	}

	// 测试数据库连接是否正常
	if db, err := orm.GetDB(); err != nil || db.Ping() != nil {
		beego.Error("数据库连接错误", err.Error())
		os.Exit(-1)
	}

	// 根据参数选择流程
	switch {
	case *init:
		// 创建表并生成admin用户
		_ = orm.RunSyncdb("default", *force, *verbose)
		o := orm.NewOrm()
		admin := &models.User{Name: "admin", IsSuperman: true}

		if err := o.Read(admin, "Name"); err == orm.ErrNoRows { // 如果admin不存在
			password := utils.RandString(6)            // 生成随机6位密码
			admin.SetPassword(password)                // 设置密码
			if _, err := o.Insert(admin); err == nil { // 插入admin数据
				beego.Informational("初始化admin成功，默认密码:", password)
			} else {
				beego.Error("初始化admin失败，错误:", err.Error())
			}
		} else {
			beego.Informational("admin用户已存在，跳过")
		}
	case *syncdb:
		// 创建表
		_ = orm.RunSyncdb("default", *force, *verbose)
		beego.Informational("同步数据库")
	default:
		go sync()
		beego.Run()
	}
}
