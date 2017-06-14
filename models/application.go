package models

import (
    "github.com/astaxie/beego"
    "database/sql"
    "fmt"
)


type Application struct {
    Id int64
    Title string
}

func GetFrist10() (result []Application, err error){
    var rows *sql.Rows
    rows, err = DB.Query("select id, title from Application order by id desc limit 10")
    if err != nil {
        beego.Critical("Query failed: ", err)
        return
    }

    result = []Application{}
    for rows.Next() {
        app := Application{}
        if err = rows.Scan(&app.Id, &app.Title); err != nil {
            beego.Critical("Scanning data error: ", err)
            return
        }
        beego.Info(app)
        result = append(result, app)
    }
    if err = rows.Err(); err != nil {
        beego.Error("Reading data error: ", err)
        return
    }

    return
}

func GetOneById(id int) (result Application, err error){
    fmt.Println("Scan sql.row by ScanStruct()")
    row := DB.QueryRow("select id, title from Application where id=?", id)
    ScanStruct(row, &result)
    return
}