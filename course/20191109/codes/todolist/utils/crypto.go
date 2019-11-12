package utils

import (
	"crypto/md5"
	"fmt"
)

// 计算md5值
func Md5(txt string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(txt)))
}
