package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Test struct {
	Id        int       `orm:"auto" json:"id"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"createdAt"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updatedAt"`
	Name      string    `orm:"size(100)" json:"name"`
	Age       int       `json:"age"`
}

func init() {
	// 注册模型
	orm.RegisterModel(new(Test))
}
