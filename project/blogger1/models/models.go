package models

import (
	"blogger/pkg/setting"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

//model全局的db句柄
var db *gorm.DB

//init 引用依赖models包的时候默认初始化db
//db 成为models包全局的db
//这里注意这里是models的全局db，db并没有大写
func init() {
	var err error
	db, err = gorm.Open(setting.SqlType, setting.DbDns)

	if err != nil {
		log.Println(err)
	}

	//设置全局表名禁用复数
	db.SingularTable(true)
	//添加表前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.TablePrefix + defaultTableName
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDb() {
	defer db.Close()
}
