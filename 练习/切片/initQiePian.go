package main

import "fmt"

func test1() {
	var arr = [5]int{1, 2, 3, 4, 5}
	var p = arr[2:3]
	fmt.Printf("切片： %T \t %v \t len()=%d \t cap()=%d \n", p, p, len(p), cap(p))
	p = append(p, 6)
	fmt.Printf("切片： %T \t %v \t len()=%d \t cap()=%d \n", p, p, len(p), cap(p))

	for i := 0; i < 3; i++ {
		p = append(p, 6)
		fmt.Printf("切片： %T \t %v \t len()=%d \t cap()=%d \n", p, p, len(p), cap(p))
	}
}

func test2() {
	var a []int
	a = make([]int, 1, 3)
	fmt.Printf("切片： %T \t %v \t len()=%d \t cap()=%d \n", a, a, len(a), cap(a))
	a = append(a, 1)
	fmt.Printf("切片： %T \t %v \t len()=%d \t cap()=%d \n", a, a, len(a), cap(a))
}

func test3() {
	var names []string
	if names == nil {
		fmt.Println("切片是空的需要使用append函数添加")
		names = append(names, "yyt", "yaoyoutian", "yaoyt")
		fmt.Println(names)
	}

}

func main() {
	//test1()
	//test2()
	test3()
}
