package utils

import (
	"crypto/md5"
	"fmt"
)

func Md5New(txt string)string{
	// hash := md5.New()
	// fmt.Fprintf(hash, txt)
	// return fmt.Sprintf("%x", hash.Sum(nil))

	return fmt.Sprintf("%x", md5.Sum([]byte(txt)))
}
