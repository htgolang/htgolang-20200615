package main

import "fmt"

// a 源 b借助的 c 目的
func tower(a, b, c string, layer int) {
	if layer == 1 {
		fmt.Println(a, "->", c)
		return
	}
	// a n-1  借助 c  移动到 b
	tower(a, c, b, layer-1)
	fmt.Println(a, "->", c)
	// b n -1 借助a 移动到 c
	tower(b, a, c, layer-1)
}

func main() {
	fmt.Println("1层:")
	tower("A", "B", "C", 1)
	fmt.Println("2层:")
	tower("A", "B", "C", 2)
	fmt.Println("3层:")
	tower("A", "B", "C", 3)
	fmt.Println("4层:")
	tower("A", "B", "C", 4)
}
