package GoStudynotes

import (
	"fmt"
	"strings"
)
func QueryUser(users map[int]map[string]string) {
	q := inputString("请输入查询内容:")
	for k,v := range users {
		if strings.Contains(v["name"], q) || strings.Contains(v["tel"], q) || strings.Contains(v["age"], q) || strings.Contains(v["addr"], q) {
			printUser(k,v)
			fmt.Println("`````````````````````****`````````````````````````````")
		}else {
			fmt.Printf("[err]并没有找到%v的相关内容\n",q)
		}
	}
}