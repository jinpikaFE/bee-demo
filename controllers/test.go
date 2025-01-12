package controllers

import (
	"bee-demo/formvalidate"
	"bee-demo/models"
	"bee-demo/utils"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
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

	querySeter := o.QueryTable(new(models.Test))

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
	_, err := querySeter.Limit(pageSize, offset).All(&tests)
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
		"records":  tests,
		"page":     page,
		"pageSize": pageSize,
		"totals":   totals,
	}

	models.RespondWithJSON(&c.Controller, "查询成功", response)
}

// @Title 创建测试数据
// @Description 创建一条测试数据
// @Param	body		body 	formvalidate.Test	true	"传入的测试数据"
// @Success 201 {object} models.Test
// @Failure 400 请求体格式错误
// @Failure 500 创建数据失败
// @router / [post]
func (c *TestController) CreateTest() {
	var testForm formvalidate.Test

	// 使用通用解析函数处理请求体
	if err := utils.ParseRequestBody(&c.Controller, &testForm); err != nil {

		models.RespondWithJSON(&c.Controller, "", map[string]string{"error": err.Error()}, 400, 400)
		return
	}

	if err := utils.ValidParams(&testForm); err != nil {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		models.RespondWithJSON(&c.Controller, "创建失败", err.Key+err.Message, 400, 400)
		return
	}

	testModel := models.Test{
		Name: testForm.Name,
		Age:  testForm.Age,
	}

	o := orm.NewOrm()
	_, err := o.Insert(&testModel)
	if err != nil {
		models.RespondWithJSON(&c.Controller, "创建失败", map[string]string{"error": err.Error()}, 500, 500)
		return
	}
	models.RespondWithJSON(&c.Controller, "创建成功", testModel)
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
// @Param	body	body 	formvalidate.Test	true	"更新后的测试数据"
// @Success 200 {object} models.Test
// @Failure 400 请求体格式错误
// @Failure 404 数据不存在
// @Failure 500 更新失败
// @router /:id [put]
func (c *TestController) UpdateTest() {
	var testForm formvalidate.Test
	id, err := c.GetInt(":id")
	if err != nil {
		models.RespondWithJSON(&c.Controller, "更新失败", map[string]string{"error": err.Error()}, 404, 400)
		return
	}

	if err := utils.ParseRequestBody(&c.Controller, &testForm); err != nil {
		models.RespondWithJSON(&c.Controller, "更新失败", map[string]string{"error": err.Error()}, 400, 400)
		return
	}

	o := orm.NewOrm()
	if err := o.Read(&models.Test{Id: id}); err != nil {
		models.RespondWithJSON(&c.Controller, "更新失败", map[string]string{"error": err.Error()}, 404, 500)
		return
	}

	test := models.Test{
		Id: id,
	}

	if err := utils.UpdateModel(o, id, &test, testForm); err != nil {
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
