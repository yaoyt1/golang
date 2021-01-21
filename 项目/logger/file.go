package logger

import (
	"fmt"
	"os"
)

//FileLogger 文件日志
type FileLogger struct {
	level          int
	filePath       string
	fileName       string
	file           *os.File
	errorFile      *os.File
	loggerDataChan chan *loggerData
}

//NewFileLogger FileLogger 结构体构造函数
func NewFileLogger(level int, filePath, fileName string) LoggerInterface {
	log := &FileLogger{
		level:    level,
		filePath: filePath,
		fileName: fileName,
	}

	return log
}

//NewFileLoggerMap FileLogger 结构体构造函数
func NewFileLoggerMap(config map[string]string) (logger LoggerInterface, err error) {
	if len(config) == 0 {
		err = fmt.Errorf("参数为空")
		return
	}

	loglevel, ok := config["log_level"]
	if !ok {
		err = fmt.Errorf("参数config中没有找到level")
		return
	}

	filePath, ok := config["log_filePath"]
	if !ok {
		err = fmt.Errorf("参数config中没有找到filePath")
		return
	}

	fileName, ok := config["log_fileName"]
	if !ok {
		err = fmt.Errorf("参数config中没有找到fileName")
		return
	}

	level := getLoggerLevel(loglevel)
	filelogger := &FileLogger{
		level:          level,
		filePath:       filePath,
		fileName:       fileName,
		loggerDataChan: make(chan *loggerData, 50000),
	}

	filelogger.LoggerInit()
	return filelogger, err
}

//LoggerInit 初始化打印日志需要的文件
func (f *FileLogger) LoggerInit() {
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

	//这里另外开一个线程
	go f.backstageWriteLog()
}

//backstageWriteLog 后台线程写入
func (f *FileLogger) backstageWriteLog() {
	var file *os.File

	for data := range f.loggerDataChan {
		file = f.file
		if data.isError {
			file = f.errorFile
		}
		fmt.Fprintf(file, "【%s】%s\t%s\t%s【line:%d】\n%s\n\n", data.levelStr, data.timeStr, data.fileName, data.funcName, data.lineNo, data.message)
	}
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

	//获取日志信息
	logData := getLogData(f.level, format, args...)

	//写入队列 这里运用了select 判断队列是否已堵塞
	select {
	case f.loggerDataChan <- logData:
	default:
	}
}

//Trace 跟踪日志
func (f *FileLogger) Trace(format string, args ...interface{}) {
	if f.level < LogLevelTrace || f.level > LogLevelFatal {
		return
	}

	//获取日志信息
	logData := getLogData(f.level, format, args...)

	//写入队列 这里运用了select 判断队列是否已堵塞
	select {
	case f.loggerDataChan <- logData:
	default:
	}
}

//Info 访问日志
func (f *FileLogger) Info(format string, args ...interface{}) {
	if f.level < LogLevelInfo || f.level > LogLevelFatal {
		return
	}

	//获取日志信息
	logData := getLogData(f.level, format, args...)

	//写入队列 这里运用了select 判断队列是否已堵塞
	select {
	case f.loggerDataChan <- logData:
	default:
	}
}

//Warn 警告日志
func (f *FileLogger) Warn(format string, args ...interface{}) {
	if f.level < LogLevelWarn || f.level > LogLevelFatal {
		return
	}

	//获取日志信息
	logData := getLogData(f.level, format, args...)

	//写入队列 这里运用了select 判断队列是否已堵塞
	select {
	case f.loggerDataChan <- logData:
	default:
	}
}

//Error 错误日志
func (f *FileLogger) Error(format string, args ...interface{}) {
	//获取日志信息
	var logData *loggerData
	if f.level < LogLevelError || f.level > LogLevelFatal {
		msg := fmt.Sprintf("当前日志级别为【%s】不能写入错误日志\n%s", getLoggerLevelText(f.level), format)
		logData = getLogData(f.level, msg)
	} else {
		logData = getLogData(f.level, format, args...)
	}

	//写入队列 这里运用了select 判断队列是否已堵塞
	select {
	case f.loggerDataChan <- logData:
	default:
	}
}

//Fatal 严重错误日志
func (f *FileLogger) Fatal(format string, args ...interface{}) {
	//获取日志信息
	var logData *loggerData
	if f.level != LogLevelFatal {
		msg := fmt.Sprintf("当前日志级别为【%s】不能写入错误日志\n%s", getLoggerLevelText(f.level), format)
		logData = getLogData(f.level, msg)
	} else {
		logData = getLogData(f.level, format, args...)
	}

	//写入队列 这里运用了select 判断队列是否已堵塞
	select {
	case f.loggerDataChan <- logData:
	default:
	}
}

//Close 关闭文件
func (f *FileLogger) Close() {
	f.file.Close()
	f.errorFile.Close()
}
