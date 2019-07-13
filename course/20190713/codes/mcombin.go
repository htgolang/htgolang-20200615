package main

import (
	"fmt"
)

type Address struct {
	Region string
	Street string
	No     string
}

func (addr Address) String() string {
	return fmt.Sprintf("%s-%s-%s", addr.Region, addr.Street, addr.No)
}

type User struct {
	ID   int
	Name string
	Addr Address
}

func (user User) String() string {
	return fmt.Sprintf("[%d]%s: %s", user.ID, user.Name, user.Addr)
}

func main() {
	var u User = User{
		ID:   1,
		Name: "kk",
		Addr: Address{"西安市", "锦业路", "001"},
	}
	fmt.Println(u)
	fmt.Println(u.Addr)
}
