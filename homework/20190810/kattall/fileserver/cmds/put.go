package cmds

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Put struct {
	Basedir string
}

// 存放路径， 本地路径
type PutRequest struct {
	Path string
	SrcPath string
	Content []byte
}

// 判断目的是否存在,  message返回信息
type PutResponse struct {
	Message string
}

// 服务端
func (p *Put) Exec(request *PutRequest, response *PutResponse) error {
	path := filepath.Join(p.Basedir, request.Path)
	fmt.Printf("put %s to  %s\n", request.SrcPath, path)

	fileinfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		// 源不存在, 直接返回，报错 server path is not exist
		response.Message = "源目录不存在, 请检查"
		return errors.New("server path is not exist.")
	} else if !fileinfo.IsDir(){
		// 源如果是文件， 判断源文件和上传的文件是否一样， 一样就返回错误
		if fileinfo.Name() == filepath.Base(request.SrcPath) {
			response.Message = "server端存在此文件, 请重新put."
			return  errors.New("server已经存在此文件, 请重新put.")
		}
	} else if fileinfo.IsDir() {
		// 源如果是目录，循环目录下，是否有相同的文件，如果有相同的文件，就返回错误。
		file, _ := os.Open(path)
		if names, err := file.Readdir(-1); err == nil {
			for _, name := range names {
				if name.Name() == filepath.Base(request.SrcPath) {
					response.Message = "server端存在此文件, 请重新put."
					return  errors.New("server已经存在此文件, 请重新put.")
				}
			}
		}
		defer file.Close()
	}

	err = ioutil.WriteFile(filepath.Join(path, filepath.Base(request.SrcPath)), request.Content, os.ModePerm)
	if err == nil {
		response.Message = "put成功."
	} else {
		fmt.Println("put失败:", err)
	}
	return nil
}