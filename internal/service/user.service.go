package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/DasJalapa/reportes-salud/internal/models"
	"github.com/DasJalapa/reportes-salud/internal/storage"
)

type userService struct{}

var Userstorage storage.UserStorage

// NewUserService retorna un nuevo servicio para los usuarios
func NewUserService(userstorage storage.UserStorage) UserService {
	Userstorage = userstorage
	return &userService{}
}

// UserService implementa el conjunto de metodos de servicio para usuario
type UserService interface {
	Create(ctx context.Context, user *models.User) (string, error)
	Login(ctx context.Context, user *models.User) (models.User, error)
}

// UserCreate es el servicio de conexion al storage de crear usuario
func (*userService) Create(ctx context.Context, user *models.User) (string, error) {
	user.ID = uuid.New().String()
	user.IDRol = 2

	return Userstorage.Create(ctx, user)
}

// UserLogin es el servicio de conexion al storage de login de usuario
func (*userService) Login(ctx context.Context, user *models.User) (models.User, error) {
	return Userstorage.Login(ctx, user)
}
