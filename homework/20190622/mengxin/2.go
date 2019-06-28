package main

import "fmt"

func main() {
	s1 := [6]int{3, 4, 5, 1, 8, 2}
	max := s1[0]

	for _, v := range s1 {
		if v > max {
			max = v
		}
	}
	fmt.Println(max)
}

/*
 评分: 8
*/
