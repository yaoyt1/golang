package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type student struct {
		stuName string
		stuAge  int
		stuSex  string
	}

	var zhangSan = &student{
		stuName: "yyt",
		stuAge:  18,
		stuSex:  "男",
	}

	zhangShanStr, err := json.Marshal(zhangSan)
	if err != nil {
		fmt.Println("json序列化报错")
	}
	fmt.Println("json 序列化后：", zhangShanStr)

	err = json.Unmarshal([]byte(zhangShanStr), zhangSan)
	if err != nil {
		fmt.Println("json序列化报错")
	}
	fmt.Printf("json 反序列化后：%#v\n", zhangSan)

	fmt.Println(zhangSan)
}
