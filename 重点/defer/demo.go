package main

import "fmt"

/*
defer 推迟执行： 类似于栈 后进先出
作用：
1. 用于释放某些资源
2. 实现代码追踪
3. 记录函数的参数和返回值

return 原理：
1. 先设置返回值
2. 返回

defer 原理：
1. 先设置返回值
2. 执行defer中函数
3. 返回

*/

func deferA() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func deferB() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func deferC() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func deferD() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func main() {
	var num int
	num = deferA()
	fmt.Println("deferA：", num) //5

	num = deferB()
	fmt.Println("deferB：", num) //6

	num = deferC()
	fmt.Println("deferC：", num) //5

	num = deferD()
	fmt.Println("deferD：", num) //5

}
