package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	Id       int `json:"id"`
	UserName string `orm:"size(100)" json:"userName" xml:"userName" form:"userName" valid:"Required"`
	Password string `json:"password" xml:"password" form:"password" valid:"Required"`
}

func init() {
	// 注册模型
	orm.RegisterModel(new(User))
}
