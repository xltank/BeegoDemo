package main

import (
	_ "test/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {

	//fmt.Println(beego.AppConfig.String("staging::mysqlUser"))
	//fmt.Println(beego.BConfig.WebConfig.StaticDir) // map[/static:static]

	if beego.BConfig.RunMode == "dev" {
		//beego.BConfig.WebConfig.DirectoryIndex = true
		//beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.SetLogger(logs.AdapterFile, `{"filename": "logs/test.log"}`)

	beego.Run()
}
