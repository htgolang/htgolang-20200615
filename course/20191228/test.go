package main

import (
	"fmt"
)

func main() {
	a := 1

	b := &a

	c := []*int{}
	c = append(c, b)
	c = append(c, b)
	c = append(c, b)

	*b = 3

	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(*c[0], *c[1], *c[2])

	s := "abcdef"
	fmt.Println(string([]rune(s)[:3]))
	s = "a"
	fmt.Println(string([]rune(s)[:3]))
	s = "我爱中国"
	fmt.Println(string([]rune(s)[:3]))
	s = "我爱"
	fmt.Println(string([]rune(s)[:3]))
}

