package config

import (
	"fmt"
	"testing"
)

type Config struct {
	ServerConf ServerConfig `ini:"server"`
	MysqlConf  MysqlConfig  `ini:"mysql"`
}

type ServerConfig struct {
	Ip   string `ini:"ip"`
	Port uint   `ini:"port"`
}

type MysqlConfig struct {
	Username string  `ini:"username"`
	Passwd   string  `ini:"passwd"`
	Database string  `ini:"database"`
	Host     string  `ini:"host"`
	Port     int     `ini:"port"`
	Timeout  float32 `ini:"timeout"`
}

func TestConfig(t *testing.T) {
	var config1 Config
	err := UnConfigFileSerialization("config.ini", &config1)
	if err != nil {
		t.Errorf("反序列配置文件错误，%v", err)
	}
	fmt.Printf("打印配置文件结构体：%#v\n", config1)

	_ = ConfigFileSerialization("config1.ini", config1)
	fmt.Println("序列化数据完成")
}
