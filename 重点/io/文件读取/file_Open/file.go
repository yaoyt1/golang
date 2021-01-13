package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("重点/io/文件读取/file_Open/file.go")
	if err != nil {
		fmt.Printf("读取文件报错， 错误如下：%v\n", err)
		return
	}
	defer file.Close()

	var allBuf []byte
	var buf [128]byte

	for {
		n, err := file.Read(buf[:])

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("读取文件报错， 错误如下：", err)
			return
		}

		allBuf = append(allBuf, buf[:n]...)
	}

	fmt.Println("读取文件成功")
	fmt.Printf("%s\n", string(allBuf))
}
