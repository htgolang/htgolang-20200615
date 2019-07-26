package main

import (
	"flag"
	"fmt"
	"os"
)
func copydir(src string,dest string)(err error){
	sourceinfo,err := os.Stat(src)
	if err != nil{
		return err
	}
	err = os.MkdirAll(dest,sourceinfo.Mode())
	if err != nil{
		return err
	}
	dir,err := os.Open(src)
	if err != nil{
		return err
	}
	obj,err := dir.Readdir(-1)
	for _,obj := range obj{
		src_splice := src + "/" + obj.Name()
		dst_splice := dest + "/" + obj.Name()
		if obj.IsDir(){			// 如果是目录则进行创建
			fmt.Println("[COPY Dir]:",src_splice,dst_splice)
			err = copydir(src_splice,dst_splice)
			if err != nil{
				fmt.Println(err)
			}
		}else {		// 如果是文件调用copyFile创建
			fmt.Println("[COPY File]:",src_splice,dst_splice)
			err =  copyFile(src_splice,dst_splice)
			if err != nil{
				fmt.Println(err)
			}
		}
	}
	return
}
func copyFile(src,dest string)(err error) {
	sourcefile,err := os.Open(src)
	if err != nil{
		return err
	}else {
		defer sourcefile.Close()
		destfile,err := os.Create(dest)
		if err != nil{
			fmt.Println(err)
		}else {
			defer destfile.Close()
			sourceinfo,err := os.Stat(src)
			if err != nil{
				err = os.Chmod(dest,sourceinfo.Mode())
			}
		}
	}
	return
}
func main(){
	src := flag.String("s","","Source directory")
	dest := flag.String("d","","Destination directory")
	help := flag.Bool("h",false,"help")
	flag.Usage = func() {
		go fmt.Print(`
Usage: Copy directory -s [Source directory] -d [Destination directory]
Options:
	`)
		flag.PrintDefaults()
	}
	flag.Parse()
	if *help || *src == "" || *dest == ""{

		flag.Usage()
	}else  {
		err := copydir(*src,*dest)
		if err !=nil{
			fmt.Println(err)
		}else {
			fmt.Printf("[%v]复制到[%v]完成。",*src,*dest)
		}

	}
}