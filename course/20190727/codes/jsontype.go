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
	default:5
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

func main() {
	var size Size = Medium

	bytes, _ := json.Marshal(size)

	fmt.Println(string(bytes))

	var size02 Size
	json.Unmarshal(bytes, &size02)

	fmt.Println(size02)

	sizes := []Size{Large, Large, Small, Medium}

	bytes, _ = json.Marshal(sizes)

	fmt.Println(string(bytes))

	var sizes02 []Size

	json.Unmarshal(bytes, &sizes02)

	fmt.Println(sizes02)
}
