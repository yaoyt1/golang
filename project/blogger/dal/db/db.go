package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"yytGithub/project/blogger/model"
	"yytGithub/project/blogger/util"
)

var DB *sqlx.DB

func Init() error {
	var err error
	var config model.Config
	err = util.UnConfigFileSerialization("./config/config.ini", &config)
	if err != nil {
		err = fmt.Errorf("反序列配置文件错误，%v", err)
		return err
	}
	fmt.Printf("打印配置文件结构体：%#v\n", config)

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.MysqlConf.UserName,
		config.MysqlConf.Passwd, config.MysqlConf.Host, config.MysqlConf.Port, config.MysqlConf.Database)
	DB, err = sqlx.Open("mysql", dns)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(20)
	return nil
}
