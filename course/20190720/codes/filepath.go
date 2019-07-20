package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Abs("."))
	fmt.Println(os.Args)

	path, _ := filepath.Abs(os.Args[0])
	dirPath := filepath.Dir(path)

	fmt.Println(filepath.Base("c:/test/a.txt"))
	fmt.Println(filepath.Base("c:/test/xxxx/"))
	fmt.Println(filepath.Base(path))

	fmt.Println(filepath.Dir("c:/test/a.txt"))
	fmt.Println(filepath.Dir("c:/test/xxxx/"))
	fmt.Println(filepath.Dir(path))

	fmt.Println(filepath.Ext("c:/test/a.txt"))
	fmt.Println(filepath.Ext("c:/test/xxxx/a"))
	fmt.Println(filepath.Ext(path))

	iniPath := dirPath + "/conf/ip.ini"
	fmt.Println(iniPath)
	fmt.Println(filepath.Join(dirPath, "conf", "ip.ini"))

	fmt.Println(filepath.Glob("./[ab]*/*.go"))

	filepath.Walk(".", func(path string, fileInfo os.FileInfo, err error) error {
		fmt.Println(path, fileInfo.Name())
		return nil
	})

}
