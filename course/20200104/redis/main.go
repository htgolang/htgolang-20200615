package main

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	// 认证
	reply, err := conn.Do("AUTH", "123456")
	fmt.Println(reply, err)

	reply, err = conn.Do("PING")
	fmt.Println(reply, err)

	// 字符串
	reply, err = conn.Do("SET", "test:name", "kk1")
	fmt.Println(reply, err)

	reply, err = conn.Do("SET", "test:name", "kk2")
	fmt.Println(reply, err)

	reply, err = conn.Do("SETNX", "test:name2", "kk1")
	fmt.Println(reply, err)

	reply, err = conn.Do("SETNX", "test:name2", "kk2")
	fmt.Println(reply, err)

	svalue, err := redis.String(conn.Do("GET", "test:name"))
	fmt.Println(svalue, err)

	reply, err = conn.Do("MSET", "test:1", 1, "test:2", "2", "test:3", true)
	fmt.Println(reply, err)

	ssvalue, err := redis.Strings(conn.Do("MGET", "test:1", "test:2","test:3"))
	fmt.Println(ssvalue, err)

	conn.Do("INCR", "test:1")
	conn.Do("DECR", "test:2")

	// key操作
	svalue, err = redis.String(conn.Do("TYPE", "test:name"))
	fmt.Println(svalue, err)

	bvalue, err := redis.Bool(conn.Do("EXISTS", "test:name"))
	fmt.Println(bvalue, err)

	reply, err = conn.Do("EXPIRE", "test:name", 30)
	fmt.Println(reply, err)

	time.Sleep(3 * time.Second)

	ivalue, err := redis.Int(conn.Do("TTL", "test:name"))
	fmt.Println(ivalue, err)

	reply, err = conn.Do("DEL", "test:name")
	fmt.Println(reply, err)

	// list
	reply, err = conn.Do("LPUSH", "test:list", 1, 2, 3, 4, 5)
	fmt.Println(reply, err)


	ivalue, err = redis.Int(conn.Do("LLEN", "test:list"))
	fmt.Println(ivalue, err)

	isvalue, err := redis.Ints(conn.Do("LRANGE", "test:list", 0, -1))
	fmt.Println(isvalue, err)

	ivalue, err = redis.Int(conn.Do("RPOP", "test:list"))
	fmt.Println(ivalue, err)

	ivalue, err = redis.Int(conn.Do("RPOP", "test:list"))
	fmt.Println(ivalue, err)

	ivalue, err = redis.Int(conn.Do("LLEN", "test:list"))
	fmt.Println(ivalue, err)

	ssvalue, err = redis.Strings(conn.Do("BRPOP", "test:list:block", "test:list:block2", 3))
	fmt.Println(ssvalue, err)

	ssvalue, err = redis.Strings(conn.Do("BRPOP", "test:list:block", "test:list:block2", 3))
	fmt.Println(ssvalue, err)

	// hash

	reply, err = conn.Do("HSET", "user:1", "name", "kk")
	fmt.Println(reply, err)

	reply, err = conn.Do("HSETNX", "user:1", "name", "kk2")
	fmt.Println(reply, err)

	reply, err = conn.Do("HSETNX", "user:1", "sex", "男")
	fmt.Println(reply, err)

	reply, err = conn.Do("HMSET", "user:1", "tel", "15200000000", "addr", "西安")
	fmt.Println(reply, err)

	smvalue, err := redis.StringMap(conn.Do("HGETALL", "user:1"))
	fmt.Println(smvalue, err)

	svalue, err = redis.String(conn.Do("HGET", "user:1", "name"))
	fmt.Println(svalue, err)

	smvalue, err = redis.StringMap(conn.Do("HMGET", "user:1", "name", "tel"))
	fmt.Println(smvalue, err)

	ivalue, err = redis.Int(conn.Do("HLEN", "user:1"))
	fmt.Println(ivalue, err)

	reply, err = conn.Do("HDEL", "user:1", "addr")
	fmt.Println(reply, err)

	bvalue, err = redis.Bool(conn.Do("HEXISTS", "user:1", "name"))
	fmt.Println(bvalue, err)

	bvalue, err = redis.Bool(conn.Do("HEXISTS", "user:1", "addr"))
	fmt.Println(bvalue, err)

	//SET
	reply, err = conn.Do("SADD", "test:set", 1, 2, 3, 1, 2, 3)
	fmt.Println(reply, err)

	ivalue, err = redis.Int(conn.Do("SCARD", "test:set"))
	fmt.Println(ivalue, err)

	isvalue, err = redis.Ints(conn.Do("SMEMBERS", "test:set"))
	fmt.Println(isvalue, err)

	bvalue, err = redis.Bool(conn.Do("SISMEMBER", "test:set", 1))
	fmt.Println(bvalue, err)

	reply, err = conn.Do("SREM", "test:set", 1)
	fmt.Println(reply, err)

	bvalue, err = redis.Bool(conn.Do("SISMEMBER", "test:set", 1))
	fmt.Println(bvalue, err)

	//ZSET
	reply, err = conn.Do("ZADD", "test:zset", 1, "test1", 2, "test2", 3, "test3")
	fmt.Println(reply, err)

	ivalue, err = redis.Int(conn.Do("ZCARD", "test:zset"))
	fmt.Println(ivalue, err)

	ssvalue, err = redis.Strings(conn.Do("ZRANGE", "test:zset", 0, -1))
	fmt.Println(ssvalue, err)

	ssvalue, err = redis.Strings(conn.Do("ZREVRANGE", "test:zset", 0, -1))
	fmt.Println(ssvalue, err)

	imvalue, err := redis.IntMap(conn.Do("ZRANGEBYSCORE", "test:zset", 1, 2, "WITHSCORES"))
	fmt.Println(imvalue, err)

	ssvalue, err = redis.Strings(conn.Do("ZRANGEBYSCORE", "test:zset", 1, 2, "WITHSCORES"))
	fmt.Println(ssvalue, err)

	ssvalue, err = redis.Strings(conn.Do("ZREVRANGEBYSCORE", "test:zset", 2, 1, "WITHSCORES"))
	fmt.Println(ssvalue, err)

	reply, err = conn.Do("ZREM", "test:zset", "test2")
	fmt.Println(reply, err)

	ivalue, err = redis.Int(conn.Do("ZCARD", "test:zset"))
	fmt.Println(ivalue, err)


	// 发布订阅
	go func() {
		conn, _ := redis.Dial("tcp", "localhost:6379")
		defer conn.Close()
		conn.Do("AUTH", "123456")

		psc := redis.PubSubConn{Conn: conn}
		defer psc.Close()
		psc.Subscribe("test:channel")
		for {
			switch v := psc.Receive().(type) {
			case redis.Message:
				fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
			case redis.Subscription:
				fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
			case error:
				fmt.Printf("error: %s\n", v)
			}
		}

	}()

	for i := 0; i < 10; i++ {
		conn.Do("PUBLISH", "test:channel", fmt.Sprintf("msg:%d", i))
		time.Sleep(time.Second)
	}

	time.Sleep(10 * time.Second)
}