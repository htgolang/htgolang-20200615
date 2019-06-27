package main

import "fmt"

func main() {

	sliceData := []int{9, 10, 8, 4, 1, 3, 20}

	for i := 1; i < len(sliceData); i++ {
		low := 0
		high := i
		for low <= high {
			middle := (low + high) / 2
			if sliceData[middle] > sliceData[i] {
				high = middle - 1
			} else {
				low = middle + 1
			}
		}
		for j := i - 1; j >= low; j-- {
			sliceData[j], sliceData[j+1] = sliceData[j+1], sliceData[j]
		}
	}
	fmt.Println(sliceData)

}
