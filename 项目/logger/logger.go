package logger

import "fmt"

var log LoggerInterface

/*InitLogger 第三方调用初始化函数
file: 记录日志到指定的文件
console: 把日志打印到控制台
*/
func InitLogger(typeName string, config map[string]string) (err error) {
	switch typeName {
	case "file":
		log, err = NewFileLoggerMap(config)
	case "console":
		log, err = NewConsoleLoggerMap(config)
	default:
		err = fmt.Errorf("logger日志不支持%s", typeName)
	}
	return
}

//InitFileLogger 文件日志初始化
func InitFileLogger(level int, filePath, fileName string) {
	log = &FileLogger{
		level:    level,
		filePath: filePath,
		fileName: fileName,
	}
}

//InitConsoleLogger 控制台日志初始化
func InitConsoleLogger(level int) {
	log = &consoleLogger{
		level: level,
	}
}

//Debug 调试日志
func Debug(format string, args ...interface{}) {
	log.SetLevel(LogLevelDebug)
	log.Debug(format, args...)
}

//Trace 跟踪日志
func Trace(format string, args ...interface{}) {
	log.SetLevel(LogLevelTrace)
	log.Trace(format, args...)
}

//Info 访问日志
func Info(format string, args ...interface{}) {
	log.SetLevel(LogLevelInfo)
	log.Info(format, args...)
}

//Warn 警告日志
func Warn(format string, args ...interface{}) {
	log.SetLevel(LogLevelWarn)
	log.Warn(format, args...)
}

//Error 错误日志
func Error(format string, args ...interface{}) {
	log.SetLevel(LogLevelError)
	log.Error(format, args...)
}

//Fatal 严重错误日志
func Fatal(format string, args ...interface{}) {
	log.SetLevel(LogLevelFatal)
	log.Fatal(format, args...)
}

func Close() {
	log.Close()
}
