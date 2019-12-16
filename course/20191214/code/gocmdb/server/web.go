package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"github.com/imsilence/gocmdb/server/models"
	"github.com/imsilence/gocmdb/server/utils"

	_ "github.com/imsilence/gocmdb/server/routers"
	_ "github.com/imsilence/gocmdb/server/cloud/plugins"
)

func main() {
	// 初始化命令行参数
	h := flag.Bool("h", false, "help")
	help := flag.Bool("help", false, "help")
	init := flag.Bool("init", false, "init server")
	syncdb := flag.Bool("syncdb", false, "sync db")
	force := flag.Bool("force", false, "force sync db(drop table)")
	verbose := flag.Bool("v", false, "verbose")

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
	beego.SetLogger("file", `{
		"filename" : "logs/web.log",
		"level" : 7}`,
	)
	if !*verbose {
		//删除控制台日志
		beego.BeeLogger.DelLogger("console")
	} else {
		orm.Debug = true
	}

	// 初始化orm
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("dsn"))

	// 测试数据库连接是否正常
	if db, err := orm.GetDB(); err != nil || db.Ping() != nil {
		beego.Error("数据库连接错误")
		os.Exit(-1)
	}
	// 根据参数选择执行流程
	switch {
	case *init:
		orm.RunSyncdb("default", *force, *verbose)
		ormer := orm.NewOrm()
		admin := &models.User{Name: "admin", IsSuperman: true}
		if err := ormer.Read(admin, "Name"); err == orm.ErrNoRows {
			password := utils.RandString(6)
			admin.SetPassword(password)
			if _, err := ormer.Insert(admin); err == nil {
				beego.Informational("初始化admin成功, 默认密码:", password)
			} else {
				beego.Error("初始化用户失败, 错误:", err)
			}
		} else {
			beego.Informational("admin用户已存在, 跳过")
		}
	case *syncdb:
		orm.RunSyncdb("default", *force, *verbose)
		beego.Informational("同步数据库")
	default:
		beego.Run()
	}
}
