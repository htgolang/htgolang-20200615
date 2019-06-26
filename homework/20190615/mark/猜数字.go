package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)
func dadata(sjs int) int{
	rand.Seed(time.Now().UnixNano())   // 系统时间为随机种子
	return rand.Intn(sjs)
}
func main(){
	sjs := dadata(100)
	fmt.Println(sjs)
	num := 3
	var dA1 string

	for i:=0;i<=3;i++{
		rnum := num - i
		fmt.Println("开始猜数字游戏，请输入:")
		fmt.Scanf("%v",&dA1)
		dA1,err := strconv.ParseInt(dA1,10,0)  // 判断输入的是整数
		if err != nil {
			fmt.Println("你输入的不是数字",err)
			break
		}else if int(dA1) == sjs {
			fmt.Printf("恭喜你猜对了数字是%v!", sjs)
			goto ENDD
		} else if int(dA1) > sjs {
			fmt.Printf("猜大了,你还可以猜测%v次\n",rnum)
		} else if int(dA1) < sjs {
			fmt.Printf("猜小了,你还可以猜测%v次\n",rnum)
		}
	}
ENDD:
}