package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["bee-demo/controllers:LoginController"] = append(beego.GlobalControllerRouter["bee-demo/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee-demo/controllers:LoginController"] = append(beego.GlobalControllerRouter["bee-demo/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee-demo/controllers:TestController"] = append(beego.GlobalControllerRouter["bee-demo/controllers:TestController"],
        beego.ControllerComments{
            Method: "GetTests",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee-demo/controllers:TestController"] = append(beego.GlobalControllerRouter["bee-demo/controllers:TestController"],
        beego.ControllerComments{
            Method: "CreateTest",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee-demo/controllers:TestController"] = append(beego.GlobalControllerRouter["bee-demo/controllers:TestController"],
        beego.ControllerComments{
            Method: "GetTest",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee-demo/controllers:TestController"] = append(beego.GlobalControllerRouter["bee-demo/controllers:TestController"],
        beego.ControllerComments{
            Method: "UpdateTest",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee-demo/controllers:TestController"] = append(beego.GlobalControllerRouter["bee-demo/controllers:TestController"],
        beego.ControllerComments{
            Method: "DeleteTest",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee-demo/controllers:TestController"] = append(beego.GlobalControllerRouter["bee-demo/controllers:TestController"],
        beego.ControllerComments{
            Method: "GetTestsPage",
            Router: `/page`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee-demo/controllers:UserController"] = append(beego.GlobalControllerRouter["bee-demo/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetUsers",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee-demo/controllers:UserController"] = append(beego.GlobalControllerRouter["bee-demo/controllers:UserController"],
        beego.ControllerComments{
            Method: "CreateUser",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee-demo/controllers:UserController"] = append(beego.GlobalControllerRouter["bee-demo/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetUser",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee-demo/controllers:UserController"] = append(beego.GlobalControllerRouter["bee-demo/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateUser",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee-demo/controllers:UserController"] = append(beego.GlobalControllerRouter["bee-demo/controllers:UserController"],
        beego.ControllerComments{
            Method: "DeleteUser",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee-demo/controllers:UserController"] = append(beego.GlobalControllerRouter["bee-demo/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetUserCurrent",
            Router: `/current`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee-demo/controllers:UserController"] = append(beego.GlobalControllerRouter["bee-demo/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetUsersPage",
            Router: `/page`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
