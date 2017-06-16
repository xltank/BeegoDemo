// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/context"
	"strings"
	"BeegoTest/controllers"
)


func auth(ctx *context.Context){
	//fmt.Println(ctx.Input.URL(), ctx.Input.Query("a"))
	url := ctx.Input.URL()
	if !strings.HasPrefix(url, "/api/public") && url != "/login.html" {
		beego.Debug("auth ...")
		// check cookie

		// to login page
		//ctx.Redirect(302, "/login.html")
	}
}

func init() {

	beego.InsertFilter("*", beego.BeforeExec, auth)

	beego.Include(&controllers.ApplicationsController{})

/*
	beego.Get("/", func(ctx *context.Context){
		ctx.Redirect(302, "/static/index.html")
	})
*/
	/*beego.Get("/:path:[^static].html", func(ctx *context.Context){
		fmt.Println(ctx.Input.Param(":splat"))
		ctx.Redirect(301, "/static/"+ctx.Input.Param(":splat"))
	})*/

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)

	/*
        beego.Get("/path*//*", func(ctx *context.Context){
            ctx.Output.Body([]byte(ctx.Input.Param(":splat"))) // a/b/c.html
        })

        beego.AutoRouter(&controllers.ItemController{})

        beego.Get("/item", func(ctx *context.Context){
            //beego.Debug("urlfor", beego.URLFor("ItemController.Get")) // urlfor /item/get
            ctx.Output.JSON(models.FindOne(), true, false)
        })*/

	// ?? does not work
	//beego.AutoRouter(&controllers.ApplicationsController{})

}
