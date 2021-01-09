package main

import (
	"fmt"
	"os"
)

func testFscanf() {
	var a int
	var b string
	var c float32

	fmt.Fscanf(os.Stdin, "%d%s%f", &a, &b, &c)
	fmt.Printf("a=%d;b=%s;c=%0.2f\n", a, b, c)
	/*
		fmt.Scanf()的原理还是调用fmt.Fscanf()
			   func Scanf(format string, a ...interface{}) (n int, err error) {
			   	return Fscanf(os.Stdin, format, a...)
			   }
	*/
}

func main() {
	testFscanf()
}
