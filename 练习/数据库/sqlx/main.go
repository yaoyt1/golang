package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type dept struct {
	DeptNo int    `db:"DEPTNO"`
	Dname  string `db:"DNAME"`
	Loc    string `db:"LOC"`
}

//mysqlInit mysql初始化
func mysqlInit() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/yaoyoutian"

	var err error
	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("连接MySQL数据库时候报错：%#v", err)
		return
	}
	//设置程序的最大连接,休闲时候的连接
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
}

func main() {
	mysqlInit()
	dbQueryRow()
	dbQuery()
	//dbInsertValues()
}

//dbQueryRow 查询单行
func dbQueryRow() {
	var deptData dept
	var sqlStr string = "SELECT DEPTNO,DNAME,LOC FROM dept where DEPTNO=?"
	err := db.Get(&deptData, sqlStr, 1)

	if err != nil {
		fmt.Printf("sql 查询数据失败：%#v\n", err)
		return
	}

	fmt.Printf("数据库查询成功%#v\n", deptData)
}

//dbQuery 查询多行
func dbQuery() {
	var deptData []dept
	var sqlStr string = "SELECT DEPTNO,DNAME,LOC FROM dept where DEPTNO>0"
	err := db.Select(&deptData, sqlStr)
	if err != nil {
		fmt.Printf("sql 查询数据失败：%#v\n", err)
		return
	}

	fmt.Printf("数据库查询成功%v\n", deptData)
}

//dbInsertValues 新增数据
func dbInsertValues() {
	var sqlStr string = "insert dept (DNAME,LOC) values(?,?)"
	result, err := db.Exec(sqlStr, "yyt", "y")

	if err != nil {
		fmt.Printf("sql 执行sql语句失败：%#v", err)
		return
	}
	affected, err := result.RowsAffected()
	lastId, err := result.LastInsertId()

	fmt.Printf("sql插入成功，影响行数:%d,插入id:%d", affected, lastId)
}
