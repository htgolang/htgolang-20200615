package entity

import (
	"time"
	"encoding/json"
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

func NewLog(uuid string, typ int, msg interface{}) Log {
	bytes, _ := json.Marshal(msg)
	return Log{
		UUID: uuid,
		Type: typ,
		Msg:  string(bytes),
		Time: time.Now(),
	}
}
