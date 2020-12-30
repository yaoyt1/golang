package main

import "fmt"

//只有返回值，无输入参数
func testFunc() int {
	return 10
}

//传递变长参数函数
//找出最小的值
func testfunc1(arg ...int) int {
	if len(arg) == 0 {
		return 0
	}
	min := arg[0]
	for _, v := range arg {
		if v < min {
			min = v
		}
	}
	return min
}

func main() {
	num := testFunc()
	fmt.Println("1: num = ", num)

	var s []int = []int{1, 2, 4, 5}
	min := testfunc1(s...)
	fmt.Printf("2: 最小值是%d\n", min)

	min = testfunc1([]int{1, 2, 4, 5}...)
	fmt.Printf("3: 最小值是%d\n", min)

}
