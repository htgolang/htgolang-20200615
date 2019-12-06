package utils

import (
	"crypto/md5"
	"fmt"
	"strings"
)

// 默认生成8位随机的salt，之后对salt:randPassword进行md5加密，以salt:md5加密的密码这种形式进行存储
func Md5Salt(text, salt string) string {
	if salt == "" {
		salt = RandString(8)
	}
	return fmt.Sprintf("%s:%x", salt, md5.Sum([]byte(fmt.Sprintf("%s:%s", salt, text))))
}

// 对数据库中的密码进行拆分
func SplitMd5Salt(text string) (string, string) {
	nodes := strings.SplitN(text, ":", 2)
	if len(nodes) >= 2 {
		return nodes[0], nodes[1]
	} else {
		return "", nodes[0]
	}
}
