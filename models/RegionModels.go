package models

import (
    "github.com/astaxie/beego/orm"
)

type Province struct {
    Id   int    `orm:"auto"`
    Name string `orm:"size(100)"`
}

type City struct {
    Id         int      `orm:"auto"`
    Name       string   `orm:"size(100)"`
    ProvinceId *Province `orm:"rel(fk)"`
}

type District struct {
    Id     int    `orm:"auto"`
    Name   string `orm:"size(100)"`
    CityId *City  `orm:"rel(fk)"`
}

func init() {
    orm.RegisterModel(new(Province), new(City), new(District))
}
