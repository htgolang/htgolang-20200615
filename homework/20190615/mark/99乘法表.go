package main
/*
九九乘法表
 */
import "fmt"

func jiujiu()  {
	for i :=0; i < 9; i ++ {
		for j :=0; j <= i; j++ {
			fmt.Printf("%d*%d=%d ",(i + 1),j + 1, (i+1)*(j+1))
		}
		fmt.Println()
	}
}
func main()  {
	jiujiu()
}
