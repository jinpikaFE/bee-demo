package main

import (
	_ "bee-demo/routers"
	"bee-demo/utils"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"
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

// 初始化数据库配置
func initDBConfig() *DBConfig {
	// 从配置文件中读取数据库配置
	return &DBConfig{
		User:     utils.GetConfigValue("db_user"),
		Password: utils.GetConfigValue("db_password"),
		Host:     utils.GetConfigValue("db_host"),
		Port:     utils.GetConfigValue("db_port"),
		Name:     utils.GetConfigValue("db_name"),
		Charset:  utils.GetConfigValue("db_charset"),
		Debug:    utils.GetConfigBoolValue("db_debug"),
	}
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

	// 注册数据库
	orm.RegisterDataBase("default", "mysql", connStr)

	// 自动建表
	orm.RunSyncdb("default", false, true)

	// 是否启用 ORM 调试模式
	if config.Debug {
		orm.Debug = true
	}
}

// 项目运行命令 bee run -gendoc=true -downdoc=true

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
