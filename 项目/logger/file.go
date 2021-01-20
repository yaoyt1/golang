package logger

import (
	"fmt"
	"os"
)

//FileLogger 文件日志
type FileLogger struct {
	level     int
	filePath  string
	fileName  string
	file      *os.File
	errorFile *os.File
}

//NewFileLogger FileLogger 结构体构造函数
func NewFileLogger(level int, filePath, fileName string) LoggerInterface {
	logger := &FileLogger{
		level:    level,
		filePath: filePath,
		fileName: fileName,
	}

	logger.fileIinit()
	return logger
}

//init 初始化打印日志需要的文件
func (f *FileLogger) fileIinit() {
	filePath := fmt.Sprintf("%s/%s.log", f.filePath, f.fileName)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)

	if err != nil {
		panic(fmt.Sprintf("打开文件【%s】失败,错误如下：%v", filePath, err))
	}
	f.file = file

	filePath = fmt.Sprintf("%s/%s.err.log", f.filePath, f.fileName)
	file, err = os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)

	if err != nil {
		panic(fmt.Sprintf("打开文件【%s】失败,错误如下：%v", filePath, err))
	}
	f.errorFile = file
}

//SetLevel 设置日志等级
func (f *FileLogger) SetLevel(levl int) {
	f.level = levl
}

//Debug 调试日志
func (f *FileLogger) Debug(format string, args ...interface{}) {
	if f.level < LogLevelDebug || f.level > LogLevelFatal {
		return
	}

	writeLog(f.file, format, args...)
}

//Trace 跟踪日志
func (f *FileLogger) Trace(format string, args ...interface{}) {
	if f.level < LogLevelTrace || f.level > LogLevelFatal {
		return
	}

	writeLog(f.file, format, args...)
}

//Info 访问日志
func (f *FileLogger) Info(format string, args ...interface{}) {
	if f.level < LogLevelInfo || f.level > LogLevelFatal {
		return
	}

	writeLog(f.file, format, args...)
}

//Warn 警告日志
func (f *FileLogger) Warn(format string, args ...interface{}) {
	if f.level < LogLevelWarn || f.level > LogLevelFatal {
		return
	}

	writeLog(f.file, format, args...)
}

//Error 错误日志
func (f *FileLogger) Error(format string, args ...interface{}) {
	if f.level < LogLevelError || f.level > LogLevelFatal {
		msg := fmt.Sprintf("当前日志级别为【%s】不能写入错误日志\n%s", getLoggerLevelText(f.level), format)
		writeLog(f.file, msg, args...)
		return
	}

	writeLog(f.errorFile, format, args...)
}

//Fatal 严重错误日志
func (f *FileLogger) Fatal(format string, args ...interface{}) {
	if f.level != LogLevelFatal {
		msg := fmt.Sprintf("当前日志级别为【%s】不能写入错误日志\n%s", getLoggerLevelText(f.level), format)
		writeLog(f.file, msg, args...)
		return
	}

	writeLog(f.errorFile, format, args...)
}

//Close 关闭文件
func (f *FileLogger) Close() {
	f.file.Close()
	f.errorFile.Close()
}
