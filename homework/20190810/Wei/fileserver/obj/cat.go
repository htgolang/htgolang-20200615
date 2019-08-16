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
