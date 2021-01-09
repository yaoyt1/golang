package main

import "fmt"

//提供字符串 空格作为分隔符
func testSscanf() {
	var a int
	var b string
	var c float32

	fmt.Sscanf("88 测试sScanf 8.8", "%d%s%f", &a, &b, &c)
	fmt.Printf("a=%d;b=%s;c=%f\n", a, b, c)
}

//提供字符串 空格和换行作为分隔符
func testSscan() {
	var a int
	var b string
	var c float32

	fmt.Sscan("8 测试Sscan 1.2", &a, &b, &c)
	fmt.Printf("a=%d;b=%s;c=%f\n", a, b, c)
}

//提供字符串 空格作为分隔符，
func testSscanln() {
	var a int
	var b string
	var c float32

	fmt.Sscanln("1 测试Sscanln 1.2", &a, &b, &c)
	fmt.Printf("a=%d;b=%s;c=%f\n", a, b, c)
}

func main() {
	testSscanf()
	testSscan()
	testSscanln()
}
