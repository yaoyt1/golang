package main

import "fmt"

// 冒泡排序
func bubble_Sort(a [5]int) [5]int {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}

func main() {
	var i [5]int = [5]int{2, 4, 3, 1, 7}
	j := bubble_Sort(i)
	fmt.Println(i)
	fmt.Println(j)

}
