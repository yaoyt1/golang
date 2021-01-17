package main

import (
	"flag"
	"fmt"
)

func main() {
	var b bool
	var s string
	var i int

	flag.BoolVar(&b, "sex", false, "是否是男生")
	flag.StringVar(&s, "n", "yyt", "您的姓名")
	flag.IntVar(&i, "a", 20, "年龄多少")
	flag.Parse()

	fmt.Printf("b=%v s=%s i=%d\n", b, s, i)
}
