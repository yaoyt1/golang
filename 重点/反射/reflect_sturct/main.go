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

func reflectSturct(s interface{}) {
	v := reflect.ValueOf(s)
	t := v.Type()
	k := t.Kind()
	switch k {
	case reflect.Int16:
		fmt.Println("变量类型是int16")
	case reflect.Struct:
		fmt.Printf("总共有%d个字段\n", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			fmt.Printf("名称%v，类型%v，值%v.\n", t.Field(i).Name, field.Type(), field.Interface())
		}
	default:
		fmt.Println("default")
	}
}

func reflectSturctSetValue(s interface{}) {
	v := reflect.ValueOf(s)
	v.Elem().Field(0).SetString("001")
	v.Elem().FieldByName("Age").SetInt(12)
}

func main() {
	var s student
	reflectSturct(s)

	reflectSturctSetValue(&s)
	fmt.Println(s)
}
