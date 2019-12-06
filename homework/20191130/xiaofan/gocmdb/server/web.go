package main

import (
	"flag"
	"fmt"
	_ "gocmdb/cloud/plugins"
	"gocmdb/models"
	_ "gocmdb/routers"
	"gocmdb/utils"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

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
		beego.Run()
	}
}
