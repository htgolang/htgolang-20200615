package main

import "fmt"

func main() {
	var nums = [6]int{6, 2, 8, 3, 7, 9}
	var maxNum int
	maxNum = nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
		}
	}

	// 找出切片中最大的一个数

	fmt.Printf("切片中最大一个数:%d \n", maxNum)

	// 找出切片中第二大的一个数

	SecondNum := 0
	for _, value := range nums {
		if value > SecondNum && value != maxNum {
			SecondNum = value
		}
	}
	fmt.Printf("切片中第二大的数:%d \n", SecondNum)

	dicts := map[string]int{"tim": 87, "ava": 98, "tom": 52}
	var keys []string
	var values []int
	for k, v := range dicts {
		keys = append(keys, k)
		values = append(values, v)
	}
	//  打印映射中所有key组成的切片和打印映射中所有value组成的切片
	fmt.Printf("keys组成的切片:%v\nvalues组成的切片:%v\n",keys, values)
}
