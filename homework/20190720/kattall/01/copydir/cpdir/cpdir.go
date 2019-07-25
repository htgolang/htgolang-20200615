package cpdir

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// 命令输入
func Input_flag() {
	// 设置日志文件
	logfile := "copy.log"
	src_dir := flag.String("s", "", "-s srcdir")
	dest_dir := flag.String("d", "", "-d destdir")
	help := flag.Bool("h", false, "help")
	flag.Usage = func() {
		fmt.Println(`
Usage: copydir.ext [-s srcdir] [-d destdir]

Options:
	`)
		flag.PrintDefaults()
	}
	flag.Parse()

	// 如果输入help或者 src和dest其中一个不输入, 都返回Usage
	if !*help || *src_dir != "" && *dest_dir != "" {
		cdir(*src_dir, *dest_dir, logfile)
		fmt.Printf("请查看日志文件：%s", logfile)
	} else {
		flag.Usage()
	}
}

// 拷贝目录  src_dir 源目录  dest目录
func cdir(src_dir, dest_dir, log_file string) {
	// 判断源目录是否存在, 不存在则退出
	if exist, _ := PathExists(src_dir);  !exist {
		fmt.Println("源目录不存在, 请确认原目录是否存在.")
		return
	}

	// 打开源目录
	src, err := os.Open(src_dir)
	if err != nil {
		fmt.Printf("源目录打开失败: %s\n", err)
	} else {
		// 循环源目录
		fileinfo, err := src.Readdir(-1)

		// 判断打开目录是否成功
		if err == nil {
			// 写入日志
			writeLog(log_file)

			// 打开成功, 创建dest_dir
			err := os.MkdirAll(dest_dir, 0644)
			if err != nil {
				log.Printf("创建%s目录失败：%s\n", dest_dir, err)
			} else {
				log.Printf("创建%s目录成功.\n", dest_dir)
			}
		} else {
			// 打开失败, 返回错误信息
			log.Printf("打开%s目录失败: %s\n", src_dir, err)
		}

		// 打开源目录后,循环, 判断是目录或者是文件
		for _, file :=  range  fileinfo {
			if file.IsDir() {
				// 如果是目录, 创建目录，然后继续调用此函数, 递归
				//fmt.Println("mkdir path:", filepath.Join(src_dir, file.Name()))
				err := os.MkdirAll(filepath.Join(src_dir, file.Name()), 0644)
				if err != nil {
					log.Printf("创建%s目录失败: %s\n", filepath.Join(src_dir,file.Name()), err)
				}
				// 递归
				cdir(filepath.Join(src_dir, file.Name()), filepath.Join(dest_dir, file.Name()), log_file)
			} else {
				// 如果是文件则拷贝文件
				copyfile(filepath.Join(src_dir, file.Name()), filepath.Join(dest_dir, file.Name()))
				log.Printf("%s 拷贝至 %s成功.\n", filepath.Join(src_dir, file.Name()), filepath.Join(dest_dir, file.Name()))
			}
		}
	}
}

// 拷贝文件
func copyfile(srcfile, destfile string) {
	// 如果是文件则打开文件, 开始复制
	readfile, err := os.Open(srcfile)
	if err != nil {
		fmt.Println(err)
	} else {
		defer readfile.Close()
		writefile, err := os.Create(destfile)
		if err != nil {
			fmt.Println(err)
		} else {
			defer writefile.Close()
			reader := bufio.NewReader(readfile)
			writer := bufio.NewWriter(writefile)
			bytes := make([]byte, 1024*1024*10)
			// 拷贝文件
			for {
				n ,err := reader.Read(bytes)
				if err != nil {
					if err != io.EOF {
						fmt.Println(err)
					}
					break
				} else {
					writer.Write(bytes[:n])
					writer.Flush()
				}
			}
		}
	}
}

// 写入日志
func writeLog(log_file string) {
	logfile, err := os.OpenFile(log_file, os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err == nil {
		log.SetOutput(logfile)
	}
}

// 判断目录或者文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}