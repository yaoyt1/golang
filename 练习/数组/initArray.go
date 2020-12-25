// 数组长度是数组类型的一部分
package main

import "fmt"

// 初始化
// 然后赋值
func exmaple1() {
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println(a)
}

// 创建变量的时候赋值
func exmaple2() {
	var a [3]int = [3]int{1, 2, 3}
	fmt.Println(a)
}

// 类型推导式
func exmaple3() {
	a := [3]int{1, 2, 3}
	fmt.Println(a)
}

// 类型推导式 不声明数组长度
func exmaple4() {
	a := [...]int{1, 2, 3}
	fmt.Println(a)
}

// 赋值一部分， 其余部分的使用类型默认值
func exmaple5() {
	a := [3]int{1}
	fmt.Println(a)
}

// 使用下标赋值
func exmaple6() {
	a := [3]int{1: 100, 2: 300}
	fmt.Println(a)
}

func main() {
	exmaple1()
	exmaple2()
	exmaple3()
	exmaple4()
	exmaple5()
	exmaple6()
}
