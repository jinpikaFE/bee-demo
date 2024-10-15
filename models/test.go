package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Test struct {
	Id   int
	Name string `orm:"size(100)" json:"name" xml:"name" form:"name"`
	Age  int    `json:"age" xml:"age" form:"age"`
}

func init() {
	// 注册模型
	orm.RegisterModel(new(Test))
}
