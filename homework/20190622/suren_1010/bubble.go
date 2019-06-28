package main

import "fmt"

func main() {

	sliceData := []int{1, 3, 10, 5, 7, 20, 100}
	for i := 0; i < len(sliceData)-1; i++ {
		for j := 0; j < len(sliceData)-1-i; j++ {
			if sliceData[j] > sliceData[j+1] {
				sliceData[j], sliceData[j+1] = sliceData[j+1], sliceData[j]
			}
		}
	}
	fmt.Println(sliceData)

}

/*
 评分: 8
*/
