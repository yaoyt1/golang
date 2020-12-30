package main

import "fmt"

func test(num int) (res int) {
	if num <= 1 {
		res = 1
	} else {
		res = test(num-1) + test(num-2)
	}
	return
}

func main() {
	res := test(20)
	fmt.Println(res)
}
