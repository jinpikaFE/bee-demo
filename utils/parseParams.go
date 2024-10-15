package utils

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/beego/beego/v2/server/web"
)

// ParseRequestBody 解析请求体到目标结构体，支持 JSON 和 application/x-www-form-urlencoded
func ParseRequestBody(c *web.Controller, target interface{}) error {
	contentType := c.Ctx.Input.Header("Content-Type")
	log.Println(c.Ctx.Request.Form)

	if contentType == "application/json" {
		// 处理 JSON 格式
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, target); err != nil {
			return err
		}
	} else if contentType == "application/x-www-form-urlencoded" {
		// 处理 application/x-www-form-urlencoded 格式

		if err := c.ParseForm(target); err != nil {
			return err
		}
	} else {
		return errors.New("Unsupported Content-Type")
	}

	return nil
}
