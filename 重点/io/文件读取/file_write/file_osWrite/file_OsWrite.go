package main

import (
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

	f.Write([]byte("1. write: hello word\n"))
	f.WriteString("2. writeString: hello word")
	f.WriteAt([]byte("nihoaa "), 12) // 这里的写入是直接覆盖原有的  WriteAt()如果我们使用的是os.O_APPEND 打开文件的话， writeAt()函数报错
}
