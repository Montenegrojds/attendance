package main

import (
    "fmt"
    "log"
    "github.com/Montenegrojds/attendance/internal/db"
)

func main() {
    database, err := db.Init()
    if err != nil {
        log.Fatal(err)
    }
    err = db.CreateTables(database)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Database ready!")
}