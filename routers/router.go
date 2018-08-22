package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	// Routes for Accounts
	router = SetUserRoutes(router)

	return router
}
