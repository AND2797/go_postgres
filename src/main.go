package main

import (
        "gorm.io/driver/postgres"
        "gorm.io/gorm"
    )
func main() {
    dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=EST"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(err.Error())
    }


    database := db.DB()

    err := database.Ping()

    if err != nil{
        panic(err.Error())
    }

}
