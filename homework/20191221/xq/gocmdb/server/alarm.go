package main

import (
	"flag"
	"fmt"

	"os"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/xlotz/gocmdb/server/routers"
	"github.com/xlotz/gocmdb/server/models"

)

func main() {
	// 初始化命令行参数
	h := flag.Bool("h", false, "help")
	help := flag.Bool("help", false, "help")
	//verbose := flag.Bool("v", false, "verbose")

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
	//if !*verbose {
	//	//删除控制台日志
	//	beego.BeeLogger.DelLogger("console")
	//
	//} else {
		orm.Debug = true
	//}

	// 初始化orm
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("dsn"))

	// 测试数据库连接是否正常
	if db, err := orm.GetDB(); err != nil || db.Ping() != nil {
		beego.Error("数据库连接错误")
		os.Exit(-1)
	}
	go func() {
		// 判断心跳时间超过1分钟
		for now := range time.Tick(time.Minute){
		//for now := range time.Tick(time.Second) {
			offlineTime := 5
			endTime := now.Add(-1 * time.Duration(offlineTime) * time.Minute)

			var result []orm.Params

			orm.NewOrm().Raw("select uuid from agent where deleted_time is null and heartbeat_time < ?", endTime).Values(&result)

			models.DefaultAlarmManager.Create(1, result, "离线超过5分钟")

		}
	}()

	//// cpu 使用率
	go func() {
		for now := range time.Tick(time.Second){

			windowTime := 1

			cpuThreshold := 1
			cpuCounter := 1

			startTime := now.Add(-1 * time.Duration(windowTime) * time.Minute) // 5 可根据配置

			var result []orm.Params

			orm.NewOrm().Raw("select uuid, count(*) as cnt from resource where deleted_time is null and created_time > ?  and cpu_precent > ? group by uuid having cnt >= ?", startTime, cpuThreshold, cpuCounter).Values(&result)

			fmt.Println(result)

			models.DefaultAlarmManager.Create(2, result, "CPU在5分钟内, 连续3次 超过10%")



		}

	}()

	//// MEM 使用率
	go func() {
		for now := range time.Tick(time.Minute){

			windowTime := 5

			memThreshold := 10
			memCounter := 3

			startTime := now.Add(-1 * time.Duration(windowTime) * time.Minute) // 5 可根据配置

			var result []orm.Params

			orm.NewOrm().Raw("select uuid, count(*) as cnt from resource where deleted_time is null and created_time > ?  and mem_precent > ? group by uuid having cnt >= ?", startTime, memThreshold, memCounter).Values(&result)

			fmt.Println(result)

				models.DefaultAlarmManager.Create(3, result, "MEM使用率在5分钟内，连续3次超过10%")

		}

	}()

}

