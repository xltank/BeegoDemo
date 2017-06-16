package controllers

import (
    "github.com/astaxie/beego"
    "BeegoTest/models"
    "fmt"
    "strconv"
)


type ApplicationsController struct {
    beego.Controller
}

func (App *ApplicationsController) URLMapping(){
    App.Mapping("List", App.List)
}

// @router /apps [get]
func (App *ApplicationsController) List(){
    apps, err := models.GetFrist10()
    if err != nil {
        beego.Critical("Error: ", err)
    }
    App.Ctx.Output.JSON(apps, false, false)
}

// @router /apps/panic [get]
func (App *ApplicationsController) TestPanic(){

    defer func (){
        beego.Info("defer 1")
    }()

    defer func (){
        if r := recover(); r != nil {
            beego.Info("recover with:", r)
            App.Ctx.Output.JSON(r, false, false)
        }
        beego.Info("defer with recover")
    }()

    defer func (){
        beego.Info("defer 2")
    }()

    panic("Panic test")
}


// @router /apps/:id [get]
func (App *ApplicationsController) GetOne(){
    id, _ := strconv.Atoi(App.Ctx.Input.Param(":id"))
    app, err := models.GetOneById(id)
    fmt.Println(app)
    if err != nil {

    }
    App.Ctx.Output.JSON(app, false, false)
}

