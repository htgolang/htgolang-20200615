package newcopy

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func CopyFile(src, dest string){
	srcfile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
	} else {
		defer srcfile.Close()
		destfile, err := os.Create(dest)
		if err != nil {
			fmt.Println(err)
		} else {
			defer destfile.Close()

			reader := bufio.NewReader(srcfile)
			writer := bufio.NewWriter(destfile)

			bytes := make([]byte, 1024*1024*10)

			for {
				n, err := reader.Read(bytes)
				if err != nil {
					if err != io.EOF {
						fmt.Println(err)
					}
					break
				}
				writer.Write(bytes[:n])
				writer.Flush()
			}
		}
	}
}

func CopyDir(src, dest string)  {
	files, err := ioutil.ReadDir(src)
	if err == nil {
		for _, file := range files {
			if file.IsDir() {
				CopyDir(filepath.Join(src, file.Name()), filepath.Join(dest, file.Name()))
			}else {
				CopyFile(filepath.Join(src, file.Name()), filepath.Join(dest, file.Name()))
			}
		}
	}


}
