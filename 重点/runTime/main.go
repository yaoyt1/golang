package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpu := runtime.NumCPU()
	fmt.Println("cpu个数：", cpu)
	runtime.GOMAXPROCS(cpu)
	var i int = 0
	for {
		i++
	}
}
