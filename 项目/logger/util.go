package logger

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"time"
)

//GetLineInfo 得到行数
func getLineInfo() (fileName, funcName string, lineNo int) {
	//得到报错的行数
	pc, fileName, lineNo, ok := runtime.Caller(3)
	if ok {
		//得到方法名
		funcName = runtime.FuncForPC(pc).Name()
	}
	fileName = path.Base(fileName)
	funcName = path.Base(funcName)

	return
}

//WriteLog 写日志的函数
func writeLog(f *os.File, format string, args ...interface{}) {
	//获取时间
	now := time.Now()
	nowStr := now.Format("2006-01-02 15:04:05.999")
	//获取行数
	filenName, funcName, lineNo := getLineInfo()
	msg := fmt.Sprintf("%s\t%s\t%s【line:%d】\n%s\n\n", nowStr, filenName, funcName, lineNo, fmt.Sprintf(format, args...))

	//写入
	fmt.Fprintf(f, msg)
}
