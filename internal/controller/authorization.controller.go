package controller

import (
	"encoding/json"
	"net/http"

	"github.com/DasJalapa/reportes-salud/internal/lib"
	"github.com/DasJalapa/reportes-salud/internal/middleware"
	"github.com/DasJalapa/reportes-salud/internal/models"
	"github.com/DasJalapa/reportes-salud/internal/service"
)

type authorizationController struct{}

var AuthorizationService service.AuthorizationService

// NewAuthorizationController retorna un nuevo controller de tipo usuario controller
func NewAuthorizationController(authorizationService service.AuthorizationService) AuthorizationController {
	AuthorizationService = authorizationService
	return &authorizationController{}
}

// AuthorizationController contiene todos los controladores de usuario
type AuthorizationController interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetManyWorks(w http.ResponseWriter, r *http.Request)
}

func (*authorizationController) Create(w http.ResponseWriter, r *http.Request) {
	id, ok := middleware.IsAuthenticated(r.Context())
	if !ok {
		respond(w, response{Message: lib.ErrUnauthenticated.Error()}, http.StatusUnauthorized)
		return
	}

	defer r.Body.Close()
	var authorization models.Authorization

	if err := json.NewDecoder(r.Body).Decode(&authorization); err != nil {
		respond(w, response{
			Ok:      false,
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	authorization.User = id

	result, err := AuthorizationService.Create(r.Context(), authorization)
	if err == nil {
		respond(w, response{
			Ok:      true,
			Message: "Registro creado satisfactoriamente",
			Data:    result,
		}, http.StatusCreated)
		return
	}

	if err != nil {
		respondError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (*authorizationController) GetManyWorks(w http.ResponseWriter, r *http.Request) {
	_, ok := middleware.IsAuthenticated(r.Context())
	if !ok {
		respond(w, response{Message: lib.ErrUnauthenticated.Error()}, http.StatusUnauthorized)
		return
	}

	data, err := AuthorizationService.GetManyWorkDependency(r.Context())
	if err == lib.ErrNotFound {
		respond(w, response{
			Ok:      false,
			Data:    data,
			Message: lib.ErrNotFound.Error(),
		}, http.StatusOK)
		return
	}

	if err == nil {
		respond(w, response{
			Ok:   true,
			Data: data,
		}, http.StatusOK)
		return
	}

	if err != nil {
		respondError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
