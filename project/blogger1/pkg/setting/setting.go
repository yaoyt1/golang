package setting

import (
	"fmt"
	"github.com/go-ini/ini"
	"log"
	"os"
	"time"
)

var (
	cfg *ini.File

	RunMode string

	PageSize  int
	JwtSecret string

	HttpPort     int
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration

	SqlType     string
	TablePrefix string
	DbDns       string
)

func init() {
	appConfigPath := "conf/app.ini"
	var err error
	cfg, err = ini.Load(appConfigPath)
	if err != nil {
		fmt.Printf("加载配置文件[%s]失败:%v\n", appConfigPath, err)
		os.Exit(1)
	}

	loadBase()
	loadApp()
	loadServer()
	loadDb()
}

//加载基础配置
func loadBase() {
	RunMode = cfg.Section("").Key("RunMode").String()
}

//加载App配置
func loadApp() {
	section, err := cfg.GetSection("App")
	if err != nil {
		log.Fatalf("加载节点App失败：%v\n", err)
	}

	PageSize = section.Key("PageSize").MustInt(10)
	JwtSecret = section.Key("JwtSecret").String()
}

//加载Server配置
func loadServer() {
	section, err := cfg.GetSection("Server")
	if err != nil {
		log.Fatalf("加载节点App失败：%v\n", err)
	}

	HttpPort = section.Key("HttpPort").MustInt(8080)
	ReadTimeOut = time.Duration(section.Key("ReadTimeOut").MustInt(60)) * time.Second
	WriteTimeOut = time.Duration(section.Key("WriteTimeOut").MustInt(60)) * time.Second
}

//加载loadDb配置
func loadDb() {
	section, err := cfg.GetSection("Database")
	if err != nil {
		log.Fatalf("加载节点App失败：%v\n", err)
	}

	SqlType = section.Key("Type").String()
	user := section.Key("User").String()
	pwd := section.Key("Pwd").String()
	host := section.Key("Host").String()
	port := section.Key("Port").In("3306", []string{"3306"})
	TablePrefix = section.Key("TablePrefix").String()
	dataBaseName := section.Key("DataBaseName").In("BloggerGin", []string{"yaoyoutian", "BloggerGin"})

	DbDns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pwd, host, port, dataBaseName)
}
