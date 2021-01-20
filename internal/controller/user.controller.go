package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Mike2397/reportes-salud/internal/lib"
	"github.com/Mike2397/reportes-salud/internal/models"
	"github.com/Mike2397/reportes-salud/internal/service"
)

type userController struct{}

var userService service.UserService

// NewUserController retorna un nuevo controller de tipo usuario controller
func NewUserController(service service.UserService) UserController {
	userService = service
	return &userController{}
}

// UserController contiene todos los controladores de usuario
type UserController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

func (*userController) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respond(w, response{
			Ok:      false,
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	idinsert, err := userService.Create(r.Context(), &user)

	if err == lib.ErrDuplicateUser {
		respond(w, response{
			Ok:      false,
			Message: lib.ErrDuplicateUser.Error(),
		}, http.StatusConflict)
		return
	}

	if err == nil {
		respond(w, response{
			Ok:       true,
			Message:  "Usuario creado satisfactoriamente",
			IDInsert: idinsert,
		}, http.StatusCreated)
		return
	}

	if err != nil {
		respondError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (*userController) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	resp, err := userService.Login(r.Context(), &user)

	if err == lib.ErrUserNotFound {
		respond(w, response{
			Ok:      false,
			Message: lib.ErrUserNotFound.Error(),
		}, http.StatusNotFound)
		return
	}

	if err == nil {
		respond(w, resp, http.StatusOK)
		return
	}

	if err != nil {
		respondError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// func (*userController) Update(w http.ResponseWriter, r *http.Request) {
// 	_, ok := middleware.IsAuthenticated(r.Context())
// 	if !ok {
// 		respond(w, response{Message: lib.ErrUnauthenticated.Error()}, http.StatusUnauthorized)
// 		return
// 	}

// 	defer r.Body.Close()
// 	var user models.User

// 	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
// 		respond(w, response{
// 			Ok:      false,
// 			Message: err.Error(),
// 		}, http.StatusBadRequest)
// 		return
// 	}

// 	vars := mux.Vars(r)

// 	err := userService.Update(r.Context(), vars["id"], user.Rol)
// 	if err == nil {
// 		respond(w, response{
// 			Ok:      true,
// 			Message: "Usuario actualizado satisfactoriamente",
// 		}, http.StatusAccepted)
// 		return
// 	}

// 	if err != nil {
// 		respondError(w, err)
// 		return
// 	}
// }

// func (*userController) GetManyUsers(w http.ResponseWriter, r *http.Request) {
// 	var users []models.User
// 	data, err := userService.GetManyUsers(r.Context())
// 	if err == lib.ErrNotFound {
// 		respond(w, response{
// 			Ok:   false,
// 			Data: users,
// 		}, http.StatusNotFound)
// 	}

// 	if err == nil {
// 		respond(w, response{
// 			Ok:   true,
// 			Data: data,
// 		}, http.StatusOK)
// 		return
// 	}

// 	if err != nil {
// 		respondError(w, err)
// 		return
// 	}

// 	w.WriteHeader(http.StatusNoContent)
// }

// func (*userController) Rols(w http.ResponseWriter, r *http.Request) {

// 	data, err := userService.Roles(r.Context())
// 	if err == lib.ErrNotFound {
// 		respond(w, response{
// 			Ok:   false,
// 			Data: data,
// 		}, http.StatusNotFound)
// 	}

// 	if err == nil {
// 		respond(w, response{
// 			Ok:   true,
// 			Data: data,
// 		}, http.StatusOK)
// 		return
// 	}

// 	if err != nil {
// 		respondError(w, err)
// 		return
// 	}

// 	w.WriteHeader(http.StatusNoContent)
// }
