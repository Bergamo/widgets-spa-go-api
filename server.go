package main

import (
	// Standard library packages
	"net/http"

	// Third party packages
	"github.com/Bergamo/widgets-spa-go-api/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	// Instantiate a new router
	router := httprouter.New()

	// Get a UserController instance
	userController := controllers.NewUserController(getSession())

	// Get a user resource
	router.GET("/", userController.Index)

	// Get a user resource
	router.GET("/user/:id", userController.GetUser)

	// Create a new user
	router.POST("/user", userController.CreateUser)

	// Remove an existing user
	router.DELETE("/user/:id", userController.RemoveUser)

	// Fire up the server
	http.ListenAndServe("localhost:3000", router)
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
