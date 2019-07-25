package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func Dirlist(path string) ([]string, []string) {
	fileinfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("%s 目录不存在\n", path)
			fmt.Println("目录不存在")
			return []string{}, []string{}
		}
		return []string{}, []string{}
	} else {
		if fileinfo.IsDir() {
			files := []string{}
			dirs := []string{}
			dirfile, err := os.Open(path)
			if err == nil {
				defer dirfile.Close()
				childrens, _ := dirfile.Readdir(-1)
				for _, children := range childrens {
					if !children.IsDir() {
						files = append(files, children.Name())
					} else if children.IsDir() {
						dirs = append(dirs, children.Name())
					}
				}
				return files, dirs
			}
			return []string{}, []string{}
		}
		return []string{}, []string{}
	}

}

func dircopy(spath, dpath string) {
	var destdir_path string
	var srcdir_path string

	sfile, sdir := Dirlist(spath)
	for _, file := range sfile {
		spath := filepath.Join(spath, file)
		dpath := filepath.Join(dpath, file)
		copyfile(spath, dpath)
	}

	if len(sdir) == 0 {
		goto END
	} else {
		for _, dir := range sdir {
			destdir_path = filepath.Join(dpath, dir)
			srcdir_path = filepath.Join(spath, dir)
			os.Mkdir(destdir_path, 0644)
			log.Printf("%s 目录已创建\n", destdir_path)
			dircopy(srcdir_path, destdir_path)
		}
	}

END:
}

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

			// 每次拷贝10M字节
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
			log.Printf("%s 文件已复制\n", src)

		}
	}
}

func main() {
	logfile := "copy.log"
	file, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE, os.ModePerm)

	if err == nil {
		log.SetOutput(file)
		commandfile := os.Args[0]
		log.SetPrefix(commandfile + ":")
		log.SetFlags(log.Flags() | log.Lshortfile)
	}

	src := flag.String("s", "", "src file")
	dest := flag.String("d", "", "dest file")
	dir := flag.Bool("R", false, "Dir file")
	help := flag.Bool("h", false, "help")

	flag.Usage = func() {
		fmt.Println(`
Usage: copyfile [-R] -s srcfile -d destfile
Optime: 
		`)
		flag.PrintDefaults()
	}

	flag.Parse()

	if *help || *src == "" || *dest == "" {
		flag.Usage()
	} else if !*dir {
		fileinfo, _ := os.Stat(*src)
		if !fileinfo.IsDir() {
			copyfile(*src, *dest)
		} else {
			log.Printf("%s 是目录文件，需使用-R复制\n", *src)
			fmt.Printf("%s 是个目录，需使用-R复制", *src)
		}

	} else if *dir {
		fileinfo, _ := os.Stat(*src)
		//spath := filepath.Join("D:/Code/goang/practise/day06", *src)
		//dpath := filepath.Join("D:/Code/goang/practise/day06", *dest)
		spath := os.Args[3]
		dpath := os.Args[5]

		if fileinfo.IsDir() {
			os.Mkdir(dpath, 0644)
			log.Printf("%s 目录已创建\n", dpath)

			dircopy(spath, dpath)
		}
	}
	file.Close()

}
