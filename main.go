package main

import (
	"log"
	"net/http"

	"github.com/dsbell81/go-pgsql-api/datastore"
	"github.com/dsbell81/go-pgsql-api/routers"
	"github.com/dsbell81/go-pgsql-api/utils"
	"github.com/urfave/negroni"
)

func main() {

	//initialize datastore connection
	datastore.InitDb()

	// Get the mux router object
	router := routers.InitRoutes()

	// Create a negroni instance
	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr:    utils.AppConfig.Server,
		Handler: n,
	}

	log.Println("Listening...")
	server.ListenAndServe()

}
