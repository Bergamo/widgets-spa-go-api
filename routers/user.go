package routers

import (
	"github.com/Bergamo/widgets-spa-go-api/common"
	"github.com/Bergamo/widgets-spa-go-api/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

// SetUserRoutes set user routes
func SetUserRoutes(router *mux.Router) *mux.Router {

	// Get a UserController instance
	userController := controllers.NewUserController(common.GetSession())
	tokenController := controllers.NewTokenController()

	router.Handle("/",
		negroni.New(
			negroni.HandlerFunc(userController.Index),
		)).Methods("GET")

	router.Handle("/users",
		negroni.New(
			negroni.HandlerFunc(tokenController.ValidateTokenHandler),
			negroni.HandlerFunc(userController.GetUsers),
		)).Methods("GET")

	router.Handle("/users/{id}",
		negroni.New(
			negroni.HandlerFunc(tokenController.ValidateTokenHandler),
			negroni.HandlerFunc(userController.GetUser),
		)).Methods("GET")

	router.Handle("/users",
		negroni.New(
			negroni.HandlerFunc(tokenController.ValidateTokenHandler),
			negroni.HandlerFunc(userController.CreateUser),
		)).Methods("POST")

	router.Handle("/users/{id}",
		negroni.New(
			negroni.HandlerFunc(tokenController.ValidateTokenHandler),
			negroni.HandlerFunc(userController.UpdateUser),
		)).Methods("PUT")

	router.Handle("/users/{id}",
		negroni.New(
			negroni.HandlerFunc(tokenController.ValidateTokenHandler),
			negroni.HandlerFunc(userController.RemoveUser),
		)).Methods("DELETE")

	return router
}
