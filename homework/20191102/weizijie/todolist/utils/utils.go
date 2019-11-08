package utils

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

func RandomString(length int) string{
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	len := len(letters)

	chars := make([]byte, length)

	for i := 0; i < length; i++ {
		chars[i] = letters[rand.Int() % len]
	}
	return string(chars)
}

func Md5String(plain string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(plain)))
}

func init() {
	rand.Seed(time.Now().Unix())
}