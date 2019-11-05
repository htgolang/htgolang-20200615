package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
	_ "todolist/routers"

	"todolist/models"
	"todolist/utils"
)

func main() {

	//通过flag包接收命令行参数
	help := flag.Bool("help", false, "help")
	h := flag.Bool("h", false, "help")

	init := flag.Bool("init", false, "Init admin")             // 初始化数据库及超级管理员账号
	syncdb := flag.Bool("syncdb", false, "Sync Db")            // 同步数据库
	force := flag.Bool("force", false, "force clear database") // 强制删除数据表并同步
	verbose := flag.Bool("verbose", false, "verbose")          // 显示同步数据库信息

	// 命令行帮助参数
	flag.Usage = func() {
		fmt.Println("Usage: todolist [--syncdb] or [--init]")
		flag.PrintDefaults()
	}

	flag.Parse() // 解析命令行参数

	// 命令行输入帮助参数，打印帮助信息
	if *h || *help {
		flag.Usage()
		os.Exit(0)
	}

	// 注册数据库连接 (dns信息从配置文件中获得)
	if err := orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("dsn")); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Println(1)
	// 检查数据库连接状态
	if db, err := orm.GetDB(); err != nil || db.Ping() != nil {
		fmt.Printf("数据库连接失败，%s", err)
	}

	switch {
	case *init:
		// 同步数据库
		if err := orm.RunSyncdb("default", *force, *verbose); err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		// 获取数据库连接
		o := orm.NewOrm()

		// 为admin账号生成随机6位密码
		name, password := "admin", utils.RandomString(6)

		user := models.User{Name: name, IsSuper: true}

		// 添加admin账号到数据库中
		if err := o.Read(&user, "Name"); err == orm.ErrNoRows {
			user.SetPassword(password)
			if _, err := o.Insert(&user); err == nil {
				fmt.Printf("初始化用户：%s, 密码：%s\n", name, password)
			}
		} else {
			fmt.Println("admin用户已存在， 跳过初始化")
		}
	case *syncdb:
		// 同步数据库
		if err := orm.RunSyncdb("default", *force, *verbose); err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
	default:
		fmt.Println("启动服务.....")
		beego.Run()

	}

}
