package models

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/astaxie/beego"
    "reflect"
    "fmt"
    "errors"
)

var DB *sql.DB

func init(){
    sqlString := beego.AppConfig.String("mysqlString")
    beego.Info("Connecting MySQL at ", sqlString)

    db, err := sql.Open("mysql", sqlString)
    if err != nil {
        beego.Critical("DB connection failed : ", err)
    }

    err = db.Ping()
    if err != nil {
        beego.Critical("DB Ping failed : ", err)
    }
    beego.Info("DB connected.")

    DB = db
}


func ScanStruct(row *sql.Row, sp interface{}) error {
    vp := reflect.ValueOf(sp)
    if vp.Kind() != reflect.Ptr {
        return errors.New("Must be a Pointer.")
    }
    v := vp.Elem()

    fieldPointers := []interface{}{}
    for i:=0; i < v.NumField(); i++ {
        f := v.Field(i)
        fp := f.Addr().Interface()
        fieldPointers = append(fieldPointers, fp)
    }
    fmt.Println(fieldPointers, len(fieldPointers))

    row.Scan(fieldPointers...)

    return nil
}