package config

import (
	//"fmt"
	"io/ioutil"
	"github.com/google/uuid"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	UUID string
	UUIDFile string  // UUID 文件
	Endpoint string  //
	Token string // 类似 accessKey\sercetKey

	LogFile string

	PID int   // 客户端ID
	PIDFile string // 存PID文件

	Heartbeat chan interface{} // 定义管道
	Register chan interface{} // 定义管道
	Log chan interface{}
}

func NewConfig() (*Config, error){

	UUIDFile := "agentd.uuid"
	PIDFile := "agentd.pid"

	LogFile := "logs/agent.log"

	UUID := ""
	if cxt, err := ioutil.ReadFile(UUIDFile); err ==nil{

		UUID = string(cxt)

	}else if os.IsNotExist(err){
		UUID = strings.ReplaceAll(uuid.New().String(), "-", "")
		ioutil.WriteFile(UUIDFile, []byte(UUID), os.ModePerm)
	}else {
		return nil, err
	}

	PID := os.Getpid()
	// int 转 字节切片
	ioutil.WriteFile(PIDFile, []byte(strconv.Itoa(PID)), os.ModePerm)

	//fmt.Println(UUID, PID)

	return &Config{
		Endpoint:"http://127.0.0.1:8080/v1/api",
		UUID: UUID,
		UUIDFile: UUIDFile,
		LogFile: LogFile,
		PID: PID,
		PIDFile: PIDFile,
		Heartbeat:make(chan interface{}, 64),
		Register: make(chan interface{}, 64),
		Log: make(chan interface{}, 10240),
	}, nil

}




