package main

import "fmt"

//只有返回值，无输入参数
func testFunc() int {
	return 10
}

func main() {
	num := testFunc()
	fmt.Println("num = ", num)
}
