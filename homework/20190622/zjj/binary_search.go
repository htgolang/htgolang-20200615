package main

import (
	"fmt"
)

//二分查找，取中间，只适用于有序的序列

func binary_search(input []int, guess int) (index int, stat int) {
	v_low := 0                //取切片第一位
	v_heigh := len(input) - 1 //取切片最后一位
	for { //循环查找，直到出现结果
		v_middle := (v_low + v_heigh) / 2 //取中间数
		if v_heigh < v_low { //查询不出结果的条件
			index = v_middle //将查询最后结果赋值给index
			stat = 0
			break
		}
		if input[v_middle] > guess {
			v_heigh = v_middle - 1 //最大值取中间数向左移一位
		} else if input[v_middle] < guess {
			v_low = v_middle + 1 //最小值取中间数向右移一位
		} else {
			index = v_middle
			stat = 1
			break
		}
	}
	return index, stat
}

func main() {
	guess_num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}       //定义一个有序的切片
	guess_index, guess_stat := binary_search(guess_num, -1) //调用函数将返回值赋予给变量
	fmt.Println(guess_index)
	if guess_stat == 1 { //根据返回值的状态进行输出判断
		fmt.Println("查找结果成功，查询到的位置:", guess_index, "值为:", guess_num[guess_index]) //列出索引及值 查找到的情况
	} else {
		fmt.Println("查找结果失败，最后查询位置:", guess_index, "近似值为:", guess_num[guess_index]) //查找不到，列出近似值
	}
}
