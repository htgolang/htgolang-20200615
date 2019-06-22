package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Compare("abc", "bac"))
	fmt.Println(strings.Contains("abc", "ad"))
	fmt.Println(strings.ContainsAny("abc", "ae"))
	fmt.Println(strings.Count("abcabcad", "ab"))
	fmt.Printf("%q\n", strings.Fields("abc def\neeee\raaaa\fxxxx\vsddddd")) // 按空白字符分割(空格, \n, \r, \f, \t)

	fmt.Println(strings.HasPrefix("abcd", "ab"))
	fmt.Println(strings.HasSuffix("abcdef", "defd"))
	fmt.Println(strings.Index("defabcdef", "dabc"))
	fmt.Println(strings.Index("defabcdef", "def"))
	fmt.Println(strings.LastIndex("defabcdef", "def"))

	fmt.Println(strings.Split("abcdef;abc;abc", ";"))
	fmt.Println(strings.Join([]string{"abc", "def", "eee"}, ":"))

	fmt.Println(strings.Repeat("abc", 3))
	fmt.Println(strings.Replace("abcabcabcab", "ab", "xxx", 1))
	fmt.Println(strings.Replace("abcabcabcab", "ab", "xxx", -1))
	fmt.Println(strings.ReplaceAll("abcabcabcab", "ab", "xxx"))

	fmt.Println(strings.ToLower("abcABC"))
	fmt.Println(strings.ToUpper("abcABC"))
	fmt.Println(strings.Title("hi, kk"))

	fmt.Println(strings.Trim("xyzabcyzx", "xz"))
	fmt.Println(strings.TrimSpace(" abcd xxx \n \r"))
}
