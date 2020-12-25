package main

import (
	"fmt"
	"strings"
	"time"
)

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// 匿名函数传参
func testClosure() {
	for i := 0; i < 5; i++ {
		go func(index int) {
			fmt.Printf("%d \n", index)
		}(i)
	}
	time.Sleep(time.Second)
}

func main() {
	func1 := makeSuffixFunc(".bmp")
	func2 := makeSuffixFunc(".jpg")

	fmt.Printf(func1("test"))
	fmt.Println()
	fmt.Printf(func2("test"))
	fmt.Println()
	testClosure()
}
