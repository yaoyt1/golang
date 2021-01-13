package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, _ := os.Open("重点/io/文件读取/file_BufIo/bufio.go")
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		//这里是单独读取一行 故所以我们要加入for循环
		readerStr, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		fmt.Println(readerStr)
	}
}
