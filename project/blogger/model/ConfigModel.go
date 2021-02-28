package model

type Config struct {
	MysqlConf MySqlConfig `ini:"mysql"`
}

type MySqlConfig struct {
	UserName string `ini:"username"`
	Passwd   string `ini:"passwd"`
	Database string `ini:"database"`
	Host     string `ini:"host"`
	Port     string `ini:"port"`
}
