package routers

import (
	"github.com/bockbone/podtask/controllers"
	"github.com/gorilla/mux"
)

func SetUserRoutes(router *mux.Router) *mux.Router {

	router.HandleFunc("/api/v1/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/api/v1/users", controllers.CreateUser).Methods("POST")

	return router
}
