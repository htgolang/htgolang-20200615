package main

import (
	"fmt"
)

func main() {
	var k1 []string
	var v1 []int
	m1 := map[string]int{"age": 18, "height": 180}
	for k, v := range m1 {
		k1 = append(k1, k)
		v1 = append(v1, v)
	}
	fmt.Println(k1, v1)
}
