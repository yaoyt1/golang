package main

import "fmt"

func addr() func(int) int {
	var x int
	return func(d int) int {
		x += d
		return x
	}
}

func main() {
	f := addr()
	ret := f(1)
	fmt.Printf("f(1)=%d\n", ret)
	ret = f(10)
	fmt.Printf("f(100)=%d\n", ret)
	ret = f(100)
	fmt.Printf("f(100)=%d\n", ret)
	ret = f(1000)
	fmt.Printf("f(1000)=%d\n", ret)
	fmt.Printf("重新初始化\n")
	f = addr()
	ret = f(1)
	fmt.Printf("f(1)=%d\n", ret)
}
