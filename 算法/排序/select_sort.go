package main

import "fmt"

// 插入排序
func select_Sort(a [5]int) [5]int {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[i] {
				a[j], a[i] = a[i], a[j]
			}
		}
	}
	return a
}

func main() {
	var i [5]int = [5]int{2, 4, 3, 1, 7}
	j := select_Sort(i)
	fmt.Println(i)
	fmt.Println(j)

}
