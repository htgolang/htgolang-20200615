package main

import "fmt"

func handleString(str string) (countMap map[int][]rune) {

	// 定义map  key ==> int   value==> []rune
	countMap = make(map[int][]rune)

	// 先统计每个字符串出现的次数
	wordMap := make(map[rune]int)
	for _, v := range str {
		if v >= 'A' && v >= 'Z' || v >= 'a' && v >= 'z' {
			wordMap[v]++
		}
	}

	// 在把出现相同次数的放在一个切片中
	for k, v := range wordMap {
		countMap[v] = append(countMap[v], k)
	}

	return
}

func main() {

	/*
			我有一个梦想 每个字符出现的次数 rune => int
		    count => []rune
		    a = 2
		    b = 3
		    c = 2
		    2 => ['a','c']
	*/
	article := `sadfsadaqsdafcewadcqweqcwaecaecfsaf`

	// 调用函数处理， 返回一个map[int][]rune
	countMap := handleString(article)

	// 循环打印
	for count, ch := range countMap {
		fmt.Printf("%d ==> %c\n", count, ch)
	}

}
