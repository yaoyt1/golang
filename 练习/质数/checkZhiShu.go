package main

import (
	"fmt"
)

func checkZhiShu(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func testZhiShu() {
	sum := 0
	num := 100
	for i := 2; i < num; i++ {
		if checkZhiShu(i) {
			sum++
			defer fmt.Printf("[%d] 是质数\t\n", i)
		}
	}
	fmt.Printf("质数总数是【%d】个\n", sum)
}

func main() {
	testZhiShu()
}
