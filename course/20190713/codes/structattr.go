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
	me := User{
		Name: "KK",
		ID:   1,
		Addr: "西安",
		Tel:  "11231231320",
	}

	fmt.Println(me.ID, me.Name, me.Tel)

	me.Tel = "15200000000"
	fmt.Printf("%#v\n", me)

	me2 := &User{
		ID:   2,
		Name: "woniu",
	}
	fmt.Printf("%T\n", me2)

	(*me2).Tel = "1520000000"
	fmt.Printf("%#v\n", me2)

	me2.Addr = "北京"
	fmt.Printf("%#v\n", me2)
}
