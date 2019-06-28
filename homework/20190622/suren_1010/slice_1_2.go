package main

import "fmt"

func main() {
	sliceData := []int{1, 3, 5, 2, 4, 7}

	for i := 0; i < len(sliceData)-1; i++ {
		if sliceData[i] > sliceData[i+1] {
			sliceData[i+1], sliceData[i] = sliceData[i], sliceData[i+1]
		}
	}
	fmt.Println("Slice Max is: ", sliceData[len(sliceData)-1:])
	fmt.Println("Slice Second is: ", sliceData[len(sliceData)-2:len(sliceData)-1])
}

/*
 评分: 8
*/
