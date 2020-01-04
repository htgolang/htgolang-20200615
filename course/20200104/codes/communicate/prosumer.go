package main

import (
	"time"
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, _ := redis.Dial("tcp", "localhost:6379")

	conn.Do("AUTH", "jKAFXjnE5f8kwFG3kFpcetPTlQHuIzXb")

	for {
		conn.Do("LPUSH", "test:logs", time.Now())
		time.Sleep(time.Second)
	}
}
