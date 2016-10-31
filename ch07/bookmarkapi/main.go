package main

import (
	"log"
	"net/http"

	"github.com/urfave/negroni"

	"github.com/shijuvar/go-recipes/ch07/bookmarkapi/common"
	"github.com/shijuvar/go-recipes/ch07/bookmarkapi/routers"
)

// Entry point of the program
func main() {

	// Calls startup logic
	common.StartUp()
	// Get the mux router object
	router := routers.InitRoutes()
	// Create a negroni instance
	n := negroni.Classic()
	n.UseHandler(router)

	// Create the Server
	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}
	log.Println("Listening...")
	// Running the HTTP Server
	server.ListenAndServe()
}
