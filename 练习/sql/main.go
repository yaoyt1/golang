package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type dept struct {
	DeptNo int    `db:"deptno"`
	Dname  string `db:"dname"`
	Loc    string `db:"loc"`
}

//mysqlInit mysql初始化
func mysqlInit() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/yaoyoutian"

	var err error
	db, err = sql.Open("mysql", dsn)
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
	//dbQueryRow()
	//dbQuery()
	//dbInsertValues()
	//dbUpdateValues(5)
	//dbDeleteValues(5)
	//dbPrepareQuery()
	dbTrian()
}

//dbQueryRow 查询单行
//go自动做了数据库连接池，所以我们不用维护db
//但是如果查询出结果集没有解析，我们需要把结果集关闭掉
func dbQueryRow() {
	var sqlStr string = "SELECT DEPTNO,DNAME,LOC FROM dept where DEPTNO=?"
	row := db.QueryRow(sqlStr, 1)

	var deptData dept
	err := row.Scan(&deptData.DeptNo, &deptData.Dname, &deptData.Loc)
	if err != nil {
		fmt.Printf("sql 查询数据解析失败：%#v", err)
		return
	}

	fmt.Printf("数据库连接成功%#v\n", deptData)
}

//dbQuery 查询多行
//多行解析的时候， 并没有关闭， 需要手动关闭
func dbQuery() {
	var sqlStr string = "SELECT DEPTNO,DNAME,LOC FROM dept where DEPTNO>0"
	row, err := db.Query(sqlStr)

	if err != nil {
		fmt.Printf("sql 查询数据解析失败：%#v", err)
		return
	}

	defer func() {
		if row != nil {
			row.Close()
		}
	}()

	var deptData dept
	for row.Next() {
		err = row.Scan(&deptData.DeptNo, &deptData.Dname, &deptData.Loc)
		if err != nil {
			fmt.Printf("sql 查询数据解析失败：%#v", err)
			return
		}

		fmt.Printf("数据库连接成功deptno=%d,dname=%s,loc=%s\n", deptData.DeptNo, deptData.Dname, deptData.Loc)
	}
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

//dbUpdateValues 修改
func dbUpdateValues(id int) {
	var sqlStr string = "update dept set  DNAME='y1' where DEPTNO=?"
	result, err := db.Exec(sqlStr, id)

	if err != nil {
		fmt.Printf("sql 执行sql语句失败：%#v", err)
		return
	}
	affected, err := result.RowsAffected()

	fmt.Printf("sql修改成功，影响行数:%d", affected)
}

//dbUpdateValues 修改
func dbDeleteValues(id int) {
	var sqlStr string = "delete from dept where DEPTNO=?"
	result, err := db.Exec(sqlStr, id)

	if err != nil {
		fmt.Printf("sql 执行sql语句失败：%#v", err)
		return
	}
	affected, err := result.RowsAffected()

	fmt.Printf("sql删除成功，影响行数:%d", affected)
}

//dbPrepareQuery 预处理查询
//先把sql 命令传到sql服务器上
//然后在把数据传过去
//这样当有大量重复sql ，sql服务可以使用sql服务器的缓存
func dbPrepareQuery() {
	var sqlStr string = "SELECT DEPTNO,DNAME,LOC FROM dept where DEPTNO>?"
	stmt, err := db.Prepare(sqlStr) //先把sql 命令传到sql服务器上

	if err != nil {
		fmt.Printf("sql 查询数据解析失败：%#v", err)
		return
	}

	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()

	row, err := stmt.Query(1) //然后在把数据传过去
	if err != nil {
		fmt.Printf("sql 查询数据解析失败：%#v", err)
		return
	}

	defer func() {
		if row != nil {
			row.Close()
		}
	}()

	var deptData dept
	for row.Next() {
		err = row.Scan(&deptData.DeptNo, &deptData.Dname, &deptData.Loc)
		if err != nil {
			fmt.Printf("sql 预处理查询数据解析失败：%#v", err)
			return
		}

		fmt.Printf("数据库查询成功deptno=%d,dname=%s,loc=%s\n", deptData.DeptNo, deptData.Dname, deptData.Loc)
	}
}

//dbTrian go事务
func dbTrian() {
	conn, err := db.Begin()
	if err != nil {
		if conn != nil {
			conn.Rollback()
		}
		fmt.Printf("开启事务失败：%v", err)
		return
	}

	var sqlStr string = "update dept set  DNAME='y12' where DEPTNO=?"
	_, err = conn.Exec(sqlStr, 6)

	if err != nil {
		conn.Rollback()
		fmt.Printf("sql 执行sql语句失败：%#v", err)
		return
	}

	sqlStr = "update dept set  loc='yloc1' where DEPTNO=?"
	_, err = conn.Exec(sqlStr, 6)

	if err != nil {
		conn.Rollback()
		fmt.Printf("sql 执行sql语句失败：%#v", err)
		return
	}
	conn.Commit()

	fmt.Println("事务执行完成")
}
