package utils

import (
	"crypto/md5"
	"fmt"
)

func Md5(txt string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(txt)))
}
