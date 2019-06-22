package main

import "fmt"

func main() {
	heights := []int{10, 6, 7, 9, 5}
	// 先把最高的人移动到最后
	// 从第一个人开始，两两开始比较，如果前面的人高，这两人就交换位置
	// 1: 10, 6 => 交换 => 6, 10, 7, 9, 5
	// 2: 10, 7 => 交换 => 6, 7, 10, 9, 5
	// 3: 10, 9 => 交换 => 6, 7, 9, 10, 5
	// 4: 10, 5 => 交换 => 6, 7, 9, 5, 10

	// 第二轮
	// 1: 6, 7 => 不交换 => 6, 7, 9, 5, 10
	// 2: 7, 9 => 不交换 => 6, 7, 9, 5, 10
	// 3: 9, 5 => 交换 => 6, 7, 5, 9, 10
	// 4: 9, 10 => 不交换 => 6, 7, 5, 9, 10

	// 第三轮

	// 1: 不交换
	// 2: 交换6, 5, 7, 9, 10
	// 3: 不交换
	// 4:  不交换

	// 第四轮
	// 1 交换: 5, 6, 7, 9, 10

	// n 个人比较 n-1 5
	// 0, 1, 2, 3

	// 交换 a = 1, b = 2  a,b
	// 左 => 苹果  右  => 梨
	// 左手 => 苹果 => 放在桌子上 (左手空)
	// 右手 => 梨 => 左手 (左手梨，右手空)
	// 桌子 => 苹果 => 右手(右手苹果, 左手梨)
	// tmp := a
	// a = b
	// b = tmp
	// 方法二
	// a, b = b, a

	// for j := 0; j < len(heights)-1; j++ {
	// 	fmt.Println("第", j, "轮")
	// 	for i := 0; i < len(heights)-1; i++ {
	// 		if heights[i] > heights[i+1] {
	// 			fmt.Println("交换：", heights[i], heights[i+1])
	// 			tmp := heights[i]
	// 			heights[i] = heights[i+1]
	// 			heights[i+1] = tmp
	// 		}
	// 		fmt.Println(i, "交换完毕", heights)
	// 	}
	// 	fmt.Println("第", j, "轮, 结果:", heights)
	// }

	for j := 0; j < len(heights)-1; j++ {
		fmt.Println("第", j, "轮")
		for i := 0; i < len(heights)-1-j; i++ {
			if heights[i] > heights[i+1] {
				fmt.Println("交换：", heights[i], heights[i+1])
				tmp := heights[i]
				heights[i] = heights[i+1]
				heights[i+1] = tmp
			}
			fmt.Println(i, "交换完毕", heights)
		}
		fmt.Println("第", j, "轮, 结果:", heights)
	}

}
