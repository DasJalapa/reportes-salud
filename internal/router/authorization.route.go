package router

import (
	"github.com/gorilla/mux"

	"github.com/DasJalapa/reportes-salud/internal/controller"
	"github.com/DasJalapa/reportes-salud/internal/middleware"
	"github.com/DasJalapa/reportes-salud/internal/service"
	"github.com/DasJalapa/reportes-salud/internal/storage"
)

var (
	authorizationStorage    storage.AuthorizationStorage       = storage.NewAuthorizationStorage()
	authorizationService    service.AuthorizationService       = service.NewAuthorizationService(authorizationStorage)
	authorizationController controller.AuthorizationController = controller.NewAuthorizationController(authorizationService)
)

// SetAuthorizationRoutes registra la rutas a usar para los controladires de usuario
func SetAuthorizationRoutes(router *mux.Router) *mux.Router {

	user := router.PathPrefix("/authorization").Subrouter()
	user.Use(middleware.AuthForAmdmin)
	user.HandleFunc("/emmit", authorizationController.Create).Methods("POST")
	user.HandleFunc("/works", authorizationController.GetManyWorks).Methods("GET")
	user.HandleFunc("/jobs", authorizationController.ManyJobs).Methods("GET")

	return router
}
