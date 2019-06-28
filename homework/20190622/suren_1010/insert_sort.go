package main

import "fmt"

func main() {

	sliceData := []int{1, 5, 20, 17, 3, 9}
	for i := 1; i < len(sliceData); i++ {
		for j := i - 1; j >= 0 && sliceData[j] > sliceData[j+1]; j-- {
			sliceData[j], sliceData[j+1] = sliceData[j+1], sliceData[j]
		}
	}
	fmt.Println(sliceData)
}

/*
 评分: 8
*/
