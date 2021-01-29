package main

import (
	"fmt"
	"reflect"
)

//reflectType 通过内置包reflect来获取类型信息
func reflectType(a interface{}) {
	t := reflect.TypeOf(a)
	k := t.Kind()
	switch k {
	case reflect.Int16:
		fmt.Println("类型是int16")
	case reflect.Int32:
		fmt.Println("类型是int32")
	case reflect.Int64:
		fmt.Println("类型是int64")
	case reflect.Float64:
		fmt.Println("类型是Float64")
	default:
		fmt.Println("没有找到类型")
	}
}

//reflectValue 通过内置包reflect来获取值信息
func reflectValue(a interface{}) {
	t := reflect.ValueOf(a)
	fmt.Printf("获取类型：%v\n", t.Type().Name())
	k := t.Kind()
	switch k {
	case reflect.Int16:
		fmt.Printf("类型是int16,值%d\n", t.Int())
	case reflect.Int32:
		fmt.Printf("类型是int32,值%\n", t.Int())
	case reflect.Int64:
		fmt.Printf("类型是int64,值%d\n", t.Int())
	case reflect.Float64:
		fmt.Printf("类型是Float64,值%f\n", t.Float())
	default:
		fmt.Printf("没有找到类型\n")
	}
}

//reflectSetValue 通过内置包reflect来设置值信息
func reflectSetValue(a interface{}) {
	t := reflect.ValueOf(a)
	fmt.Printf("获取类型：%v\n", t.Type().Name())
	k := t.Kind()
	switch k {
	case reflect.Int16:
		fmt.Printf("类型是int16,值%d\n", t.Int())
	case reflect.Int32:
		fmt.Printf("类型是int32,值%\n", t.Int())
	case reflect.Int64:
		fmt.Printf("类型是int64,值%d\n", t.Int())
	case reflect.Float64:
		fmt.Printf("类型是Float64,值%f\n", t.Float())
	case reflect.Ptr:
		t.Elem().SetInt(6)
		fmt.Println("设置值信息为6")
	default:
		fmt.Printf("没有找到类型\n")
	}
}

func main() {
	var x int64 = 12
	reflectType(x)
	fmt.Println()
	reflectValue(x)
	fmt.Println()
	reflectSetValue(&x)

	fmt.Println(x)
}
