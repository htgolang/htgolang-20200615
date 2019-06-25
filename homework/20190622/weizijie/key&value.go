package main

import "fmt"

func main() {
	// 定义map 类型user
	user := map[string]int{"wei": 90, "Jevon": 80, "dan": 25, "ran": 57}
	// 分别定义key和value的切片
	key_slice := []string{}
	value_slice := []int{}

	// 遍历user，分别将key和value添加到key_slice和value_slice切片中
	for key, value := range user {
		key_slice = append(key_slice, key)
		value_slice = append(value_slice, value)
	}

	fmt.Printf("user的key slice: %v\n", key_slice)
	fmt.Printf("user的value slice: %v\n", value_slice)

}
