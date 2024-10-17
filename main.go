package main

import (
	_ "bee-demo/routers"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	beego "github.com/beego/beego/v2/server/web"
)

// DBConfig 存储数据库配置
type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
	Charset  string
	Debug    bool
}

// 获取配置值的通用函数
func getConfigValue(key string) string {
	value, err := web.AppConfig.String(key)
	if err != nil {
		log.Fatalf("Failed to load %s: %v", key, err)
	}
	return value
}

// 初始化数据库配置
func initDBConfig() *DBConfig {
	// 从配置文件中读取数据库配置
	return &DBConfig{
		User:     getConfigValue("db_user"),
		Password: getConfigValue("db_password"),
		Host:     getConfigValue("db_host"),
		Port:     getConfigValue("db_port"),
		Name:     getConfigValue("db_name"),
		Charset:  getConfigValue("db_charset"),
		Debug:    getConfigBoolValue("db_debug"),
	}
}

// 获取布尔值配置的通用函数
func getConfigBoolValue(key string) bool {
	value, err := web.AppConfig.Bool(key)
	if err != nil {
		log.Printf("Failed to load %s, using default (false): %v", key, err)
		return false
	}
	return value
}

func init() {
	// logs.SetLogger("console")
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"./log/test.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
	logs.EnableFuncCallDepth(true)

	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 初始化数据库配置
	config := initDBConfig()

	// 构造数据库连接字符串
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		config.User, config.Password, config.Host, config.Port, config.Name, config.Charset)

	logs.Info("this %s cat is %v years old", "yellow", 3)

	// 注册数据库
	orm.RegisterDataBase("default", "mysql", connStr)

	// 自动建表
	orm.RunSyncdb("default", false, true)

	// 是否启用 ORM 调试模式
	if config.Debug {
		orm.Debug = true
	}
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
