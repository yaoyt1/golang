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
