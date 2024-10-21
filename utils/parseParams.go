package utils

import (
	"encoding/json"
	"errors"
	"log"
	"reflect"

	"github.com/beego/beego/v2/client/orm"
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

// UpdateModel 更新模型，根据非零字段进行更新 excludeFields 排除不想修改的字段
func UpdateModel(o orm.Ormer, id int, model interface{}, form interface{}, excludeFields ...string) error {
	vModel := reflect.ValueOf(model).Elem()
	vForm := reflect.ValueOf(form)

	// 读取当前模型信息
	if err := o.Read(model); err != nil {
		return err
	}

	// 使用反射更新非零值
	for i := 0; i < vForm.NumField(); i++ {
		field := vForm.Type().Field(i)
		value := vForm.Field(i)

		if value.IsValid() && !value.IsZero() {
			// 检查是否在排除字段中
			if !contains(excludeFields, field.Name) {
				vModel.FieldByName(field.Name).Set(value)
			}
		}
	}

	// 获取需要更新的字段
	updatedFields := []string{}
	for i := 0; i < vModel.NumField(); i++ {
		field := vModel.Type().Field(i)
		if field.Name != "Id" && !vModel.Field(i).IsZero() {
			updatedFields = append(updatedFields, field.Name)
		}
	}

	// 执行更新
	if _, err := o.Update(model, updatedFields...); err != nil {
		return err
	}

	return nil
}

// contains 检查切片中是否包含指定字段
func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
