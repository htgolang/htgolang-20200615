package main

import (
	"encoding/json"
	"fmt"
)

const (
	Large  = iota // large
	Medium        //medium
	Small         //small
)

type Size int

func (s Size) MarshalText() ([]byte, error) {
	switch s {
	case Large:
		return []byte("large"), nil
	case Medium:
		return []byte("medium"), nil
	case Small:
		return []byte("small"), nil
	default:
		return []byte("unknow"), nil
	}
}

func (s *Size) UnmarshalText(bytes []byte) error {
	switch string(bytes) {
	case "large":
		*s = Large
	case "medium":
		*s = Medium
	case "small":
		*s = Small
	default:
		*s = Small
	}
	return nil
}

type Addr struct {
	Region string `json:"region"`
	Street string `json:"street"`
	No     int    `json:"No"`
}

//1. 需要进行序列化或反序列化的属性必须公开
type User struct {
	ID   int    `json:"id,string"`
	Name string `json:"name"`
	Sex  int    `json:",,omitempty"`
	Tel  string `json:"-"`
	Addr Addr   `json:"addr"`
	Size Size   `json:"size"`
}

func main() {

	user := User{1, "kk", 1, "15200000000", Addr{"西安", "锦业路", 100}, Medium}

	bytes, _ := json.MarshalIndent(user, "", "\t")

	fmt.Println(string(bytes))

	var user02 User

	json.Unmarshal(bytes, &user02)

	fmt.Println(user02)
}
