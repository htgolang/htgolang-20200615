package main

import "fmt"

func test() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()

	panic("error")
	return
}

func main() {
	err := test()
	fmt.Println(err)
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println(err)
	// 	}
	// }()

	// fmt.Println("main start")
	// panic("error xxxx")

	// fmt.Println("over")
}
