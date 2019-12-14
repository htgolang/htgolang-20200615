package utils

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func Md5Salt(text string, salt string) string {
	if salt == "" {
		salt = RandString(8)
	}
	return fmt.Sprintf("%s:%x", salt, md5.Sum([]byte(fmt.Sprintf("%s:%s", salt, text))))
}

func SplitMd5Salt(text string) (string, string) {
	nodes := strings.SplitN(text, ":", 2)
	if len(nodes) >= 2 {
		return nodes[0], nodes[1]
	} else {
		return "", nodes[0]
	}
}
