package main

import "fmt"

// 插入排序
// 2 4 3 1 7
// 1 2 3 4 7
func insert_sort(a [5]int) [5]int {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
			} else {
				break
			}
		}
	}
	return a
}

func main() {
	var i = [5]int{2, 4, 3, 1, 7}
	j := insert_sort(i)
	fmt.Println(i)
	fmt.Println(j)
}
