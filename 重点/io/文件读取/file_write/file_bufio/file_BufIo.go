package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile("重点/io/文件读取/file_write/file_osWrite/test.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	//这里第三位参数是linux的权限控制   第三位参数是个三位数  （非同组的人员，同组其他人员， 我）
	//参数数字代表 r(只读)--004
	//w--002
	//x--001
	//权限控制是三个数加起来
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
