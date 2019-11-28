package utils

import (
	"fmt"
	"strings"
)
import "crypto/md5"


func Md5Salt(text string, salt string) string {

	if salt == "" {
		salt = RandString(8)
	}

	return fmt.Sprintf("%s:%x", salt, md5.Sum([]byte(fmt.Sprintf("%s:%x", salt, text))))


}

func SplitMd5Salt(text string) (string, string){

	nodes := strings.SplitN(text, ":", 2)

	if len(nodes) >= 2{
		return nodes[0], nodes[1]
	}else {
		return "", nodes[0]
	}

}