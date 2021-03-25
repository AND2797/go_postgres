package main

import(
        "database/sql"
        "fmt"
        _ "github.com/lib/pq"
    )

func main() {

    connStr := "user=postgres dbname=postgres password=postgres host=localhost sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil{
        panic(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil{
        panic(err)
    }
    fmt.Println("successful!")
}
