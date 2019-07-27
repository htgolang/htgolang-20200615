package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func copyfile(src, dest string) {
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

func copydir(src, dst string) {
	files, err := ioutil.ReadDir(src)
	if err == nil {
		for _, file := range files {
			if file.IsDir() {
				copydir(filepath.Join(src, file.Name()), filepath.Join(dst, file.Name()))
			} else {
				copyfile(filepath.Join(src, file.Name()), filepath.Join(dst, file.Name()))
			}
		}
	}
}

func main() {
	src := flag.String("s", "", "src file")
	dest := flag.String("d", "", "dest file")
	help := flag.Bool("h", false, "help")

	flag.Usage = func() {
		fmt.Println(`
Usage: copyfile -s srcfile -d destfile
Options:
		`)
		flag.PrintDefaults()
	}

	flag.Parse()

	if *help || *src == "" || *dest == "" {
		flag.Usage()
	} else {
		// src 是不是存在, 不存在 退出
		// src 文件 copyfile
		// src 目录 copydir

		// dst 判断 存在 退出

		if _, err := os.Stat(*dest); err == nil {
			fmt.Println("目的文件已存在")
			return
		} else {
			if !os.IsNotExist(err) {
				fmt.Println("目的文件获取错误 ", err)
			}
		}

		if info, err := os.Stat(*src); err != nil {
			if os.IsNotExist(err) {
				fmt.Println("源文件不存在")
			} else {
				fmt.Println("源文件获取错误: ", err)
			}
		} else {
			if info.IsDir() {
				copydir(*src, *dest)
			} else {
				copyfile(*src, *dest)
			}
		}

	}
}
