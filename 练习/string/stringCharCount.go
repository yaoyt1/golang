package main

import "fmt"

func example() {
	str := "你好啊 a,123456"
	charCount, numCount, spCount, otherCount := stringCount(str)
	fmt.Printf("charCount=%d个\n numCount=%d个\n spCount=%d个\n otherCount=%d个\n ", charCount, numCount, spCount, otherCount)
}

func stringCount(str string) (charCount, numCount, spCount, otherCount int) {
	utfChars := []rune(str)
	for i := 0; i < len(utfChars); i++ {
		if utfChars[i] >= 'a' && utfChars[i] <= 'z' || utfChars[i] >= 'A' && utfChars[i] <= 'Z' {
			charCount++
			continue
		} else if utfChars[i] >= '0' && utfChars[i] <= '9' {
			numCount++
			continue
		} else if utfChars[i] == ' ' {
			spCount++
			continue
		} else {
			otherCount++
		}
	}
	return
}

func main() {
	example()
}
