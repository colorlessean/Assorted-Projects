package main

import (
	"log"
	"net/http"

	"github.com/colorlessean/Assorted-Projects/go-projects/twitter-clone/backend/apiserver/internal/api/routes"
)

func main() {
	router := routes.SetupRoutes()

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Printf("Server is listening on port 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
