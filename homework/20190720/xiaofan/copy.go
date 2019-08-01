package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func copyFile(src, dest string) {
	source, err := os.Open(src)
	defer source.Close()
	destination, err := os.Create(dest)
	defer destination.Close()
	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewReader(source)
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println(err)
	}

	writer := bufio.NewWriter(destination)
	_, err = writer.Write(bytes)

	if err != nil {
		fmt.Println(err)
	}
	writer.Flush()
}

func copyDir(src, dest string) {
	files, err := ioutil.ReadDir(src)
	if err == nil {
		_ = os.Mkdir(dest, 0744)
		for _, file := range files {
			if file.IsDir() {
				copyDir(path.Join(src, file.Name()), path.Join(dest, file.Name()))
			} else {
				copyFile(path.Join(src, file.Name()), path.Join(dest, file.Name()))
			}
		}
	}
}

func main() {

	src := flag.String("s", "", "source file")
	dest := flag.String("d", "", "destination file")
	recu := flag.String("r", "", "recursive copy")
	flag.Usage = func() {
		fmt.Println(`Usage: copy -s srcfile -d destfile`)
		flag.PrintDefaults()
	}
	flag.Parse()
	if *src == "" || *dest == "" {
		flag.Usage()
	} else {
		if _, err := os.Stat(*dest); err == nil {
			fmt.Println("文件已存在")
			return
		} else {
			if !os.IsNotExist(err) {
				fmt.Println("获取文件信息失败", err)
				return
			}
		}

		if info, err := os.Stat(*src); err != nil {
			if os.IsNotExist(err) {
				fmt.Println("源文件不存在", err)
			} else {
				fmt.Println("获取文件信息失败", err)
			}
		} else {
			if info.IsDir() {
				if *recu == "" {
					fmt.Println("目录文件请使用 -r 参数")
				} else {
					copyDir(*src, *dest)
				}

			} else {
				copyFile(*src, *dest)
			}

		}

	}

}
