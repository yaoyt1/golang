//字典包
/*
1 map基础知识点
1.1 未初始化的 map 的值是 nil。
1.2 所以数组、切片和结构体不能作为 key
(译者注：含有数组切片的结构体不能作为 key，只包含内建类型的 struct 是可以作为 key 的），
	但是指针和接口类型可以。如果要用结构体作为 key 可以提供 `Key()` 和 `Hash()` 方法，
	这样可以通过结构体的域计算出唯一的数字或者字符串的 key。

value 可以是任意类型

1.3 go ebook 说不要使用new,永远使用make来创建字典

2 map容量
当达到map容量最大值的时候， map的大小会自动+1
出于性能要求： 我们处理大的map或者需要快速扩容的map尽量先标明容量

3 map 循环输出不是有序的是乱序

4 map类型的切片

*/
package main

import "fmt"

func test() {
	var a map[string]int
	if a == nil {
		fmt.Println("1.1 未初始化map的值是nil")
	} else {
		fmt.Println("1.1 未初始化map的值不是nil")
	}

	mp1 := make(map[int][]int)
	mp1[0] = []int{1, 2, 3}
	mp1[0] = func() []int {
		return []int{1, 2, 3}
	}()
	fmt.Printf("1.2 map %v value 值可以为任何类型\n", mp1)

	newMap := new(map[string]int)
	//newMap["yyt"] = 0
	//这里赋值编译器会报错
	//new出来的这里我们不能直接这样赋值 new 出来的是直接把地址指向一块空地址
	fmt.Printf("1.3 map new初始化 的值%v\t地址%p\n", newMap, &newMap)

}

//字典初始化
func mapInit() {
	//正常赋值
	var a map[string]int
	a = make(map[string]int, 5)
	if a == nil || len(a) == 0 {
		a["yaoyoutian"] = 22
		a["yyt"] = 23
	}

	fmt.Printf("字典%v \n", a)

	//直接声明字典的时候赋值
	var b = map[string]int{
		"yyt":        10,
		"yaoyoutian": 20,
	}
	fmt.Printf("字典%v \n", b)

}

/* 字典的基本操作
1. 判断字典是否存在
2. 删除元素

*/
func mapBasic() {
	map1 := map[string]int{"y": 1, "yyt": 2}

	// 1. 判断字典是否存在
	isKey := "yy"
	_, ok := map1[isKey]
	if ok == true {
		fmt.Printf("字典%vkey中有%s \n", map1, isKey)
	} else {
		fmt.Printf("字典%vkey中无%s \n", map1, isKey)
	}

	if _, ok1 := map1[isKey]; ok1 {
		fmt.Printf("if混合使用 字典%vkey中有%s \n", map1, isKey)
	} else {
		fmt.Printf("if混合使用 字典%vkey中无%s \n", map1, isKey)
	}

	//删除元素
	delete(map1, "y")
	fmt.Printf("map删除元素 字典%v 正确的删除字典key \n", map1)

	delete(map1, "yaoyoutian")
	fmt.Printf("map删除元素 字典%v 错误的删除字典key，但是程序无报错 \n", map1)
}

/*map类型的切片
这里我们必须使用两次make 函数来分配
第一次分配切片
第二次分配map

*/
func mapAdvanced() {
	//版本1循环
	items := make([]map[string]int, 5)
	for index := range items {
		//这里我们需要make map
		items[index] = make(map[string]int)
		items[index]["yyt"] = 2
	}

	fmt.Printf("4 map类型的切片使用索引来make分配：%v\n", items)

	//版本2循环
	items1 := make([]map[string]int, 5)
	for _, item := range items1 {
		//这里我们需要make map
		item = make(map[string]int)
		item["yyt"] = 2
	}
	fmt.Printf("4 map类型的切片使用值来make分配：%v\n", items1)
}

// map 字典包
func main() {
	//map的一些知识点
	fmt.Println("------- 字典初始化 --------")
	test()

	//字典初始化
	fmt.Println()
	fmt.Println("------- 字典初始化 --------")
	mapInit()

	//字典基本操作
	fmt.Println()
	fmt.Println("------- 字典基本操作 --------")
	mapBasic()

	//map类型的切片
	fmt.Println()
	fmt.Println("------- map类型的切片 --------")
	mapAdvanced()
}
