package main

import (
	"log"
	"net/http"

	"github.com/Bergamo/widgets-spa-go-api/common"
	"github.com/Bergamo/widgets-spa-go-api/routers"
	"github.com/codegangsta/negroni"
)

func main() {

	// Instantiate a new router
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	log.Printf("Listening on : %s\n", common.AppConfig.Server)
	http.ListenAndServe(common.AppConfig.Server, n)
}
