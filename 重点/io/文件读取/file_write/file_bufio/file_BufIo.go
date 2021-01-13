package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile("重点/io/文件读取/file_write/file_osWrite/test.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("写入文件失败，错误内容如下：", err)
		return
	}
	defer f.Close()

	//加入缓存区
	wt := bufio.NewWriter(f)

	wt.WriteString("1. writeString: hello word")
	wt.Flush() //缓存区刷新写入文件
}
