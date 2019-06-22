package main

import "fmt"

func main() {
	var yes string
	fmt.Print("有卖西瓜的吗?(Y/N):")
	fmt.Scan(&yes)

	fmt.Println("老婆的想法:")
	fmt.Println("十个包子")
	if yes != "Y" && yes != "y" {
		goto END
	}

	fmt.Println("买一个西瓜")

END:

	result := 0
	i := 1

	//1...100
START:
	if i > 100 {
		goto FOREND
	}
	result += i
	i++
	goto START

FOREND:
	fmt.Println(result)

BREAKEND:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i*j == 9 {
				break BREAKEND
			}
			fmt.Println(i, j, i*j)
		}
	}

}
