package main

import (
	"flag"
	"fmt"
	"todolist/models"
	"github.com/astaxie/beego/orm"
	"todolist/utils"
	_ "github.com/lib/pq"
	"github.com/astaxie/beego"
	_ "todolist/routers"
)

func main() {
	init := flag.Bool("init", false, "init admin")
	force := flag.Bool("force", false, "force clear database")
	flag.Parse()

	orm.RegisterDataBase("default","postgres", "user=test password=test dbname=test host=127.0.0.1 port=5432 sslmode=disable", 30)

	if *init {
		orm.RunSyncdb("default", *force, true)

		user := models.User{Name:"admin", IsSuper:true}
		password := utils.RandomString(6)
		user.SetPassword(password)

		if err := models.AddUser(&user); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Admin password: %s\n", password)
		}
	} else {
		fmt.Println("Run todolist app")
		beego.Run()
	}

}
