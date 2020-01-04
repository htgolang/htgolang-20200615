package models

const (
	StatusUnlock = 0
	StatusLock   = 1
)


const (
	AlarmTypeOffline = iota
	AlarmTypeCPU
	AlarmTypeRam
)

const (
	AlarmStatusNew = iota
	AlarmStatusDoing
	AlarmStatusComplete
)