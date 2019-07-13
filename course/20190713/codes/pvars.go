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
	Addr *Address
}

func NewUser(id int, name string, region, street, no string) *User {
	return &User{
		ID:   id,
		Name: name,
		Addr: &Address{region, street, no},
	}
}

func main() {
	me := User{
		ID:   1,
		Name: "KK",
		Addr: &Address{"西安市", "锦业路", "0001"},
	}

	me2 := me
	me2.Name = "kk"
	me2.Addr.Street = "未央路"

	fmt.Printf("%#v\n", me.Addr)
	fmt.Printf("%#v\n", me2.Addr)

	woniu := NewUser(2, "woniu", "北京市", "海淀路", "0001")
	fmt.Printf("%#v\n", woniu)
}
