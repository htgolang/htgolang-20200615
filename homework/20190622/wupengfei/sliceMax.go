package main

import "fmt"

func main() {
	//找int切片中最大的元素(不准用排序)

	var numslice []int = []int{1, 22, 66, 7, 99, 3, 44}

	for i := 0; i < len(numslice)-1; i++ {
		if numslice[i] > numslice[i+1] {
			tmp := numslice[i]
			numslice[i] = numslice[i+1]
			numslice[i+1] = tmp
		}
	}
	fmt.Println(numslice[len(numslice)-1])
}
