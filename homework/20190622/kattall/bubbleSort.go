package main

import "fmt"

// 排序函数，传入一个切片int  返回一个切片int
func bubble(sli []int) []int {
	for j := 0; j < len(sli)-1; j++ {
		for i := 0; i < len(sli)-j-1; i++ {
			if sli[i] > sli[i+1] {
				sli[i], sli[i+1] = sli[i+1], sli[i]
			}
		}
	}

	return sli
}

func main() {
	height := []int{6, 7, 9, 10, 5, 3, 1, 4, 8}
	fmt.Println("排序前的切片为：", height)

	sli := bubble(height)
	fmt.Println("排序后的切片为：", sli)
}
