package main

import (
	"fmt"
	"sort"
)

// 切片的简单使用
func test() {
	var sa = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		sa = append(sa, fmt.Sprintf("%v", i))
	}
	fmt.Println(len(sa), cap(sa), sa)
}

//使用golang中的标准包sort进行排序
func test1() {
	a := [8]int{1, 3, 2, 5, 4}
	sort.Ints(a[:])
	fmt.Println("升序a=", a)
	sort.Sort(sort.Reverse(sort.IntSlice(a[:])))
	fmt.Println("降序a=", a)

}

func main() {
	//test()
	test1()
}
