package models

import (
	"time"

	"github.com/beego/beego/v2/server/web"
)

type ApiResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

type CommonModel struct {
	Id        int       `orm:"auto" json:"id"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"createdAt"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updatedAt"`
}

// 返回统一格式的响应
func RespondWithJSON(c *web.Controller, msg string, data interface{}, code ...int,) {
	// 设置默认的 code 和 statusCode 为 200
	responseCode := 200
	statusCode := 200

	// 如果传递了 code，则使用传入的值
	if len(code) > 0 {
		responseCode = code[0]
		statusCode = code[1]
	}

	// 设置 HTTP 状态码
	c.Ctx.Output.SetStatus(statusCode)

	// 构建返回的 JSON 响应
	response := ApiResponse{
		Code: responseCode,
		Msg:  msg,
		Data: data,
	}
	c.Data["json"] = response
	c.ServeJSON()
}
