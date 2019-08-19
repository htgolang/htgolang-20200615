package models

import (
	"crypto/md5"
	"fmt"
)

var Cookies = make(map[string]string)

func SetCookie(userName, remoteAddr, userAgent string) {
	key := fmt.Sprintf("%x", md5.Sum([]byte(remoteAddr+userAgent)))
	Cookies[key] = userName
}

func GetCookie(remoteAddr, userAgent string) string {
	key := fmt.Sprintf("%x", md5.Sum([]byte(remoteAddr+userAgent)))
	if username, ok := Cookies[key]; ok {
		return username
	}
	return ""
}
