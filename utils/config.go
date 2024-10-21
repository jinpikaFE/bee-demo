package utils

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

func GetConfigValue(key string) string {
	value, err := web.AppConfig.String(key)
	if err != nil {
		logs.Error("Failed to load %s: %v", key, err)
	}
	return value
}

// 获取布尔值配置的通用函数
func GetConfigBoolValue(key string) bool {
	value, err := web.AppConfig.Bool(key)
	if err != nil {
		logs.Error("Failed to load %s, using default (false): %v", key, err)
		return false
	}
	return value
}
