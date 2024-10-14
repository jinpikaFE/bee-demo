package controllers

import (
	"beego-demo/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type RegionController struct {
    beego.Controller
}

// 获取所有省份
func (r *RegionController) GetProvinces() {
    o := orm.NewOrm()
    var provinces []models.Province
    _, err := o.QueryTable("province").All(&provinces)
    if err != nil {
        r.Data["json"] = map[string]interface{}{"code": 500, "message": "获取省份失败"}
    } else {
        r.Data["json"] = map[string]interface{}{"code": 200, "data": provinces}
    }
    r.ServeJSON()
}

// 根据省份ID获取城市
func (r *RegionController) GetCities() {
    provinceId, _ := r.GetInt("province_id")
    o := orm.NewOrm()
    var cities []models.City
    _, err := o.QueryTable("city").Filter("province_id", provinceId).All(&cities)
    if err != nil {
        r.Data["json"] = map[string]interface{}{"code": 500, "message": "获取城市失败"}
    } else {
        r.Data["json"] = map[string]interface{}{"code": 200, "data": cities}
    }
    r.ServeJSON()
}

// 根据城市ID获取区县
func (r *RegionController) GetDistricts() {
    cityId, _ := r.GetInt("city_id")
    o := orm.NewOrm()
    var districts []models.District
    _, err := o.QueryTable("district").Filter("city_id", cityId).All(&districts)
    if err != nil {
        r.Data["json"] = map[string]interface{}{"code": 500, "message": "获取区县失败"}
    } else {
        r.Data["json"] = map[string]interface{}{"code": 200, "data": districts}
    }
    r.ServeJSON()
}
