package inputstring

import (
	"fmt"
	"strings"
)

// 定义从键盘输入函数，并返回输入的值
func InputString(s string) string {
	var in string
	fmt.Print(s)
	fmt.Scan(&in)
	return strings.TrimSpace(in)
}
