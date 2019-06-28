package main

import "fmt"

func main() {

	s1 := []int{10, 7, 13, 8, 21, 6, 11, 14, 5}

	for j := 0; j < len(s1)-1; j++ {
		for i := 0; i < len(s1)-1-j; i++ {
			if s1[i] > s1[i+1] {
				s1[i], s1[i+1] = s1[i+1], s1[i]
			}
		}
	}

	fmt.Println(s1)

}

/*
 评分: 8
*/
