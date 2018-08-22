package routers

import (
	"github.com/dsbell81/go-pgsql-api/controllers"
	"github.com/gorilla/mux"
)

func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/users/register", controllers.Register).Methods("POST")
	router.HandleFunc("/users/login", controllers.Login).Methods("POST")
	return router
}
