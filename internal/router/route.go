package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/DasJalapa/reportes-salud/internal/middleware"
)

// InitRoutes inicializa todas las rutas
func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.PathPrefix("/images").Handler(http.StripPrefix("/images", http.FileServer(http.Dir("public/"))))
	api := router.PathPrefix("/das/v1").Subrouter()
	api.Use(middleware.Auth)

	api = SetUserRoutes(api)
	api = SetPersonRoutes(api)
	api = SetAuthorizationRoutes(api)
	api = SetJobRoutes(api)
	api = SetWorkDependencyRoutes(api)

	router.Use(middleware.WriteJSONHeader)

	return router
}
