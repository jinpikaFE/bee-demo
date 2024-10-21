package utils

import (
	"log"

	"github.com/beego/beego/v2/core/validation"
)

// ParseRequestBody 解析请求体到目标结构体，支持 JSON 和 application/x-www-form-urlencoded
func ValidParams(target interface{}) *validation.Error {
	valid := validation.Validation{}

	valid.Valid(target)
	log.Println(target)
	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)

			return err
		}
	}
	return nil
}
