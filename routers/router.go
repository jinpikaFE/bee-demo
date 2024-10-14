package routers

import (
	"beego-demo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/api/provinces", &controllers.RegionController{}, "get:GetProvinces")
    beego.Router("/api/cities", &controllers.RegionController{}, "get:GetCities")
    beego.Router("/api/districts", &controllers.RegionController{}, "get:GetDistricts")
}
