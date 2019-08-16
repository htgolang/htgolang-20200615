package cmds

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type FileInfo struct {
	Name       string
	Type       string
	Size       int64
	CreateTime time.Time
	ModifyTime time.Time
	AccessTime time.Time
}

type Ls struct {
	Basedir string
}

type LsRequest struct {
	Path string
}

type LsResponse struct {
	FileInfos []FileInfo
}

func (l *Ls) Exec(request *LsRequest, response *LsResponse) error {
	path := filepath.Join(l.Basedir, request.Path)
	fmt.Printf("ls path: %s\n", path)

	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("[error] %s open failed.\n", path)
		return errors.New("open failed.")
	}
	defer file.Close()

	fileInfos, err := file.Readdir(-1)
	if err != nil {
		fmt.Printf("[error] 打开%s目录发生错误：%s\n", path, err)
	}

	response.FileInfos = make([]FileInfo, len(fileInfos))
	for i, fileInfo := range fileInfos {
		fileType := "F"
		if fileInfo.IsDir() {
			fileType = "D"
		}

		response.FileInfos[i] = FileInfo{
			Name:       fileInfo.Name(),
			Type:       fileType,
			Size:       fileInfo.Size(),
			CreateTime: time.Now(),
			ModifyTime: fileInfo.ModTime(),
			AccessTime: time.Now(),
		}
	}
	return nil
}