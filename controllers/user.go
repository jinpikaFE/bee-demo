package controllers

import (
	"bee-demo/models"
	"bee-demo/utils"
	"errors"
	"log"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/server/web"
)

type UserController struct {
	web.Controller
}

// @Title 获取所有用户数据
// @Description 获取所有用户数据
// @Success 200 {array} models.User
// @Failure 500 获取数据失败
// @router / [get]
func (c *UserController) GetUsers() {
	var users []models.User
	o := orm.NewOrm()
	_, err := o.QueryTable(new(models.User)).All(&users)
	if err != nil {
		models.RespondWithJSON(&c.Controller, "查询失败", map[string]string{"error": err.Error()}, 500, 500)
		return
	}
	models.RespondWithJSON(&c.Controller, "查询成功", users)
}

// @Title 获取所有用户数据
// @Description 获取所有用户数据
// @Param page query int false "页码，默认 1"
// @Param pageSize query int false "每页数量，默认 10"
// @Success 200 {array} models.User
// @Failure 500 获取数据失败
// @router /page [get]
func (c *UserController) GetUsersPage() {
	var users []models.User
	o := orm.NewOrm()

	// 获取查询参数
	page, _ := c.GetInt("page", 1)          // 默认页码为 1
	pageSize, _ := c.GetInt("pageSize", 10) // 默认每页数量为 10

	// 计算偏移量
	offset := (page - 1) * pageSize

	querySeter := o.QueryTable(new(models.User))

	// 获取可选的查询条件
	name := c.GetString("name")
	// status, _ := c.GetInt("status", -1) // 使用 -1 作为默认值，表示不筛选状态
	age, ageErr := c.GetInt("age") // 使用 -1 作为默认值，表示不筛选状态

	valid := validation.Validation{}
	valid.Required(name, "name").Message("必须的")
	// valid.MaxSize(age, 15, "ageMax")
	valid.Range(age, 0, 110, "age")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			models.RespondWithJSON(&c.Controller, "查询失败", err.Key+err.Message, 400, 400)
			return
		}
	}

	// 动态添加条件查询
	if name != "" {
		querySeter = querySeter.Filter("name__icontains", name) // 模糊查询
	}

	if ageErr == nil {
		querySeter = querySeter.Filter("age", age) // 精确匹配状态
	}

	// 查询数据，使用 Limit 和 Offset 实现分页
	_, err := querySeter.Limit(pageSize, offset).All(&users)
	if err != nil {
		models.RespondWithJSON(&c.Controller, "查询失败", map[string]string{"error": err.Error()}, 500, 500)
		return
	}

	// 查询总记录数
	var totals int64
	totals, err = querySeter.Count()
	if err != nil {
		models.RespondWithJSON(&c.Controller, "查询失败", map[string]string{"error": err.Error()}, 500, 500)
		return
	}

	// 构造返回数据
	response := map[string]interface{}{
		"records":  users,
		"page":     page,
		"pageSize": pageSize,
		"totals":   totals,
	}

	models.RespondWithJSON(&c.Controller, "查询成功", response)
}

// @Title 创建用户数据
// @Description 创建一条用户数据
// @Param	body		body 	models.User	true	"传入的用户数据"
// @Success 201 {object} models.User
// @Failure 400 请求体格式错误
// @Failure 500 创建数据失败
// @router / [post]
func (c *UserController) CreateUser() {
	var user models.User

	// 使用通用解析函数处理请求体
	if err := utils.ParseRequestBody(&c.Controller, &user); err != nil {

		models.RespondWithJSON(&c.Controller, "", map[string]string{"error": err.Error()}, 400, 400)
		return
	}

	if err := utils.ValidParams(&user); err != nil {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		models.RespondWithJSON(&c.Controller, "创建失败", err.Key+err.Message, 400, 400)
		return
	}

	log.Println(&user)
	password, hashErr := utils.HashPassword(user.Password)
	if hashErr != nil {
		models.RespondWithJSON(&c.Controller, "创建失败", map[string]string{"error": hashErr.Error()}, 500, 500)
		return
	}
	user.Password = password

	o := orm.NewOrm()
	_, err := o.Insert(&user)
	if err != nil {
		models.RespondWithJSON(&c.Controller, "创建失败", map[string]string{"error": err.Error()}, 500, 500)
		return
	}
	models.RespondWithJSON(&c.Controller, "创建成功", user)
}

// @Title 获取指定用户数据
// @Description 根据ID获取单条用户数据
// @Param	id		path 	int	true	"用户数据ID"
// @Success 200 {object} models.User
// @Failure 400 无效的ID
// @Failure 404 数据不存在
// @router /:id [get]
func (c *UserController) GetUser() {
	id, err := c.GetInt(":id")
	if err != nil {
		models.RespondWithJSON(&c.Controller, "查询失败", map[string]string{"error": err.Error()}, 400, 400)
		return
	}
	o := orm.NewOrm()
	user := models.User{Id: id}
	if err := o.Read(&user); err != nil {
		models.RespondWithJSON(&c.Controller, "查询失败", map[string]string{"error": err.Error()}, 404, 500)
		return
	}
	models.RespondWithJSON(&c.Controller, "查询成功", user)
}

// @Title 更新用户数据
// @Description 更新指定ID的用户数据
// @Param	id		path 	int	true	"用户数据ID"
// @Param	body	body 	models.User	true	"更新后的用户数据"
// @Success 200 {object} models.User
// @Failure 400 请求体格式错误
// @Failure 404 数据不存在
// @Failure 500 更新失败
// @router /:id [put]
func (c *UserController) UpdateUser() {
	id, err := c.GetInt(":id")
	if err != nil {
		models.RespondWithJSON(&c.Controller, "更新失败", map[string]string{"error": err.Error()}, 404, 400)
		return
	}
	o := orm.NewOrm()
	user := models.User{Id: id}
	if err := o.Read(&user); err != nil {
		models.RespondWithJSON(&c.Controller, "更新失败", map[string]string{"error": err.Error()}, 404, 500)
		return
	}
	if err := utils.ParseRequestBody(&c.Controller, &user); err != nil {
		models.RespondWithJSON(&c.Controller, "更新失败", map[string]string{"error": err.Error()}, 400, 400)
		return
	}
	if _, err := o.Update(&user); err != nil {
		models.RespondWithJSON(&c.Controller, "更新失败", map[string]string{"error": err.Error()}, 500, 500)
		return
	}
	models.RespondWithJSON(&c.Controller, "更新成功", user)
}

// @Title 删除用户数据
// @Description 根据ID删除用户数据
// @Param	id		path 	int	true	"用户数据ID"
// @Success 204 {string} 空
// @Failure 400 无效的ID
// @Failure 404 数据不存在
// @router /:id [delete]
func (c *UserController) DeleteUser() {
	id, err := c.GetInt(":id")
	if err != nil {
		models.RespondWithJSON(&c.Controller, "删除失败", map[string]string{"error": err.Error()}, 404, 400)
		return
	}
	o := orm.NewOrm()
	user := models.User{Id: id}
	if err := o.Read(&user); err != nil {
		models.RespondWithJSON(&c.Controller, "删除失败", map[string]string{"error": err.Error()}, 404, 500)
		return
	}
	if _, err := o.Delete(&user); err != nil {
		models.RespondWithJSON(&c.Controller, "删除失败", map[string]string{"error": err.Error()}, 404, 500)
		return
	}
	models.RespondWithJSON(&c.Controller, "删除成功", nil)
}

func (c *UserController) GetUserByLogin(loginParams *models.LoginParams) (*models.User, error) {
	o := orm.NewOrm()
	user := models.User{UserName: loginParams.UserName}

	// 读取用户数据
	if err := o.Read(&user, "UserName"); err != nil {
		return nil, err
	}

	// 检查密码
	if !utils.CheckPasswordHash(loginParams.Password, user.Password) {
		return nil, errors.New("密码错误")
	}

	return &user, nil
}
