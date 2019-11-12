package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/imsilence/todolist/routers"

	"github.com/imsilence/todolist/models"
	"github.com/imsilence/todolist/utils"
)

func main() {

	// 通过flag包接收命令行参数
	help := flag.Bool("help", false, "help")
	h := flag.Bool("h", false, "help")

	init := flag.Bool("init", false, "init Db")                     //初始化数据库及超级管理员账号
	syncdb := flag.Bool("syncdb", false, "Sync Db")                 // 同步数据库
	force := flag.Bool("force", false, "force sync db(drop table)") //强制删除数据表并同步
	verbose := flag.Bool("verbose", false, "verbose")               //显示同步数据库信息

	// 命令行帮助函数
	flag.Usage = func() {
		fmt.Println("Usage: todolist [--web] [--syncdb]")
		flag.PrintDefaults()
	}

	flag.Parse() //解析命令行参数

	// 命令行输入帮助参数，打印帮助信息
	if *h || *help {
		flag.Usage()
		os.Exit(0)
	}
	if *verbose {
		orm.Debug = true
	}

	// 注册mysql驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// 注册mysql数据库连接（dsn信息从配置文件中获取）
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("dsn"))

	// 检查数据连接状态
	if db, err := orm.GetDB(); err != nil || db.Ping() != nil {
		log.Fatal("数据库连接失败")
	}

	switch {
	case *init:
		// 同步数据库
		orm.RunSyncdb("default", *force, *verbose)

		//获取数据库连接
		ormer := orm.NewOrm()

		// 为admin账号随机生成6位密码
		name, password := "admin", utils.RandString(6)

		user := &models.User{Name: name, IsSuper: true}

		// 添加admin账号到数据库中
		if err := ormer.Read(user, "Name"); err == orm.ErrNoRows {
			user.SetPassword(password)
			if _, err := ormer.Insert(user); err == nil {
				fmt.Printf("初始化用户: %s, 密码: %s\n", name, password)
			}
		} else {
			fmt.Println("admin用户已存在, 跳过初始化...")
		}
	case *syncdb:

		// 同步数据库
		orm.RunSyncdb("default", *force, *verbose)
	default:

		//启动web服务
		beego.Run()
	}
}
