package main

import (
	"fmt"
	"reflect"
)

//struct 初始化
func structIntoTest1() {
	type student struct {
		name string
		age  int
	}

	zhangshan := new(student)
	zhangshan.name = "zhangshan"
	zhangshan.age = 18

	fmt.Printf("学生%s\t 年龄：%d岁\n", zhangshan.name, zhangshan.age)
	fmt.Println(zhangshan)

	var lisi student = student{name: "lisi", age: 18}
	fmt.Printf("学生%s\t 年龄：%d岁\n", lisi.name, lisi.age)
	fmt.Println(lisi)

	ms := &student{"Chris", 18}
	fmt.Printf("学生%s\t 年龄：%d岁\n", ms.name, ms.age)
	fmt.Println(ms)

}

// struct 注释以及注释使用
func structTag() {
	type student struct {
		name string "名称"
		age  int    "年龄"
	}

	zhangshan := student{"yyt", 18}

	ttType := reflect.TypeOf(zhangshan)
	for i := 0; i < 2; i++ {
		ixField := ttType.Field(i)
		fmt.Printf("%v\n", ixField.Tag)
	}

	fmt.Println(zhangshan)
}

//匿名结构类型
//注： 一个数据类型只能有一个匿名类型
func anonymousStruct() {
	type a struct {
		name string
		age  int
	}
	type b struct {
		name string
		age  int
		ba   a
		int
	}

	anonymousA := a{"yyt", 20}

	anonymousB := new(b)
	anonymousB.name = "yyt"
	anonymousB.age = 18
	anonymousB.ba = anonymousA
	anonymousB.int = 20

	fmt.Printf("anonymousA=%v\nanonymousB=%v\n", anonymousA, anonymousB)
}

func main() {
	structIntoTest1()
	structTag()

	fmt.Println("结构体匿名字段")
	anonymousStruct()
}
