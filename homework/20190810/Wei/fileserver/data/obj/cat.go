package obj

import (
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
)

type CatRequest struct {
	Path string
}

type CatResponse struct {
	Bytes []byte
}

type Cat struct {
	BaseDir string
}

func (c *Cat) Exec(request *CatRequest, response *CatResponse) error {
	// binPath, _ := filepath.Abs(os.Args[0])
	// logName := "log.txt"
	// logPath := filepath.Join(filepath.Dir(binPath), logName)
	// log_file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE, os.ModePerm)
	// if err == nil {
	// 	log.SetOutput(log_file)
	// 	commandfile := os.Args[0]
	// 	log.SetPrefix(commandfile + ":")
	// 	log.SetFlags(log.Flags() | log.Lshortfile)
	// }

	path := filepath.Join(c.BaseDir, request.Path)
	log.Printf("[info] cat %s", path)

	file, err := os.Open(path)
	if err != nil {
		log.Print("[error] file open error: ", path)
		return errors.New("打开文件错误")
	}

	defer file.Close()

	cxt := make([]byte, 1024)
	n, err := file.Read(cxt)
	if err != nil && err != io.EOF {
		log.Print("[error] file read file: ", path)
		return errors.New("读取文件错误")
	}

	response.Bytes = cxt[:n]
	return nil
}
