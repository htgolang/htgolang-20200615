package main

import "fmt"

func main() {
	var nums [10]int
	var t2 [5]bool
	var t3 [3]string

	fmt.Printf("%T", nums)
	fmt.Println(nums)
	fmt.Println(t2)
	fmt.Printf("%q\n", t3)

	// 字面量
	nums = [10]int{10, 20, 30}
	fmt.Println(nums)

	nums = [10]int{0: 10, 9: 20}
	fmt.Println(nums)

	nums = [...]int{1, 2, 3, 1, 2, 3, 1, 2, 3, 1}
	fmt.Println(nums)

	nums02 := [5]int{10, 20, 30}
	fmt.Printf("%T %#v\n", nums02, nums02)

	nums03 := [...]int{1, 2}
	fmt.Printf("%T %#v\n", nums03, nums03)

	nums04 := [15]int{1: 10, 5: 20, 14: 30}
	fmt.Printf("%T %#v\n", nums04, nums04)

	// 操作
	nums05 := [3]int{1, 3, 4}
	nums06 := [3]int{1, 3, 4}
	fmt.Println(nums05 == nums06)
	fmt.Println(nums05 != nums06)

	//获取数组的长度
	fmt.Println(len(nums04))

	// 索引 0, 1, 2, ... len(array) - 1
	fmt.Println(nums04[0], nums04[1])
	nums04[0] = 1000
	fmt.Println(nums04)

	// for i := 0; i < len(nums04); i++ {
	// 	fmt.Println(i, ":", nums04[i])
	// }

	// for _, value := range nums04 {
	// 	fmt.Println(value)
	// }

	// 切片
	fmt.Printf("%T\n", nums04[0:15])
	fmt.Printf("%T\n", nums04[1:15:15])
	fmt.Printf("%v\n", nums04[1:3])

}
