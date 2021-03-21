package models

import (
	"blogger/pkg/setting"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

//modeldb句柄
var db *sqlx.DB

//init 引用依赖models包的时候默认初始化db
//db 成为models包全局的db
//这里注意这里是models的全局db，db并没有大写
func init() {
	var err error
	db, err = sqlx.Open(setting.SqlType, setting.DbDns)

	if err != nil {
		log.Println(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)

	//测试数据库是否已经联通
	err = db.Ping()
	if err != nil {
		log.Println(err)
	}
}

func CloseDb() {
	defer db.Close()
}
