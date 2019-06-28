package main

import (
	"fmt"
)

func main() {
	s1 := []int{3, 1, 2, 18, 7, 5, 9, 12, 10}
	s2 := [2]int{}
	for _, v1 := range s1 {
		for _, v2 := range s1 {
			if v1 < v2 {
				s2[0] = v2
			} else {
				s2[0] = v1
			}
		}
	}
	for _, v1 := range s1 {
		for _, v2 := range s1 {
			if v1 < v2 && v2 != s2[0] {
				s2[1] = v2
			}
		}
	}
	fmt.Println(s2[1])
}

/*
 评分: 6
*/
