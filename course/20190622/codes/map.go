package main

import "fmt"

func main() {
	var scores map[string]int // nil 映射

	fmt.Printf("%T %#v\n", scores, scores)
	fmt.Println(scores == nil)
	// 字面量

	scores = make(map[string]int)
	fmt.Println(scores)

	// scores = map[string]int{}
	scores = map[string]int{"武鹏飞": 8, "烟灰": 9, "祥哥": 10}
	fmt.Println(scores)

	// 增，删，改，查
	// key
	fmt.Println(scores["武鹏飞"])
	fmt.Println(scores["保成"])

	if v, ok := scores["武鹏飞"]; ok {
		fmt.Println(v)
	}
	if v, ok := scores["保成"]; ok {
		fmt.Println(v)
	}

	scores["烟灰"] = 8
	fmt.Println(scores)
	scores["保成"] = 6
	fmt.Println(scores)

	// 宝成
	delete(scores, "保成")
	fmt.Println(scores)
	scores["宝成"] = 6
	fmt.Println(scores)

	fmt.Println(len(scores))
	for k, v := range scores {
		fmt.Println(k, v)
	}

	//key至少可以有==,!= 运算, bool, 整数, 字符串, 数组
	// value => 任意类型 slice map
	//名字 => 映射[字符串]字符串{"地方"，"联系方式", "成绩"}

	var users map[string]map[string]string

	users = map[string]map[string]string{"烟灰": {"地方": "福建", "成绩": "8", "联系方式": "123456789"}}

	fmt.Printf("%T, %#v\n", users, users)

	_, ok := users["武鹏飞"]
	fmt.Println(ok)
	users["武鹏飞"] = map[string]string{"地方": "北京"}
	fmt.Println(users)
	users["武鹏飞"]["成绩"] = "9"
	fmt.Println(users)
	delete(users["烟灰"], "联系方式")
	fmt.Println(users)
}
