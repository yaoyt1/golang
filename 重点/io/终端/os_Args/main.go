package main

import (
	"fmt"
	"os"
)

//终端命令行参数
func main() {
	osArgs := os.Args
	fmt.Println("osArgs[0]=", osArgs[0])

	if len(osArgs) > 1 {
		for index, v := range osArgs {
			if index == 0 {
				continue
			}
			fmt.Printf("osArgs[%d]=%v\n", index, v)
		}

	}

}
