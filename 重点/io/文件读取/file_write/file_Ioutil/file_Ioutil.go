package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	err := ioutil.WriteFile("重点/io/文件读取/file_write/file_Ioutil/file_Ioutil.txt", []byte("1. 你好啊"), 0666)
	if err != nil {
		fmt.Println("ioutil 写入文件失败，错误如下：", err)
		return
	}
}
