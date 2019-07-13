package main

import (
	"fmt"
	"sort"
)

type User struct {
	ID    int
	Name  string
	Score float64
}

func main() {
	list := [][2]int{{1, 3}, {5, 9}, {4, 5}, {6, 2}, {5, 8}}

	//排序 使用数组的第二个（索引为1）元素比较大小进行排序
	sort.Slice(list, func(i, j int) bool {
		return list[i][1] < list[j][1]
	})

	fmt.Println(list)

	users := []User{{1, "kk", 6}, {2, "烟灰", 5}, {3, "俗人", 9}, {4, "小凡", 3}}

	sort.Slice(users, func(i, j int) bool {
		return users[i].Score < users[j].Score
	})
	fmt.Println(users)
}
