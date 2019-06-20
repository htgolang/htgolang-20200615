/*
1、打印乘法口诀
要求，正阶梯的那种，左边乘数小

*/
package main

import "fmt"

func main() {

	/* 	i := 1
	   	for i <= 9 {
	   		for j := 1; j <= i; j++ {
	   			fmt.Printf("%d x %d = %d\t", j, i, i*j)
	   		}
	   		i++
	   		fmt.Println("\n")
	   	} */
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d x %d = %d\t", j, i, i*j)
		}
		fmt.Println("\n")
	}
}
