package utils

import (
	"math/rand"
	"time"
)

// 随机数对letters的长度取余，循环length遍，生成随机密码
func RandString(length int) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890-_"
	count := len(letters)
	chars := make([]byte, length)

	for i := 0; i < length; i++ {
		chars[i] = letters[rand.Int()%count]
		// rand.Intn(count)
	}
	return string(chars)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
