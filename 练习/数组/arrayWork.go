package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 数组所有元素之和
func arrayAllElementSum(array [5]int) (allElementSum int) {
	for _, v := range array {
		allElementSum += v
	}
	return
}

//计算下标之和
func twoSum(array [5]int, target int) {
	var otherNum int = 0
	for i := 0; i < len(array); i++ {
		otherNum = target - array[i]
		for j := i + 1; j < len(array); j++ {
			if array[j] == otherNum {
				fmt.Printf("下标[%d]，[%d]和为=[%d]  \t%d+%d=%d\n", i, j, target, array[i], array[j], target)
			}
		}
	}
}

func testArray() {
	var array [5]int

	//加入当前随机种子
	rand.Seed(time.Now().Unix())
	for i := 0; i < len(array); i++ {

		array[i] = rand.Intn(100)
	}

	fmt.Println(array)
	num := arrayAllElementSum(array)
	fmt.Println(num)

	var array1 [5]int = [5]int{10, 20, 30, 40, 50}
	twoSum(array1, 60)
}

func main() {
	testArray()
}
