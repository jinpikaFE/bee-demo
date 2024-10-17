package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Test struct {
	Id   int
	Name string `orm:"size(100)" json:"name" xml:"name" form:"name" valid:"Required"`
	Age  int    `json:"age" xml:"age" form:"age" valid:"Required"`
}

func init() {
	// 注册模型
	orm.RegisterModel(new(Test))
}
