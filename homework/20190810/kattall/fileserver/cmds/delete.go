package cmds

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Delete struct {
	Basedir string
}

type DeleteRequest struct {
	Path string
}

type DeleteResponse struct {
	IsExist bool
	Message string
}

func (d *Delete) Exec(request *DeleteRequest, response *DeleteResponse) error {
	path := filepath.Join(d.Basedir, request.Path)
	fmt.Printf("delete %s\n", path)
	response.IsExist = true

	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			response.IsExist = false
			response.Message = "删除的文件或目录存在."
			fmt.Printf("[error] %s 不存在, 删除失败: %s", path, response.Message)
			return errors.New("delete path or file is not exist")
		}
	}

	if err := os.RemoveAll(path); err == nil {
		response.Message = "删除成功."
	} else {
		response.Message = "删除失败, 请查看错误信息."
	}
	return nil
}