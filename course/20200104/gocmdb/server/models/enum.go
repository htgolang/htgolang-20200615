package models

const (
	StatusUnlock = iota
	StatusLock
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

const (
	TaskStatusNew = iota
	TaskStatusCancel
	TaskStatusScheduling
	TaskStatusExecing
	TaskStatusSuccess
	TaskStatusFailure
)