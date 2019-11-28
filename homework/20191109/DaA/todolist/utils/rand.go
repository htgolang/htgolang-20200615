package utils

import (
	"math/rand"
	"time"
)

func RandString(length int) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWSYZ0123456789"
	len := len(letters)
	chars := make([]byte,length)
	for i := 0; i < length ; i++ {
		chars[i] = letters[rand.Int()%len]
	}
	return string(chars)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}