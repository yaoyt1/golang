package main

import (
	"logger"
)

// 测试logger 包
func main() {
	initLogger("file", "项目", "log", "debug")
	run()

	initLogger("console", "项目", "log", "debug")
	run()

	return
}

func initLogger(typeName, logPath, logName string, level string) (err error) {
	var config map[string]string = map[string]string{
		"log_level":    level,
		"log_filePath": logPath,
		"log_fileName": logName,
	}
	err = logger.InitLogger(typeName, config)
	if err != nil {
		return
	}

	return
}

func run() {
	for {
		logger.Debug("Debug用户【%s】来自中国,qweqwewq123213jk就看到积分卡放假啊都DJ发空间待付款打飞机打飞机饭卡放假得空放的假啊发", "123")
		logger.Trace("Trace用户【%s】来自中国qweqwewq123213jk就看到积分卡放假啊都DJ发空间待付款打飞机打飞机饭卡放假得空放的假啊发", "123")
		logger.Info("Info用户【%s】来自中国qweqwewq123213jk就看到积分卡放假啊都DJ发空间待付款打飞机打飞机饭卡放假得空放的假啊发", "123")
		logger.Warn("Warn用户【%s】来自中国qweqwewq123213jk就看到积分卡放假啊都DJ发空间待付款打飞机打飞机饭卡放假得空放的假啊发", "123")
		logger.Error("Error用户【%s】来自中国qweqwewq123213jk就看到积分卡放假啊都DJ发空间待付款打飞机打飞机饭卡放假得空放的假啊发", "123")
		logger.Fatal("Fatal用户【%s】来自中国qweqwewq123213jk就看到积分卡放假啊都DJ发空间待付款打飞机打飞机饭卡放假得空放的假啊发", "123")
		logger.Error("报错")

		//	time.Sleep(time.Second)
	}
}
