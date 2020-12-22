// 4.1 local_scope.go
// 输出结果： GaG
package main

var a = "G"

func main() {
	n()
	m()
	n()
}

func n() {
	print(a)
}

func m() {
	a := "a"
	print(a)
}
