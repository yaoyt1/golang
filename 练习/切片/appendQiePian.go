package main

import "fmt"

func main() {
	//test()
	testCopy()
}

func test() {
	veggies := []string{"yaoyoutian", "yaoyt"}
	fruites := []string{"yyt"}

	//这里的特殊语法：... 表示展开fruites切片成一个个元素
	food := append(veggies, fruites...)
	fmt.Println(food)
}

// 这里copy 的时候是根据长度来复制的
func testCopy() {
	veggies := []string{"yaoyoutian", "yaoyt"}
	fruites := []string{"yyt"}
	copy(veggies, fruites)

	fmt.Println(veggies)
	fmt.Println(fruites)
}
