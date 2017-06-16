package controllers

import (
    "github.com/astaxie/beego"
    "BeegoTest/models"
    "github.com/astaxie/beego/logs"
)


type ItemController struct {
    beego.Controller
}


/*func (item *ItemController) Init(ct *context.Context, childName string, app interface{}){
    //logs.GetLogger()
    logs.Debug("controller/item/init()")
}*/

func (item *ItemController) Prepare(){
    logs.Debug("controller/item/Prepare()")
    item.Ctx.Output.JSON("shortcut before business", true, false)
}

func (item *ItemController) Finish(){
    logs.Debug("controller/item/Finish()")
}

func (item *ItemController) Destructor(){
    logs.Debug("controller/item/Destructor()")
}



func (item *ItemController) Get (){
    logs.Debug("controller/item/Get()")
    item.Ctx.Output.Body([]byte("Items..."))
}

func (item *ItemController) Post(){

}


// /item/buy
func (item *ItemController) Buy (){
    item1 := models.FindOne()
    items := []models.Item{item1}
    item.Data["json"] = items
    item.ServeJSON()
}
