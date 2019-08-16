package cmds

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Cat struct {
	Basedir string
}

type CatRequest struct {
	Path string
}

type CatResponse struct {
	Message string
	Content []byte
}

func (c *Cat) Exec(request *CatRequest, response *CatResponse) error {
	path := filepath.Join(c.Basedir, request.Path)
	fmt.Printf("cat path: %s\n", path)

	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			response.Message = "文件或目录不存在"
			fmt.Printf("[error] %s path is not found.\n", path)
			return errors.New("file is not found.")
		}
	} else {
		if info.IsDir() {
			response.Message = "不能查看目录"
			fmt.Printf("[error] %s is dir.\n", path)
			return errors.New("cat 不能查看目录")
		}
	}
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		response.Message = "读取文件失败"
		fmt.Printf("[error] readFile errors: %s\n", err)
		return errors.New("readFile errors.")
	} else {
		response.Content = bytes
	}
	return nil
}