package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	Id        int       `orm:"auto" json:"id"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"createdAt"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updatedAt"`
	UserName  string    `orm:"size(100);unique" json:"userName"`
	Password  string    `json:"password"`
}

func init() {
	// 注册模型
	orm.RegisterModel(new(User))
}
