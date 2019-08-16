package obj

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"time"
)

type FileInfo struct {
	Name  string
	Type  string
	Size  int64
	Ctime time.Time
	Mtime time.Time
}

type LsRequest struct {
	Path string
}

type LsResponse struct {
	FileInfos []FileInfo
}

type Ls struct {
	BaseDir string
}

func (l *Ls) Exec(request *LsRequest, response *LsResponse) error {
	path := filepath.Join(l.BaseDir, request.Path)
	log.Printf("[dubug] ls %s", path)

	file, err := os.Open(path)
	if err != nil {
		log.Print("[error] ls file error: ", err)
		return errors.New("打开路径错误")
	}

	defer file.Close()

	fileinfos, err := file.Readdir(-1)
	if err != nil {
		log.Print("[error] ls file error: ", err)
		return errors.New("读取目录错误")
	}

	response.FileInfos = make([]FileInfo, len(fileinfos))
	for i, fileinfo := range fileinfos {
		fileType := "f"
		if fileinfo.IsDir() {
			fileType = "d"
		}
		response.FileInfos[i] = FileInfo{
			Name:  fileinfo.Name(),
			Type:  fileType,
			Size:  fileinfo.Size(),
			Ctime: time.Now(),
			Mtime: fileinfo.ModTime(),
			Atime: time.Now(),
		}
	}
	return nil

}
