package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	_ "github.com/imsilence/gocmdb/server/routers"
)


func main() {
	// 初始化命令行参数
	h := flag.Bool("h", false, "help")
	help := flag.Bool("help", false, "help")
	verbose := flag.Bool("v", false, "verbose")

	flag.Usage = func() {
		fmt.Println("usage: alarm -h")
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
		"filename" : "logs/alarm.log",
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

	go func() {
		// 离线告警
		for now := range time.Tick(time.Minute) {
			fmt.Println("离线告警", now)
			offlineTime := 5
			endTime := now.Add(-1 * time.Duration(offlineTime) * time.Minute) // 5 根据配置
			var result []orm.Params
			orm.NewOrm().Raw("SELECT uuid from agent where deleted_time is null and heartbeat_time < ?", endTime).Values(&result)
			fmt.Println(result)
		}
	}()


	go func() {

		// CPU使用率
		for now := range time.Tick(time.Minute) {
			fmt.Println("CPU使用率告警", now)
			windowTime := 5
			cpuThreshold := 30
			cpuCounter := 3
			startTime := now.Add(-1 * time.Duration(windowTime) * time.Minute) // 5 根据配置
			var result []orm.Params
			orm.NewOrm().Raw("SELECT uuid, count(*) as cnt from resource where deleted_time is null and created_time >= ? and cpu_percent >= ? group by uuid having count(*) >= ?", startTime, cpuThreshold, cpuCounter).Values(&result)
			fmt.Println(result)
		}
	}()

	// 内存使用率
	for now := range time.Tick(time.Minute) {
		fmt.Println("内存使用率告警", now)
		windowTime := 5
		ramThreshold := 40
		ramCounter := 3
		startTime := now.Add(-1 * time.Duration(windowTime) * time.Minute) // 5 根据配置
		var result []orm.Params
		orm.NewOrm().Raw("SELECT uuid, count(*) as cnt from resource where deleted_time is null and created_time >= ? and ram_percent >= ? group by uuid having count(*) >= ?", startTime, ramThreshold, ramCounter).Values(&result)
		fmt.Println(result)
	}

}
