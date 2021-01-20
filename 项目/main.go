package main

import "logger"

// 调用logger 包
func main() {
	var config map[string]string = map[string]string{
		"log_level":    "1",
		"log_filePath": "项目",
		"log_fileName": "log",
	}
	logger.InitLogger("console", config)

	logger.Debug("Debug用户【%s】来自中国", "123")
	logger.Trace("Trace用户【%s】来自中国", "123")
	logger.Info("Info用户【%s】来自中国", "123")
	logger.Warn("Warn用户【%s】来自中国", "123")
	logger.Error("Error用户【%s】来自中国", "123")
	logger.Fatal("Fatal用户【%s】来自中国", "123")
	logger.Error("报错")

	logger.InitLogger("file", config)

	logger.Debug("Debug用户【%s】来自中国", "123")
	logger.Trace("Trace用户【%s】来自中国", "123")
	logger.Info("Info用户【%s】来自中国", "123")
	logger.Warn("Warn用户【%s】来自中国", "123")
	logger.Error("Error用户【%s】来自中国", "123")
	logger.Fatal("Fatal用户【%s】来自中国", "123")
	logger.Error("报错")
}
