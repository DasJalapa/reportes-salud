package controller

import (
	"net/http"
	"strconv"

	"github.com/DasJalapa/reportes-salud/internal/lib"
	"github.com/DasJalapa/reportes-salud/internal/middleware"
	"github.com/DasJalapa/reportes-salud/internal/service"
	"github.com/gorilla/mux"
)

type personController struct{}

var PersonService service.PersonService

// NewPersonController retorna un nuevo controller de tipo usuario controller
func NewPersonController(personService service.PersonService) PersonController {
	PersonService = personService
	return &personController{}
}

// PersonController contiene todos los controladores de usuario
type PersonController interface {
	GetOne(w http.ResponseWriter, r *http.Request)
	GetMany(w http.ResponseWriter, r *http.Request)
}

func (*personController) GetOne(w http.ResponseWriter, r *http.Request) {
	_, ok := middleware.IsAuthenticated(r.Context())
	if !ok {
		respond(w, response{Message: lib.ErrUnauthenticated.Error()}, http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)

	data, err := PersonService.GetOne(r.Context(), vars["uuid"])
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
func (*personController) GetMany(w http.ResponseWriter, r *http.Request) {
	_, ok := middleware.IsAuthenticated(r.Context())
	if !ok {
		respond(w, response{Message: lib.ErrUnauthenticated.Error()}, http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	limit, err := strconv.Atoi(vars["limit"])
	if err != nil {
		respond(w, response{
			Ok:      false,
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	data, err := PersonService.GetMany(r.Context(), vars["filter"], limit)
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
