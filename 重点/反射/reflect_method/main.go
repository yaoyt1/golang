package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Age   int
	Class string
	//xxx   int不能获取到私有变量的值
}

func (s *student) SetName(name string) {
	s.Name = name
}
func (s *student) StudentPrint() {
	fmt.Printf("反射调用： %#v\n", s)
}

func main() {
	var s student = struct {
		Name  string
		Age   int
		Class string
	}{Name: "yyt", Age: 18, Class: "高二三班"}

	v := reflect.ValueOf(&s)
	t := v.Type()
	fmt.Printf("student结构体拥有方法%d个\n", t.NumMethod())
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Printf("%d:方法名称%v,方法标签%v\n", i, t.Method(i).Name, t.Method(i).Type)
	}

	//通过reflect.Valuof来调用方法
	//调用无参方法
	meth1 := v.MethodByName("StudentPrint") //找到方法
	meth1.Call([]reflect.Value{})           //调用

	//调用有参方法
	meth2 := v.MethodByName("SetName") //找到方法
	meth2.Call([]reflect.Value{reflect.ValueOf("name")})
	meth1.Call([]reflect.Value{})
}
