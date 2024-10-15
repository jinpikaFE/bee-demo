package controllers

import (
	"bee-demo/models"
	"bee-demo/utils"
	"log"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

type TestController struct {
	web.Controller
}

// @Title 获取所有测试数据
// @Description 获取所有测试数据
// @Success 200 {array} models.Test
// @Failure 500 获取数据失败
// @router / [get]
func (c *TestController) GetTests() {
	var tests []models.Test
	o := orm.NewOrm()
	_, err := o.QueryTable(new(models.Test)).All(&tests)
	if err != nil {
		models.RespondWithJSON(&c.Controller, "查询失败", map[string]string{"error": err.Error()}, 500, 500)
		return
	}
	models.RespondWithJSON(&c.Controller, "查询成功", tests)
}

// @Title 获取所有测试数据
// @Description 获取所有测试数据
// @Param page query int false "页码，默认 1"
// @Param pageSize query int false "每页数量，默认 10"
// @Success 200 {array} models.Test
// @Failure 500 获取数据失败
// @router /page [get]
func (c *TestController) GetTestsPage() {
	var tests []models.Test
	o := orm.NewOrm()

	// 获取查询参数
	page, _ := c.GetInt("page", 1)          // 默认页码为 1
	pageSize, _ := c.GetInt("pageSize", 10) // 默认每页数量为 10

	// 计算偏移量
	offset := (page - 1) * pageSize

	// 查询数据，使用 Limit 和 Offset 实现分页
	_, err := o.QueryTable(new(models.Test)).Limit(pageSize, offset).All(&tests)
	if err != nil {
		models.RespondWithJSON(&c.Controller, "查询失败", map[string]string{"error": err.Error()}, 500, 500)
		return
	}

	// 查询总记录数
	var totals int64
	totals, err = o.QueryTable(new(models.Test)).Count()
	if err != nil {
		models.RespondWithJSON(&c.Controller, "查询失败", map[string]string{"error": err.Error()}, 500, 500)
		return
	}

	// 构造返回数据
	response := map[string]interface{}{
		"records":  tests,
		"page":     page,
		"pageSize": pageSize,
		"totals":   totals,
	}

	models.RespondWithJSON(&c.Controller, "查询成功", response)
}

// @Title 创建测试数据
// @Description 创建一条测试数据
// @Param	body		body 	models.Test	true	"传入的测试数据"
// @Success 201 {object} models.Test
// @Failure 400 请求体格式错误
// @Failure 500 创建数据失败
// @router / [post]
func (c *TestController) CreateTest() {
	var test models.Test

	// 使用通用解析函数处理请求体
	if err := utils.ParseRequestBody(&c.Controller, &test); err != nil {

		models.RespondWithJSON(&c.Controller, "", map[string]string{"error": err.Error()}, 400, 400)
		return
	}

	log.Println(&test)
	o := orm.NewOrm()
	_, err := o.Insert(&test)
	if err != nil {
		models.RespondWithJSON(&c.Controller, "创建失败", map[string]string{"error": err.Error()}, 500, 500)
		return
	}
	models.RespondWithJSON(&c.Controller, "创建成功", test)
}

// @Title 获取指定测试数据
// @Description 根据ID获取单条测试数据
// @Param	id		path 	int	true	"测试数据ID"
// @Success 200 {object} models.Test
// @Failure 400 无效的ID
// @Failure 404 数据不存在
// @router /:id [get]
func (c *TestController) GetTest() {
	id, err := c.GetInt(":id")
	if err != nil {
		models.RespondWithJSON(&c.Controller, "查询失败", map[string]string{"error": err.Error()}, 400, 400)
		return
	}
	o := orm.NewOrm()
	test := models.Test{Id: id}
	if err := o.Read(&test); err != nil {
		models.RespondWithJSON(&c.Controller, "查询失败", map[string]string{"error": err.Error()}, 404, 500)
		return
	}
	models.RespondWithJSON(&c.Controller, "查询成功", test)
}

// @Title 更新测试数据
// @Description 更新指定ID的测试数据
// @Param	id		path 	int	true	"测试数据ID"
// @Param	body	body 	models.Test	true	"更新后的测试数据"
// @Success 200 {object} models.Test
// @Failure 400 请求体格式错误
// @Failure 404 数据不存在
// @Failure 500 更新失败
// @router /:id [put]
func (c *TestController) UpdateTest() {
	id, err := c.GetInt(":id")
	if err != nil {
		models.RespondWithJSON(&c.Controller, "更新失败", map[string]string{"error": err.Error()}, 404, 400)
		return
	}
	o := orm.NewOrm()
	test := models.Test{Id: id}
	if err := o.Read(&test); err != nil {
		models.RespondWithJSON(&c.Controller, "更新失败", map[string]string{"error": err.Error()}, 404, 500)
		return
	}
	if err := utils.ParseRequestBody(&c.Controller, &test); err != nil {
		models.RespondWithJSON(&c.Controller, "更新失败", map[string]string{"error": err.Error()}, 400, 400)
		return
	}
	if _, err := o.Update(&test); err != nil {
		models.RespondWithJSON(&c.Controller, "更新失败", map[string]string{"error": err.Error()}, 500, 500)
		return
	}
	models.RespondWithJSON(&c.Controller, "更新成功", test)
}

// @Title 删除测试数据
// @Description 根据ID删除测试数据
// @Param	id		path 	int	true	"测试数据ID"
// @Success 204 {string} 空
// @Failure 400 无效的ID
// @Failure 404 数据不存在
// @router /:id [delete]
func (c *TestController) DeleteTest() {
	id, err := c.GetInt(":id")
	if err != nil {
		models.RespondWithJSON(&c.Controller, "删除失败", map[string]string{"error": err.Error()}, 404, 400)
		return
	}
	o := orm.NewOrm()
	test := models.Test{Id: id}
	if err := o.Read(&test); err != nil {
		models.RespondWithJSON(&c.Controller, "删除失败", map[string]string{"error": err.Error()}, 404, 500)
		return
	}
	if _, err := o.Delete(&test); err != nil {
		models.RespondWithJSON(&c.Controller, "删除失败", map[string]string{"error": err.Error()}, 404, 500)
		return
	}
	models.RespondWithJSON(&c.Controller, "删除成功", nil)
}
