package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("cpu个数：", runtime.NumCPU())
	for i := 0; i < 10; i++ {
		go helloWord(i)
	}
	time.Sleep(time.Second)
}

func helloWord(i int) {
	fmt.Println("新开一个线程，i=", i)
}
