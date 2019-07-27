package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	/*
		json.Marshal 序列化 内存=>字符串/字节切片
		json.Unmarshal 反序列化 字符串/字节切片 => 内存
	*/

	names := []string{"未子杰", "祥哥", "小凡", "xq"}

	users := []map[string]string{{"name": "未子杰", "addr": "上海"}, {"name": "祥哥", "addr": "杭州"}, {"name": "小凡", "addr": "北京"}}

	bytes, err := json.MarshalIndent(names, "", "\t")
	if err == nil {
		// fmt.Println(bytes)
		fmt.Println(string(bytes))
	}

	var names02 []string

	err = json.Unmarshal(bytes, &names02)
	fmt.Println(err)
	fmt.Println(names02)

	bytes, err = json.MarshalIndent(users, "", "\t")
	if err == nil {
		// fmt.Println(bytes)
		fmt.Println(string(bytes))
	}

	var user02 []map[string]string

	err = json.Unmarshal(bytes, &user02)
	fmt.Println(err)
	fmt.Println(user02)

	fmt.Println(json.Valid([]byte("[]x")))

}
