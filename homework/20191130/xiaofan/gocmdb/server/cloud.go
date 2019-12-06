package main

import (
	"flag"
	"fmt"
	"gocmdb/models"
	_ "gocmdb/routers"
	"os"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func mainm() {
	// 设置命令行参数
	h := flag.Bool("h", false, "help")
	help := flag.Bool("help", false, "help")
	verbose := flag.Bool("v", false, "verbose")

	// 设置使用帮助
	flag.Usage = func() {
		fmt.Println("usage: cloud -h")
		flag.PrintDefaults()
	}

	// 解析命令行参数
	flag.Parse()

	if *h || *help {
		flag.Usage()
		os.Exit(0)
	}

	// 设置日志到文件
	_ = beego.SetLogger("file", `{"filename":"logs/cloud.log","level":7}`)

	// 初始化数据库
	_ = orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("dsn"))

	if !*verbose {
		// 删除控制台日志
		_ = beego.BeeLogger.DelLogger("console")
	} else {
		orm.Debug = true
	}

	// 测试数据库连接是否正常
	if db, err := orm.GetDB(); err != nil || db.Ping() != nil {
		beego.Error("数据库连接错误", err.Error())
		os.Exit(-1)
	}

	for now := range time.Tick(10 * time.Second) {
		platforms, _, _ := models.DefaultCloudPlatformManager.Query("", 0, 0)
		for _, platform := range platforms {
			if platform.IsEnable() {
				fmt.Println(platform, now)
			}

		}

	}

}
