/*
终端带缓冲区的读写
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入：")

	//这里的意思是读取到\n 字符串就返回不读取了
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("程序错误")
		return
	}
	fmt.Printf("终端输入：%s\n", input)
}
