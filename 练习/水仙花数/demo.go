package main

import "fmt"

func example1() {
	num := 1000
	for i := 100; i < num; i++ {
		if isShuiXianHuaShu(i) {
			fmt.Printf("[%d] 是水仙花数\t\n", i)
		}
	}
}

func isShuiXianHuaShu(num int) bool {
	first := num % 10
	second := (num / 10) % 10
	third := (num / 100) % 10

	num1 := first*first*first + second*second*second + third*third*third
	if num1 == num {
		return true
	}
	return false
}

func main() {
	example1()
}
