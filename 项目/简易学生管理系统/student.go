package main

type student struct {
	name  string  "学生名称"
	age   int     "学生年龄"
	class string  "学生班级"
	rank  float32 "学生排名"
}

//student.go 的构造函数
func newStudent(stuName string, stuAge int, stuClass string, stuRank float32) *student {
	return &student{
		name:  stuName,
		age:   stuAge,
		class: stuClass,
		rank:  stuRank,
	}
}
