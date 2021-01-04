package main

import (
	"fmt"
	"os"
)

//显示菜单
func showMenu() {
	fmt.Println("1. 增加学生信息")
	fmt.Println("2. 删除学生信息")
	fmt.Println("3. 显示所有学生信息")
	fmt.Println("4. 退出系统")
}

//接收参数
func receiveProcessParamenter() {
	var para int
	fmt.Println("\n请选择选项？")
	fmt.Scanf("%d", &para)

	switch para {
	case 1:
		addStudent()
	case 2:
		deletStudent()
	case 3:
		showStudent()
	case 4:
		os.Exit(0)
	}
}

func main() {
	for {
		//显示菜单
		showMenu()

		//接收参数并处理
		receiveProcessParamenter()
	}
}
