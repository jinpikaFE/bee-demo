package controllers

import (
	"bee-demo/formvalidate"
	"bee-demo/models"
	"bee-demo/utils"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	web.Controller
}

// @Title 登录
// @Description 登录
// @Param	body		body 	formvalidate.LoginParams	true	"登录参数"
// @Success 200 {array} models.LoginParams
// @Failure 500 获取数据失败
// @router / [post]
func (c *LoginController) Login() {
	var loginParams formvalidate.LoginParams
	userController := UserController{}

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

	user, err := userController.GetUserByLogin(&loginParams)
	if err != nil {
		// 处理错误
		models.RespondWithJSON(&c.Controller, "登录失败", map[string]string{"error": err.Error()}, 400, 400)
		return
	}

	token := utils.CreateToken(user)
	models.RespondWithJSON(&c.Controller, "登录成功", map[string]interface{}{"token": token, "userID": user.Id, "userName": user.UserName})
}

// @Title 注册
// @Description 注册
// @Param	body		body 	formvalidate.User	true	"注册"
// @Success 200 {array} models.User
// @Failure 500 获取数据失败
// @router /register [post]
func (c *LoginController) Register() {
	var userForm formvalidate.User

	// 使用通用解析函数处理请求体
	if err := utils.ParseRequestBody(&c.Controller, &userForm); err != nil {

		models.RespondWithJSON(&c.Controller, "", map[string]string{"error": err.Error()}, 400, 400)
		return
	}

	if err := utils.ValidParams(&userForm); err != nil {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		models.RespondWithJSON(&c.Controller, "创建失败", err.Key+err.Message, 400, 400)
		return
	}

	password, hashErr := utils.HashPassword(userForm.Password)
	if hashErr != nil {
		models.RespondWithJSON(&c.Controller, "创建失败", map[string]string{"error": hashErr.Error()}, 500, 500)
		return
	}

	userModel := models.User{
		Password: password,
		UserName: userForm.UserName,
	}

	o := orm.NewOrm()
	_, err := o.Insert(&userModel)
	if err != nil {
		models.RespondWithJSON(&c.Controller, "创建失败", map[string]string{"error": err.Error()}, 500, 500)
		return
	}
	models.RespondWithJSON(&c.Controller, "创建成功", userModel)
}
