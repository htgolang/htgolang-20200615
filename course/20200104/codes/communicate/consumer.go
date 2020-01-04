package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, _ := redis.Dial("tcp", "localhost:6379")

	conn.Do("AUTH", "jKAFXjnE5f8kwFG3kFpcetPTlQHuIzXb")

	for {
		values, err := redis.Strings(conn.Do("BRPOP", "test:logs", 3))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(values[1])
		}
	}
}
