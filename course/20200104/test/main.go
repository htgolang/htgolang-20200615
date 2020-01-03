package main

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/context"

	_ "github.com/astaxie/beego/cache/redis"
)

func main() {
	memory, _ := cache.NewCache("memory", `{"interval":60}`)
	redis, _ := cache.NewCache("redis", `{"key":"cache","conn":"127.0.0.1:6379","dbNum":"0","password":""}`)

	beego.Get("/memory/", func(ctx *context.Context) {
		key := "cachetest"

		appname := beego.AppConfig.String("appname")
		node := beego.AppConfig.String("node")
		if !memory.IsExist(key) {
			memory.Put(key, 0, 60*time.Second)
		}
		memory.Incr(key)
		ctx.Output.Body([]byte(fmt.Sprintf("%s: %s: %d", appname, node, memory.Get(key))))
	})

	beego.Get("/redis/", func(ctx *context.Context) {
		key := "cachetest"

		appname := beego.AppConfig.String("appname")
		node := beego.AppConfig.String("node")
		if !redis.IsExist(key) {
			redis.Put(key, 0, 60*time.Second)
		}
		redis.Incr(key)
		ctx.Output.Body([]byte(fmt.Sprintf("%s: %s: %s", appname, node, redis.Get(key))))
	})
	beego.Run()
}
