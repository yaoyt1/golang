package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

type dept struct {
	DEPTNO string
	DNAME  string
	LOC    string
}

var engine *xorm.Engine

func main() {
	sqlDataSourceOne()
	//xormExec()
	//xormInsert()
	xormSelect()
}

//sqlDataSourceOne 一般情况下如果只操作一个数据库，只需要创建一个 engine 即可。engine 是 GoRoutine 安全的
//另外对于某些数据库有连接超时设置的，可以通过起一个定期Ping的Go程来保持连接鲜活。
//对于有大量数据并且需要分区的应用，也可以根据规则来创建多个Engine
func sqlDataSourceOne() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:123456@/yaoyoutian?charset=utf8")
	if err != nil {
		fmt.Printf("链接数据库错误:%v", err)
		return
	}

	err = engine.Ping()
	if err != nil {
		fmt.Printf("链接数据库失败:%v", err)
		return
	}

	//设置字段映射模式
	//SnakeMapper 支持struct为驼峰式命名，表结构为下划线命名之间的转换，这个是默认的Maper；
	//SameMapper 支持结构体名称和对应的表名称以及结构体field名称与对应的表字段名称相同的命名；
	//GonicMapper 和SnakeMapper很类似，但是对于特定词支持更好，比如ID会翻译成id而不是i_d。
	engine.SetMapper(names.SameMapper{})

	//对sql连接池进行设置
	engine.SetMaxOpenConns(100)
	engine.SetMaxIdleConns(20)
	engine.SetConnMaxLifetime(time.Second * 20)

	//设置日志
	engine.ShowSQL(true) //设置控制台输出sql
	//engine.Logger().SetLevel(core.log) //在控制台打印调试及以上的信息；

	////日志写入文件
	//f, err := os.Create("sql.log")
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//engine.SetLogger(xorm)
	//
	////日志记录到syslog
	//logwriter, err := syslog.New(syslog.LOG_DEBUG, "rest-xorm-example")
	//if err != nil {
	//	log.Fatal("记录日志错误")
	//}
	//logger := xorm.NewSimpleLogger(logwriter)
	//engine.SetLogger(logger)

	fmt.Println("链接数据库成功")
}

func xormExec() {
	aff, err := engine.Exec("update dept set dname = ? where loc = ?", "12", "y")
	if err != nil {
		fmt.Println(err)
		return
	}
	row, _ := aff.RowsAffected()
	id, _ := aff.LastInsertId()
	fmt.Printf("aff,row=%v,affid=%v\n", row, id)
}

func xormInsert() {
	deptInfo := &dept{
		DEPTNO: "8",
		DNAME:  "88",
		LOC:    "888",
	}

	deptInfo1 := new(dept)
	deptInfo1.DEPTNO = "9"
	deptInfo1.DNAME = "99"
	deptInfo1.LOC = "999"

	deptInfo2 := &dept{
		DEPTNO: "10",
		DNAME:  "88",
		LOC:    "888",
	}

	deptInfo3 := &dept{
		DEPTNO: "11",
		DNAME:  "88",
		LOC:    "888",
	}

	deptItems := make([]*dept, 0)
	deptItems = append(deptItems, deptInfo2)
	deptItems = append(deptItems, deptInfo3)

	// [SQL] INSERT INTO `dept` (`DEPTNO`,`DNAME`,`LOC`) VALUES (?,?,?) [8 88 888]
	// [SQL] INSERT INTO `dept` (`DEPTNO`,`DNAME`,`LOC`) VALUES (?,?,?) [9 99 999] -
	aff, err := engine.Insert(deptInfo, deptInfo1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("row=%v\n", aff)

	// [SQL] INSERT INTO `dept` (`DEPTNO`,`DNAME`,`LOC`) VALUES (?, ?, ?),(?, ?, ?) [10 88 888 11 88 888]
	aff, err = engine.Insert(&deptItems)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("row=%v\n", aff)

}

func xormSelect() {
	var d dept
	has, _ := engine.Get(&d)
	fmt.Printf("%v\n%v", has, d)
}
