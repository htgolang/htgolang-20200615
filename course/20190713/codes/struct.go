package main

import (
	"fmt"
	"time"
)

type User struct {
	ID       int
	Name     string
	Birthday time.Time
	Addr     string
	Tel      string
	Remark   string
}

func main() {
	// var counter Counter
	var me User
	fmt.Printf("%T\n", me)
	fmt.Printf("%#v\n", me)

	var me2 User = User{
		1,
		"KK",
		time.Now(),
		"西安市",
		"150XXXXXXXX",
		"kk",
	}
	fmt.Printf("%#v\n", me2)

	var me3 User = User{} // var me3 User
	fmt.Printf("%#v\n", me3)

	var me4 User = User{
		Name: "KK",
		ID:   1,
		Addr: "西安",
		Tel:  "11231231320",
	}
	fmt.Printf("%#v\n", me4)

	var pointer *User
	fmt.Printf("%T\n", pointer)
	fmt.Printf("%#v\n", pointer)

	var pointer2 *User = &me2

	fmt.Printf("%#v\n", pointer2)

	var pointer3 *User = &User{}
	fmt.Printf("%#v\n", pointer3)

	var pointer4 *User = new(User)
	fmt.Printf("%#v\n", pointer4)

	var pointer5 *int = new(int)
	fmt.Printf("%#v\n", pointer5)
}
