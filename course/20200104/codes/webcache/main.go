package main

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/cache"

	_ "github.com/astaxie/beego/cache/redis"
)

func main() {
	memoryCache, _ := cache.NewCache("memory", `{"interval" : 60}`)
	redisCache, _ := cache.NewCache("redis", `{"key" : "cache", "conn" : "localhost:6379", "dbNum" : 0, "password":"jKAFXjnE5f8kwFG3kFpcetPTlQHuIzXb" }`)

	node := beego.AppConfig.String("node")

	key := "counter"

	beego.Get("/memory/", func(ctx *context.Context) {
		if !memoryCache.IsExist(key) {
			memoryCache.Put(key, 0, 60*time.Second)
		}
		memoryCache.Incr(key)

		ctx.Output.Body([]byte(fmt.Sprintf("%s: memory: %d", node, memoryCache.Get(key))))
	})

	beego.Get("/redis/", func(ctx *context.Context) {
		if !redisCache.IsExist(key) {
			redisCache.Put(key, 0, 60*time.Second)
		}
		redisCache.Incr(key)
		ctx.Output.Body([]byte(fmt.Sprintf("%s: redis: %s", node, redisCache.Get(key))))
	})
	beego.Run()
}
