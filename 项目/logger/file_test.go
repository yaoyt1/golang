package logger

import (
	"testing"
)

func TestFileDebug(t *testing.T) {
	l := NewFileLogger(LogLevelDebug, "/Users/yaoyt98/Library/Mobile Documents/com~apple~CloudDocs/code/git/github_golang/项目/logger", "log")
	l.Debug("用户【%s】来自中国", "123")

	l.Error("报错")
	l.SetLevel(LogLevelError)
	l.Error("报错")

	l.Close()
}

func TestConsoleDebug(t *testing.T) {
	l := NewConsoleLogger(LogLevelDebug)
	l.Debug("用户【%s】来自中国", "123")

	l.Error("报错")
	l.SetLevel(LogLevelError)
	l.Error("报错")
}

func TestFileDebugMap(t *testing.T) {
	var config map[string]string = map[string]string{
		"log_level":    "1",
		"log_filePath": "/Users/yaoyt98/Library/Mobile Documents/com~apple~CloudDocs/code/git/github_golang/项目/logger",
		"log_fileName": "log",
	}
	InitLogger("file", config)

	Debug("Debug用户【%s】来自中国", "1")
	Trace("Trace用户【%s】来自中国", "1")
	Info("Info用户【%s】来自中国", "1")
	Warn("Warn用户【%s】来自中国", "1")
	Error("Error用户【%s】来自中国", "1")
	Fatal("Fatal用户【%s】来自中国", "1")
}

func TestConsoleDebugMap(t *testing.T) {
	var config map[string]string = map[string]string{
		"log_level":    "1",
		"log_filePath": "/Users/yaoyt98/Library/Mobile Documents/com~apple~CloudDocs/code/git/github_golang/项目/logger",
		"log_fileName": "log",
	}
	InitLogger("console", config)

	Debug("用户【%s】来自中国", "123")

	Error("报错")
}
