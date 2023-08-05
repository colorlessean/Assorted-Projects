package main

import (
	"fmt"

	"github.com/colorlessean/Assorted-Projects/go-projects/auth-service/internal/login"
)

func main() {
    fmt.Println("auth service starting up...")

    login.NewLoginServer()
}

