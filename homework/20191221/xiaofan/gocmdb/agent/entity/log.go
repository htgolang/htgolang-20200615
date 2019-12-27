package entity

import (
	"encoding/json"
	"time"
)

const (
	LOGResource = 0X0001
)

type Log struct {
	UUID string    `json:"uuid"`
	Type int       `json:"type"`
	Msg  string    `json:"msg"`
	Time time.Time `json:"time"`
}

// 处理 entity/resource的信息后，将Log返回给ens
func NewLog(uuid string, typ int, msg interface{}) Log {
	bytes, _ := json.Marshal(msg)
	return Log{
		UUID: uuid,
		Type: typ,
		Msg:  string(bytes),
		Time: time.Now(),
	}
}
