package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name" db:"namedb"`
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
	var s = student{Name: "yyt", Age: 18, Class: "高二三班"}
	t := reflect.TypeOf(s)

	fied0 := t.Field(0)
	fmt.Printf("json tag:%s\n", fied0.Tag.Get("json"))
	fmt.Printf("db tag:%s\n", fied0.Tag.Get("db"))
}
