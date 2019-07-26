package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	src := flag.String("s", "", "src file")
	dest := flag.String("d", "", "dest file")
	recursive := flag.Bool("R", false, "copy directories recursively")
	help := flag.Bool("H", false, "help")
	flag.Parse()

	flag.Usage = func() {
		flag.PrintDefaults()
	}

	if *help || *src == "" || *dest == "" {
		flag.Usage()
	} else {
		switch *recursive {
		case false:
			copyFile(*src, *dest)
			return
		case true:
			break
		default:
			flag.Usage()
		}
	}

	filepath.Walk(*src, func(path string, fileInfo os.FileInfo, err error) error {
		relDir := filepath.Base(*src)
		if fileInfo.IsDir() {
			destDir := filepath.Join(*dest, relDir)
			os.MkdirAll(destDir, fileInfo.Mode())
		} else {
			destFile := filepath.Join(*dest, relDir, fileInfo.Name())
			copyFile(path, destFile)
		}
		return nil
	})
}

func copyFile(src, dest string) {
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

			bytes := make([]byte, 1024*10*10)

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
