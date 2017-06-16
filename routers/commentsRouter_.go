package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["BeegoTest/controllers:ApplicationsController"] = append(beego.GlobalControllerRouter["BeegoTest/controllers:ApplicationsController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/apps`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["BeegoTest/controllers:ApplicationsController"] = append(beego.GlobalControllerRouter["BeegoTest/controllers:ApplicationsController"],
		beego.ControllerComments{
			Method: "TestPanic",
			Router: `/apps/panic`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["BeegoTest/controllers:ApplicationsController"] = append(beego.GlobalControllerRouter["BeegoTest/controllers:ApplicationsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/apps/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["BeegoTest/controllers:ObjectController"] = append(beego.GlobalControllerRouter["BeegoTest/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["BeegoTest/controllers:ObjectController"] = append(beego.GlobalControllerRouter["BeegoTest/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["BeegoTest/controllers:ObjectController"] = append(beego.GlobalControllerRouter["BeegoTest/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["BeegoTest/controllers:ObjectController"] = append(beego.GlobalControllerRouter["BeegoTest/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["BeegoTest/controllers:ObjectController"] = append(beego.GlobalControllerRouter["BeegoTest/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["BeegoTest/controllers:UserController"] = append(beego.GlobalControllerRouter["BeegoTest/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["BeegoTest/controllers:UserController"] = append(beego.GlobalControllerRouter["BeegoTest/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["BeegoTest/controllers:UserController"] = append(beego.GlobalControllerRouter["BeegoTest/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["BeegoTest/controllers:UserController"] = append(beego.GlobalControllerRouter["BeegoTest/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["BeegoTest/controllers:UserController"] = append(beego.GlobalControllerRouter["BeegoTest/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["BeegoTest/controllers:UserController"] = append(beego.GlobalControllerRouter["BeegoTest/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["BeegoTest/controllers:UserController"] = append(beego.GlobalControllerRouter["BeegoTest/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
