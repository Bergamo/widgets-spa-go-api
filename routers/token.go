package routers

import (
	"github.com/Bergamo/widgets-spa-go-api/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

/* Set up a global string for our secret */
var mySigningKey = []byte("secret")

// SetTokenRoutes set user routes
func SetTokenRoutes(router *mux.Router) *mux.Router {

	// Get a UserController instance
	tokenController := controllers.NewTokenController(mySigningKey)

	router.Handle("/get-token",
		negroni.New(
			negroni.HandlerFunc(tokenController.GetTokenHandler),
		)).Methods("GET")

	return router
}
