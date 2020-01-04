package main

import (
	"fmt"
	"flag"
	"time"

	"github.com/gomodule/redigo/redis"
)

func main() {
	node := flag.String("node", "A", "node name")
	flag.Parse()


	masterKey := "test:worker:master"
	conn, _ := redis.Dial("tcp", "localhost:6379")
	defer conn.Close()

	conn.Do("AUTH", "jKAFXjnE5f8kwFG3kFpcetPTlQHuIzXb")

	for now := range time.Tick(10 * time.Second) {
		if value, err := redis.String(conn.Do("SET", masterKey, *node, "EX", 9 , "NX")); err == nil && value == "OK" {
			fmt.Println("node is master:", *node, now)
		} else {
			value, _ := redis.String(conn.Do("GET", masterKey))
			fmt.Println("masetr:", value)
		}
	}

}
