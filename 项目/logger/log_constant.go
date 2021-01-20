package logger

//错误日志等级变量
const (
	LogLevelDebug = iota
	LogLevelTrace
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelFatal
)

func getLoggerLevelText(level int) string {
	switch level {
	case LogLevelDebug:
		return "debug"
	case LogLevelTrace:
		return "trace"
	case LogLevelInfo:
		return "info"
	case LogLevelWarn:
		return "warn"
	case LogLevelError:
		return "error"
	case LogLevelFatal:
		return "fatal"
	default:
		return "unknow"
	}

}
