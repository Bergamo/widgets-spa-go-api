package routers

import (
	"github.com/gorilla/mux"
)

// InitRoutes initiate routes
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetUserRoutes(router)
	router = SetWidgetRoutes(router)
	router = SetTokenRoutes(router)
	return router
}
