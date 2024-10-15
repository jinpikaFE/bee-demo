package controllers

import (
	"bee-demo/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

type TestController struct {
	web.Controller
}

// @router / [get]
func (c *TestController) GetTests() {
	var tests []models.Test
	o := orm.NewOrm()
	_, err := o.QueryTable(new(models.Test)).All(&tests)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Ctx.Output.Body([]byte("Error fetching tests"))
		return
	}
	c.Data["json"] = tests
	c.ServeJSON()
}

// @Title 创建test
// @Description 创建test
// @Param	body		body 	models.Test	true		"body for test content"
// @Success 200 {int} models.User
// @Failure 403 body is empty
// @router / [post]
func (c *TestController) CreateTest() {
	var test models.Test
	if err := c.ParseForm(&test); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Invalid input"))
		return
	}
	o := orm.NewOrm()
	_, err := o.Insert(&test)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Ctx.Output.Body([]byte("Error creating test"))
		return
	}
	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = test
	c.ServeJSON()
}

// @router /:id [get]
func (c *TestController) GetTest() {
	id, err := c.GetInt(":id")
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Invalid test ID"))
		return
	}
	o := orm.NewOrm()
	test := models.Test{Id: id}
	if err := o.Read(&test); err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Ctx.Output.Body([]byte("Test not found"))
		return
	}
	c.Data["json"] = test
	c.ServeJSON()
}

// @router /:id [put]
func (c *TestController) UpdateTest() {
	id, err := c.GetInt(":id")
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Invalid test ID"))
		return
	}
	o := orm.NewOrm()
	test := models.Test{Id: id}
	if err := o.Read(&test); err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Ctx.Output.Body([]byte("Test not found"))
		return
	}
	if err := c.ParseForm(&test); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Invalid input"))
		return
	}
	if _, err := o.Update(&test); err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Ctx.Output.Body([]byte("Error updating test"))
		return
	}
	c.Data["json"] = test
	c.ServeJSON()
}

// @router /:id [delete]
func (c *TestController) DeleteTest() {
	id, err := c.GetInt(":id")
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.Body([]byte("Invalid test ID"))
		return
	}
	o := orm.NewOrm()
	test := models.Test{Id: id}
	if _, err := o.Delete(&test); err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Ctx.Output.Body([]byte("Test not found"))
		return
	}
	c.Ctx.Output.SetStatus(204) // No content
}
