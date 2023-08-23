package main

import (
	"fmt"
	"log"

	"github.com/colorlessean/Assorted-Projects/go-projects/auth-service/internal/database"
	"github.com/colorlessean/Assorted-Projects/go-projects/auth-service/internal/login"
)

func main() {
    fmt.Println("auth service starting up...")

    db, err := database.NewDatabase(database.Memory) 
    if err != nil {
        log.Fatalf("failed to setup database: %v", err)
        return
    }

    loginServer, err := login.NewLoginServer(db)
    if err != nil {
        log.Fatalf("failed create server instance: %v", err)
        return
    }

    err = loginServer.Start()
    if err != nil {
        log.Fatalf("failed to start server: %v", err)
        return
    }
}

