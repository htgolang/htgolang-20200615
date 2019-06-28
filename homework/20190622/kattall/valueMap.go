package main

import "fmt"

func gainValue(kmap map[string]string) (valueSlice []string) {
	// 循环map, 并且把value加入到切片中
	for _, v := range kmap {
		valueSlice = append(valueSlice, v)
	}

	// return是返回值, 默认返回 keySlice
	return
}

func main() {
	/*
		获取映射中所有value组成的切片
	*/

	// 定义一个map
	keyMap := map[string]string{"G1001": "祥哥", "G1002": "宝成", "G1003": "小凡", "G1004": "子杰", "G1005": "大A"}
	fmt.Println("原始的map为：", keyMap)

	// 调用函数 gainKey  获取返回值并且赋值为 keySlice
	valueSlice := gainValue(keyMap)
	fmt.Println("获取映射中所有value组成的切片: ", valueSlice)
}

/*
 评分: 8
*/
