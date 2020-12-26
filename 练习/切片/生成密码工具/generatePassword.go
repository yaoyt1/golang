package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

// 定义全局变量  用来存入终端参数
var (
	flagLength int
	flagType   string
)

// 定义常量
const (
	constNums  string = "1234567890"
	constChars string = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
	constSpac  string = "`~!@#$%^&*()_+|][}{;:<>?,./"
)

// 终端接收输入参数
func terminalPara() {
	flag.IntVar(&flagLength, "l", 6, "密码长度")
	flag.StringVar(&flagType, "t", "num",
		`-t 制定密码生成的字符集，
	num: 只使用数字【0-9】
	char: 只使用英文【a-zA-Z】
	min: 使用数字和字母
	advance: 使用数字，字母以及特殊字符`)
	flag.Parse()
}

// 生成密码
func generatePwd() (pwd string) {
	var pwds []byte = make([]byte, flagLength, flagLength)
	var sourceStr string

	switch flagType {
	case "num":
		sourceStr = constNums
	case "char":
		sourceStr = constChars
	case "mix":
		sourceStr = fmt.Sprintf("%s%s", constNums, constChars)
	case "advance":
		sourceStr = fmt.Sprintf("%s%s%s", constNums, constChars, constSpac)
	default:
		sourceStr = constNums
	}

	fmt.Println("原始数据类型：", sourceStr)
	rand.Seed(time.Now().Unix())
	for i := 0; i < flagLength; i++ {
		index := rand.Intn(len(sourceStr))
		pwds[i] = sourceStr[index]
	}
	pwd = string(pwds)
	return
}

func main() {
	terminalPara()

	fmt.Printf("密码类型：%s,长度：%d\n", flagType, flagLength)
	pwd := generatePwd()
	fmt.Println("密码： ", pwd)
}
