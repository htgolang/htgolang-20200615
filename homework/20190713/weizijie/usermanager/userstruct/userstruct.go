package userstruct

import "time"

// 定义用户结构体类型Users

type Users struct {
	ID       int
	Name     string
	Birthday time.Time
	Addr     string
	Tel      string
	Desc     string
}

// 定义用户变量

var User map[int]Users = map[int]Users{}
