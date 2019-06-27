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
	fmt.Printf("keys组成的切片:%v\nvalues组成的切片:%v\n", keys, values)

	wordCount01 := map[rune]int{}
	wordCount02 := map[int][]rune{}

	for _, word := range article {
		if word >= 'A' && word <= 'Z' || word >= 'a' && word <= 'z' {
			wordCount01[word]++
		}
	}

	for key, value := range wordCount01 {
		if v, ok := wordCount02[value]; !ok {
			wordCount02[value] = []rune{key}
		} else {
			v = append(v, key)
		}
	}
	fmt.Print("统计出现特定次数的单词：")
	for count, value := range wordCount02 {
		fmt.Printf("%d:%c, ", count, value)
	}
	fmt.Println()

	// 7. 冒泡排序，这里使用第一题的切片 ints
	n := len(ints) - 1
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			if ints[j] > ints[j+1] {
				ints[j], ints[j+1] = ints[j+1], ints[j]
			}
		}
	}
	fmt.Printf("冒泡排序：%v", ints)
}
