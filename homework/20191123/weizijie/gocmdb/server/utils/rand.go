package utils

import (
	"math/rand"
	"time"
)

func RandString(length int) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"
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
