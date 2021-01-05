package main

import "fmt"

type integer int

type people struct {
	name string
	age  int
}

// 这个方法前面的参数必须不是内置数据类型
func (a integer) print() {
	fmt.Println("这个是integer的方法")
}

//这里要使用指针传入， 不然方法是值传递的
func (p *people) addPeopleInfo(peoName string, peoAge int) {
	p.name = peoName
	p.age = peoAge
}

func main() {
	var i integer = 10
	i.print()

	var p *people = new(people)
	p.name = "yyt"
	p.age = 18
	fmt.Printf("使用方法前%#v\n", p)
	p.addPeopleInfo("y", 20)
	fmt.Printf("使用方法后%#v\n", p)
}
