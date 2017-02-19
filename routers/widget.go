package routers

import (
	"github.com/Bergamo/widgets-spa-go-api/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/Bergamo/widgets-spa-go-api/common"
)

// SetWidgetRoutes set user routes
func SetWidgetRoutes(router *mux.Router) *mux.Router {

	// Get a WidgetController instance
	widgetController := controllers.NewWidgetController(common.GetSession())
	tokenController := controllers.NewTokenController()

	router.Handle("/widgets",
		negroni.New(
			//negroni.HandlerFunc(tokenController.ValidateTokenHandler),
			negroni.HandlerFunc(widgetController.GetWidgets),
		)).Methods("GET")

	router.Handle("/widgets/{id}",
		negroni.New(
			negroni.HandlerFunc(tokenController.ValidateTokenHandler),
			negroni.HandlerFunc(widgetController.GetWidget),
		)).Methods("GET")

	router.Handle("/widgets",
		negroni.New(
			negroni.HandlerFunc(tokenController.ValidateTokenHandler),
			negroni.HandlerFunc(widgetController.CreateWidget),
		)).Methods("POST")

	router.Handle("/widgets/{id}",
		negroni.New(
			negroni.HandlerFunc(tokenController.ValidateTokenHandler),
			negroni.HandlerFunc(widgetController.UpdateWidget),
		)).Methods("PUT")

	router.Handle("/widgets/{id}",
		negroni.New(
			negroni.HandlerFunc(tokenController.ValidateTokenHandler),
			negroni.HandlerFunc(widgetController.RemoveWidget),
		)).Methods("DELETE")

	return router
}
