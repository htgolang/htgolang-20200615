package main

import "fmt"

func main() {
	sort_int := []int{10, 20, 14, 57, 22}

	for i := 0; i < len(sort_int); i++ {
		tmp := sort_int[i]
		// 当i=0时，j<0,则不比较保留原值
		// 当i=1时，将sort_int[i]赋值给tmp，与sort_int[0]进行比较
		// 当i=2时，将新的tmp值与sort_int[0]和sort_int[1]逐次比较

		for j := i - 1; j >= 0; j-- {
			if sort_int[j] > tmp {
				sort_int[j+1], sort_int[j] = sort_int[j], tmp
			}
		}
	}

	fmt.Printf("从小到大排序为: %v", sort_int)

}
