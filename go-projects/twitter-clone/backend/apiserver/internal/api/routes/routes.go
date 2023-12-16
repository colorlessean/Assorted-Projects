package routes

import (
	"github.com/colorlessean/Assorted-Projects/go-projects/twitter-clone/backend/apiserver/internal/api/handlers"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.RootHandler)

	return r
}
