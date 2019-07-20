package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("xxxx")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("文件不存在")
		}

	} else {
		file.Close()
	}

	for _, path := range []string{"xxx", "reader.go", "usermanager"} {
		fileInfo, err := os.Stat(path)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println("文件不存在")
			}
		} else {
			fmt.Println(strings.Repeat("*", 20))
			fmt.Println(fileInfo.Name())
			fmt.Println(fileInfo.IsDir())
			fmt.Println(fileInfo.Size())
			fmt.Println(fileInfo.ModTime())

			if fileInfo.IsDir() {
				dirfile, err := os.Open(path)
				if err == nil {
					defer dirfile.Close()

					// childrens, _ := dirfile.Readdir(-1)
					// for _, children := range childrens {
					// 	fmt.Println(children.Name(), children.IsDir(), children.Size(), children.ModTime())
					// }

					names, _ := dirfile.Readdirnames(-1)
					for _, name := range names {
						fmt.Println(name)
					}
				}
			}
		}
	}

}
