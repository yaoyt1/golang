package logger

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"time"
)

type loggerData struct {
	timeStr  string
	message  string
	fileName string
	funcName string
	levelStr string
	lineNo   int
	isError  bool
}

//GetLineInfo 得到行数
func getLineInfo() (fileName, funcName string, lineNo int) {
	//得到报错的行数
	pc, fileName, lineNo, ok := runtime.Caller(4)
	if ok {
		//得到方法名
		funcName = runtime.FuncForPC(pc).Name()
	}
	fileName = path.Base(fileName)
	funcName = path.Base(funcName)

	return
}

//getLogData 获取写日志需要的数据
func getLogData(level int, format string, args ...interface{}) *loggerData {
	//获取时间
	now := time.Now()
	nowStr := now.Format("2006-01-02 15:04:05.999")
	//获取行数
	fileName, funcName, lineNo := getLineInfo()

	logData := &loggerData{
		message:  fmt.Sprintf(format, args...),
		timeStr:  nowStr,
		fileName: fileName,
		funcName: funcName,
		lineNo:   lineNo,
		levelStr: getLoggerLevelText(level),
		isError:  false,
	}

	if level >= LogLevelError && level <= LogLevelFatal {
		logData.isError = true
	}

	return logData
}

//consoleWriteLog 写日志的函数
func consoleWriteLog(f *os.File, level int, format string, args ...interface{}) {
	logData := getLogData(level, format, args...)
	fmt.Fprintf(f, "【%s】%s\t%s\t%s【line:%d】\n%s\n\n", logData.levelStr, logData.timeStr, logData.fileName, logData.funcName, logData.lineNo, logData.message)
}
