package controllers

import (
	"bee-demo/models"
	"bee-demo/utils"

	"github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	web.Controller
}

// @Title 登录
// @Description 登录
// // @Param	body		body 	models.LoginParams	true	"登录参数"
// @Success 200 {array} models.LoginParams
// @Failure 500 获取数据失败
// @router / [post]
func (c *LoginController) Login() {
	var loginParams []models.LoginParams

	// 使用通用解析函数处理请求体
	if err := utils.ParseRequestBody(&c.Controller, &loginParams); err != nil {

		models.RespondWithJSON(&c.Controller, "登录失败", map[string]string{"error": err.Error()}, 400, 400)
		return
	}

	if err := utils.ValidParams(&loginParams); err != nil {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		models.RespondWithJSON(&c.Controller, "登录失败", err.Key+err.Message, 400, 400)
		return
	}
	// o := orm.NewOrm()
	// _, err := o.QueryTable(new(models.User)).All(&users)
	// if err != nil {
	// 	models.RespondWithJSON(&c.Controller, "查询失败", map[string]string{"error": err.Error()}, 500, 500)
	// 	return
	// }
	// models.RespondWithJSON(&c.Controller, "查询成功", users)
}
