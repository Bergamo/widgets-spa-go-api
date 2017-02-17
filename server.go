package main

import (
	// Standard library packages
	"net/http"

	// Third party packages
	"github.com/Bergamo/widgets-spa-go-api/routers"
	"github.com/codegangsta/negroni"
)

func main() {

	// Instantiate a new router
	//router := httprouter.New()

	// Instantiate a new router
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe("localhost:3000", n)
	

	// Get a UserController instance
	//userController := controllers.NewUserController(getSession())

	// Get a user resource
	//router.GET("/", userController.Index)

	// Get a user resource
	//router.GET("/users", userController.GetUsers)

	// Get a user resource
	//router.GET("/user/:id", userController.GetUser)

	// Create a new user
	//router.POST("/user", userController.CreateUser)

	// Remove an existing user
	//router.DELETE("/user/:id", userController.RemoveUser)
}


