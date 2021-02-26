//学生操作类
package main

import "fmt"

//所有的学生信息
var allStudents []*student

//学生信息输入
func inputStudent() *student {
	var stuName string
	var stuAge int
	var stuClass string
	var stuRank float32

	fmt.Println("请输入学生名字:")
	fmt.Scanf("%s", &stuName)
	fmt.Println("请输入学生年龄:")
	fmt.Scanf("%d", &stuAge)
	fmt.Println("请输入学生所在班级:")
	fmt.Scanf("%s", &stuClass)
	fmt.Println("请输入学生班级排名:")
	fmt.Scanf("%f", &stuRank)

	stu := newStudent(stuName, stuAge, stuClass, stuRank)
	return stu
}

//查找学生是否存在
func findStudent(stuName string) (booHaveStudent bool, stuIndex int) {
	booHaveStudent = false
	if len(allStudents) == 0 {
		return
	}

	for index, item := range allStudents {
		stuIndex = -1
		if item.name == stuName {
			stuIndex = index
			booHaveStudent = true
			return
		}
	}
	return
}

//增加学生信息
func addStudent() {
	fmt.Println("\n增加学生信息")
	//获取学生信息
	stu := inputStudent()

	//查询学生信息
	booHaveStudent, stuIndex := findStudent(stu.name)

	var strUpdate string

	if booHaveStudent {
		fmt.Printf("\n提示： 学生%s已录入系统\n学生信息%#v\n是否需要更新其信息(是\\否)：\n", allStudents[stuIndex].name, allStudents[stuIndex])
		fmt.Scanf("%s", &strUpdate)
	} else {
		allStudents = append(allStudents, stu)
		fmt.Printf("结果： 学生%s信息新增成功: %#v\n", stu.name, stu)
	}

	if strUpdate == "是" {
		allStudents[stuIndex] = stu
		fmt.Printf("结果： 学生%s信息更新成功: %#v\n", stu.name, stu)
	}

	fmt.Println()
	return
}

//删除学生信息
func deletStudent() {
	var stuName string
	fmt.Println("\n删除学生信息")
	fmt.Println("请输入要删除的学生名字")
	fmt.Scanf("%s", &stuName)

	//查询学生信息
	booHaveStudent, stuIndex := findStudent(stuName)

	if !booHaveStudent || len(allStudents) == 0 {
		fmt.Printf("结果： 学生：%s 不存在\n\n", stuName)
	}

	//注意这里的三个逗号表示展开切片中的元素
	allStudents = append(allStudents[:stuIndex], allStudents[stuIndex+1:]...)
	fmt.Printf("结果： 学生%s信息删除成功\n\n", stuName)
}

//显示学生信息
func showStudent() {
	fmt.Println("\n显示学生信息")
	if len(allStudents) == 0 {
		fmt.Printf("结果： 暂无任何学生\n\n")
		return
	}

	for i, v := range allStudents {
		fmt.Printf("%d 学生：%s\t 学生信息：%#v\n", i+1, v.name, v)
	}
	fmt.Println()
}
