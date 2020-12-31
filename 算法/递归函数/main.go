package main

import (
	"fmt"
	"time"
)

const fiBonaCCiNum int = 80

var fibos [fiBonaCCiNum + 1]uint64

func fiBoNaCCi(num int) (res uint64) {
	if num <= 1 {
		res = 1
	} else {
		res = fiBoNaCCi(num-1) + fiBoNaCCi(num-2)
	}
	return
}

func fiBoNaCCi1(num int) (res uint64) {
	if fibos[num] != 0 {
		res = fibos[num]
		return
	}

	if num <= 1 {
		res = 1
	} else {
		res = fiBoNaCCi(num-1) + fiBoNaCCi(num-2)
	}
	fibos[num] = res
	return
}

func main() {
	start := time.Now()
	res := fiBoNaCCi(fiBonaCCiNum)
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("未使用缓存优化： %s \t  %d的斐波那契数列：%d\n", delta, fiBonaCCiNum, res)

	start1 := time.Now()
	res1 := fiBoNaCCi1(fiBonaCCiNum)
	end1 := time.Now()
	delta1 := end1.Sub(start1)
	fmt.Printf("使用缓存优化： %s \t  %d的斐波那契数列：%d\n", delta1, fiBonaCCiNum, res1)

}
