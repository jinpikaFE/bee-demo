// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"bee-demo/controllers"
	"bee-demo/filters"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		// bee generate routers 命令 生成 commentsRouter 路由
		beego.NSNamespace("/test",
			// 走注解路由
			beego.NSInclude(
				&controllers.TestController{},
			),
			// // 指定方法
			// beego.NSRouter("", &controllers.TestController{}, "get:GetTests;post:CreateTest"),
		),
		beego.NSNamespace("/user",
			// 走注解路由
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/login",
			// 走注解路由
			beego.NSInclude(
				&controllers.LoginController{},
			),
		),
	)
	beego.AddNamespace(ns)

	beego.InsertFilter("/v1/user/*", beego.BeforeRouter, filters.JwtFilter)

}
