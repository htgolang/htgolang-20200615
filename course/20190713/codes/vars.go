package main

import "fmt"

type Address struct {
	Region string
	Street string
	No     string
}

type User struct {
	ID   int
	Name string
	Addr Address
}

func change(u User) {
	u.Name = "xxxx"
}

func changePoint(u *User) {
	u.Name = "yyy"
}

func main() {
	me := User{}
	me2 := me
	me2.Name = "kk"
	me2.Addr.Street = "未央路"

	fmt.Printf("%#v\n", me)
	fmt.Printf("%#v\n", me2)

	change(me2)
	fmt.Printf("%#v\n", me2)

	changePoint(&me2)
	fmt.Printf("%#v\n", me2)

}
