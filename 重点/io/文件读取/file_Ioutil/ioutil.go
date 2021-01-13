package main

import (
	"fmt"
	"io/ioutil"
)

//这里他调用的还是os底层， 只是给我们固定好了每次读取的字节大小
func main() {
	content, err := ioutil.ReadFile("重点/io/文件读取/file_Ioutil/ioutil.go")
	if err != nil {
		fmt.Println("ioutil读取文件报错", err)
		return
	}
	fmt.Println(string(content))
}
