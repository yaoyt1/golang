package logger

//LoggerInterface 定义日志接口
type LoggerInterface interface {
	SetLevel(levl int)
	LoggerInit()
	Debug(format string, args ...interface{})
	Trace(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	Close()
}
