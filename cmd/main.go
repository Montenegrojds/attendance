package main

import (
    "fmt"
    "log"
    "github.com/Montenegrojds/attendance/internal/db"
	"net/http"
    "github.com/Montenegrojds/attendance/internal/handlers"
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
	mux := http.NewServeMux()
	myServer :=&http.Server{
		Addr: ":8080",
		Handler: mux,
	}
	cfg := handlers.ApiConfig{DB: database}
	mux.HandleFunc("GET /health",cfg.HandlerHealth)
	log.Fatal(myServer.ListenAndServe())
}
