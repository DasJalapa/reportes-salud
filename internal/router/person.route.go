package router

import (
	"github.com/DasJalapa/reportes-salud/internal/controller"
	"github.com/DasJalapa/reportes-salud/internal/middleware"
	"github.com/DasJalapa/reportes-salud/internal/service"
	"github.com/DasJalapa/reportes-salud/internal/storage"
	"github.com/gorilla/mux"
)

var (
	personStorage    storage.PersonStorage       = storage.NewPersonStorage()
	personService    service.PersonService       = service.NewPersonService(personStorage)
	personController controller.PersonController = controller.NewPersonController(personService)
)

// SetPersonRoutes registra la rutas a usar para los controladires de usuario
func SetPersonRoutes(router *mux.Router) *mux.Router {

	person := router.PathPrefix("/persons").Subrouter()
	person.Use(middleware.AuthForAmdmin)
	person.HandleFunc("/{uuid}", personController.GetOne).Methods("GET")
	person.HandleFunc("/{limit}/{filter}", personController.GetMany).Methods("GET")
	person.HandleFunc("/{uuid}", personController.Update).Methods("PUT")

	return router
}
