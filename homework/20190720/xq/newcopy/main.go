package main

import (
	"flag"
	"fmt"
	"github.com/xlotz/newcopy"
	"os"
)


func main() {


	src := flag.String("s", "", "src file")
	dest := flag.String("d", "", "dest file")
	help := flag.Bool("h", false, "help")

	flag.Usage = func() {
		fmt.Println(`
Usage: newcopy -s srcfile -d destfile
Optime:
		`)
		flag.PrintDefaults()
	}

	flag.Parse()

	if *help || *src == "" || *dest == "" {
		flag.Usage()
	} else {
    // src 是不是存在， 不存在 退出
    // src 是文件 copyfile
    // src 是目录 copydir

    // dest 判断 存在 退出

		if _, err := os.Stat(*dest); err == nil {

			fmt.Println("dest is exist")
			return
		}else {
			if !os.IsNotExist(err){
				fmt.Println("dest file error")
			}
		}


		if info, err := os.Stat(*src); err != nil {
			if os.IsNotExist(err){
				fmt.Println("Not find file")
			}else {
				fmt.Println("error")
			}
		}else {
			if info.IsDir(){
				newcopy.CopyDir(*src, *dest)
			}else {
				newcopy.CopyFile(*src, *dest)
			}
		}
	}


}
