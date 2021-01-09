package main

import "fmt"

//空格作为分隔符,遇到换行结束
func testScanf() {
	var a int
	var b string
	var c float32

	fmt.Scanf("%d%s%f", &a, &b, &c)
	fmt.Printf("a=%d;b=%s;c=%f\n", a, b, c)
}

//空格和换行作为分隔符
func testScan() {
	var a int
	var b string
	var c float32

	fmt.Scan(&a, &b, &c)
	fmt.Printf("a=%d;b=%s;c=%f\n", a, b, c)
}

//空格作为分隔符， 遇到换行结束
func testScanln() {
	var a int
	var b string
	var c float32

	fmt.Scanln(&a, &b, &c)
	fmt.Printf("a=%d;b=%s;c=%f\n", a, b, c)
}

func main() {
	testScanf()
	//testScan()
	//testScanln()
}
