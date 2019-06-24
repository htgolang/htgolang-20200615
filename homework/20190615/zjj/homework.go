package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	/*
		打印乘法口诀 for + if + fmt.Println + fmt.Printf
		猜数字游戏 for if continue/break a. 生成一个[0, 100)随机数 b. 让用户最多猜 5次（从命令行输入一个整数） 猜的太大 => 提示太大了，你还有N此猜测机会 猜的太小 => 提示太小了，你还有N此猜测机会 猜中了 => 猜中了
		 5次都没猜中 => 退出，并提示太笨了
		挑战： 当5此都没有猜中：提示太笨了，重新开始猜数游戏（重新生成随机数，重新开始5次计数）
	*/

	//打印99乘法表
	for i := 1; i <= 9; i++ { //控制行
		for j := 1; j <= i; j++ { //控制列
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Printf("\n")
	}

	//猜数字游戏
	var (
		v_count   int = 5 //定义计数器
		v_guest   int     //定义猜的数字
		v_scan    int     //定义输入变量
		v_guest_2 string  //定义重新猜数字的变量
	)
GUESTSTART: //开始生成随机数
	fmt.Println("开始生成随机数...")
	rand.Seed(time.Now().Unix())
	v_guest = rand.Intn(100) //100以内的随机数
	for {
		if v_count == 0 {
			fmt.Printf("你的机会已用完，答案是%d\n", v_guest)
			fmt.Println("你还想继续猜吗(Y/N)?")
			fmt.Scan(&v_guest_2)
			if strings.ToUpper(v_guest_2) == "Y" {
				fmt.Println("我佩服你的勇气，继续跳坑")
				v_count = 5 //重新初始化变量
				goto GUESTSTART //使用goto重新生成随机数
			} else {
				fmt.Println("你太笨了")
				break
			}
		}
		fmt.Println("随机数生成完毕，开始猜数字...")
		fmt.Println("请输入你的答案")
		fmt.Scan(&v_scan)
		if v_scan > v_guest {
			fmt.Println("太大了")
			v_count -= 1
			fmt.Printf("你还有%d次机会\n", v_count)
		} else if v_scan < v_guest {
			fmt.Println("太小了")
			v_count -= 1
			fmt.Printf("你还有%d次机会\n", v_count)
		} else {
			fmt.Println("猜中了")
			break
		}
	}
}
