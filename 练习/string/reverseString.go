package main

import "fmt"

// reverseStringV1 只适合英文，数字， 只占一个字节的文字
// 这里换成中文的话会乱码
func reverseStringV1() {
	var str string = "123456"
	var byteSlice []byte = []byte(str)

	for i := 0; i < len(byteSlice)/2; i++ {
		byteSlice[len(byteSlice)-i-1], byteSlice[i] = byteSlice[i], byteSlice[len(byteSlice)-i-1]
	}

	str = string(byteSlice)
	fmt.Printf("str reverse=%s\n", str)
}

// reverseStringV2 适合英文，数字， 中文
func reverseStringV2() {
	var str string = "我爱中华"
	var runeSlice []rune = []rune(str)

	for i := 0; i < len(runeSlice)/2; i++ {
		runeSlice[len(runeSlice)-i-1], runeSlice[i] = runeSlice[i], runeSlice[len(runeSlice)-i-1]
	}

	str = string(runeSlice)
	fmt.Printf("str reverse=%s\n", str)
}

// strIsHuiWei： 判断是否是回文
// 回文： 字符串倒序后还和原来的是一样的
func strIsHuiWei() {
	str := "上海自来水来自海上"

	var runeSlice []rune = []rune(str)

	for i := 0; i < len(runeSlice)/2; i++ {
		runeSlice[len(runeSlice)-i-1], runeSlice[i] = runeSlice[i], runeSlice[len(runeSlice)-i-1]
	}

	str2 := string(runeSlice)

	if str == str2 {
		fmt.Printf("str:'%s'是回文", str)
	} else {
		fmt.Printf("str:'%s'不是回文", str)
	}
}

func main() {
	reverseStringV1()
	reverseStringV2()
	strIsHuiWei()
}
