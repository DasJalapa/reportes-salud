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

	ManyUsers(ctx context.Context) ([]models.User, error)
	Roles(ctx context.Context) ([]models.Rol, error)
	ChangePassword(ctx context.Context, uuidUser, actualPassword, newPassword string) error
}

// UserCreate es el servicio de conexion al storage de crear usuario
func (*userService) Create(ctx context.Context, user *models.User) (string, error) {
	user.ID = uuid.New().String()

	return Userstorage.Create(ctx, user)
}

// UserLogin es el servicio de conexion al storage de login de usuario
func (*userService) Login(ctx context.Context, user *models.User) (models.User, error) {
	return Userstorage.Login(ctx, user)
}

func (*userService) ManyUsers(ctx context.Context) ([]models.User, error) {
	return Userstorage.GetManyUsers(ctx)
}

func (*userService) Roles(ctx context.Context) ([]models.Rol, error) {
	return Userstorage.Roles(ctx)
}

func (*userService) ChangePassword(ctx context.Context, uuidUser, actualPassword, newPassword string) error {
	return Userstorage.ChangePassword(ctx, uuidUser, actualPassword, newPassword)
}
