package util

import (
	"os"
	"path/filepath"
	"regexp"
)

//VerifyEmailFormat 检查邮箱地址
func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func GetRootDir() (rootPath string) {
	exePath := os.Args[0]
	rootPath = filepath.Dir(exePath)
	return rootPath
}
