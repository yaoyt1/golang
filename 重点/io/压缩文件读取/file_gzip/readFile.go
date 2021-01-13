package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("重点/io/压缩文件读取/测试.txt(1).gz")
	if err != nil {
		fmt.Println("os打开文件报错，报错内容如下：", err)
		return
	}
	defer f.Close()

	read, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println("gzip加载到阅读器中报错，报错内容如下：", err)
		return
	}
	defer read.Close()

	buf := bufio.NewReader(read) //加载到缓冲区
	for {
		context, err := buf.ReadString('\n')

		if err != nil {
			fmt.Println("读取缓冲区中的文件报错，报错内容如下：", err)
			os.Exit(0)
			return
		}
		fmt.Printf(string(context))
	}
}
