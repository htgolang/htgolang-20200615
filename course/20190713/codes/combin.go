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

func main() {
	var me01 User

	fmt.Printf("%#v\n", me01)

	addr := Address{"西安市", "锦业路", "01"}
	me02 := User{
		ID:   1,
		Name: "KK",
		Addr: addr,
	}

	fmt.Printf("%#v\n", me02)

	me03 := User{
		ID:   2,
		Name: "WOniu",
		Addr: Address{
			Region: "北京市",
			Street: "海淀路",
			No:     "002",
		},
	}
	fmt.Printf("%#v\n", me03)

	fmt.Println(me03.Addr.Region)
	me03.Addr.Region = "西安市"
	fmt.Printf("%#v\n", me03)
}
