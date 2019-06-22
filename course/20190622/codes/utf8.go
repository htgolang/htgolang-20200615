package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "我爱中华人民共和国"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
}
