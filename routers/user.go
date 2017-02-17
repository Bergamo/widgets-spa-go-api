package routers

import (
	"github.com/Bergamo/widgets-spa-go-api/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
    "gopkg.in/mgo.v2"
)

func SetUserRoutes(router *mux.Router) *mux.Router {

    // Get a UserController instance
	userController := controllers.NewUserController(getSession())

	router.Handle("/",
		negroni.New(
			//negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(userController.Index),
		)).Methods("GET")

    router.Handle("/users",
		negroni.New(
			//negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(userController.GetUsers),
		)).Methods("GET")

    router.Handle("/users/{id}",
		negroni.New(
			//negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(userController.GetUser),
		)).Methods("GET")

    router.Handle("/users",
		negroni.New(
			//negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(userController.CreateUser),
		)).Methods("POST")

    router.Handle("/users/{id}",
		negroni.New(
			//negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(userController.UpdateUser),
		)).Methods("PUT")

    router.Handle("/users/{id}",
		negroni.New(
			//negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(userController.RemoveUser),
		)).Methods("DELETE")

	return router
}

// getSession creates a new mongo session and panics if connection error occurs
func getSession() *mgo.Session {
	// Connect to our local mongo
	session, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	// Deliver session
	return session
}