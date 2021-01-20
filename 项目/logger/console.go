package logger

import (
	"fmt"
	"os"
)

//consoleLogger 文件日志
type consoleLogger struct {
	level int
}

//NewConsoleLogger 结构体构造函数
func NewConsoleLogger(level int) LoggerInterface {
	logger := &consoleLogger{
		level: level,
	}

	return logger
}

//NewConsoleLoggerMap 结构体构造函数
func NewConsoleLoggerMap(config map[string]string) (logger LoggerInterface, err error) {
	if len(config) == 0 {
		err = fmt.Errorf("参数为空")
		return
	}

	loglevel, ok := config["log_level"]
	if !ok {
		err = fmt.Errorf("参数config中没有找到level")
		return
	}

	level := getLoggerLevel(loglevel)
	logger = &consoleLogger{
		level: level,
	}

	return
}

//SetLevel 设置日志等级
func (f *consoleLogger) SetLevel(levl int) {
	f.level = levl
}

//Debug 调试日志
func (f *consoleLogger) Debug(format string, args ...interface{}) {
	if f.level < LogLevelDebug || f.level > LogLevelFatal {
		return
	}

	writeLog(os.Stdout, format, args...)
}

//Trace 跟踪日志
func (f *consoleLogger) Trace(format string, args ...interface{}) {
	if f.level < LogLevelTrace || f.level > LogLevelFatal {
		return
	}

	writeLog(os.Stdout, format, args...)
}

//Info 访问日志
func (f *consoleLogger) Info(format string, args ...interface{}) {
	if f.level < LogLevelInfo || f.level > LogLevelFatal {
		return
	}

	writeLog(os.Stdout, format, args...)
}

//Warn 警告日志
func (f *consoleLogger) Warn(format string, args ...interface{}) {
	if f.level < LogLevelWarn || f.level > LogLevelFatal {
		return
	}

	writeLog(os.Stdout, format, args...)
}

//Error 错误日志
func (f *consoleLogger) Error(format string, args ...interface{}) {
	if f.level < LogLevelError || f.level > LogLevelFatal {
		return
	}

	writeLog(os.Stdout, format, args...)
}

//Fatal 严重错误日志
func (f *consoleLogger) Fatal(format string, args ...interface{}) {
	if f.level != LogLevelFatal {
		return
	}

	writeLog(os.Stdout, format, args...)
}

//Close 关闭文件
func (f *consoleLogger) Close() {
}
