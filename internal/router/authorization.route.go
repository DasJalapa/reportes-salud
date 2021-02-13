package router

import (
	"github.com/gorilla/mux"

	"github.com/DasJalapa/reportes-salud/internal/controller"
	// "github.com/DasJalapa/reportes-salud/internal/middleware"
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

	authorization := router.PathPrefix("/authorization").Subrouter()
	// authorization.Use(middleware.AuthForAmdmin)
	authorization.HandleFunc("/emmit", authorizationController.Create).Methods("POST")
	authorization.HandleFunc("/works", authorizationController.GetManyWorks).Methods("GET")
	authorization.HandleFunc("/works", authorizationController.CreateWorkDependency).Methods("POST")
	authorization.HandleFunc("/jobs", authorizationController.ManyJobs).Methods("GET")
	authorization.HandleFunc("/jobs", authorizationController.CreateJob).Methods("POST")
	authorization.HandleFunc("/many", authorizationController.GetManyAuthorizations).Methods("GET")
	authorization.HandleFunc("/one/{uuid}", authorizationController.GetOnlyAuthorization).Methods("GET")
	authorization.HandleFunc("/update/{uuid}", authorizationController.UpdateAuthorization).Methods("PUT")
	authorization.HandleFunc("/pdfauthorization/{uuid}", authorizationController.GetOnlyAuthorizationPDF).Methods("GET")

	return router
}
