package service

import (
	"context"

	"github.com/DasJalapa/reportes-salud/internal/models"
	"github.com/DasJalapa/reportes-salud/internal/storage"
	"github.com/google/uuid"
)

type authorizationService struct{}

var AuthorizationStorage storage.AuthorizationStorage

// NewAuthorizationService retorna un nuevo servicio para los usuarios
func NewAuthorizationService(authorizationStorage storage.AuthorizationStorage) AuthorizationService {
	AuthorizationStorage = authorizationStorage
	return &authorizationService{}
}

type AuthorizationService interface {
	Create(ctx context.Context, auhtorization models.Authorization) (models.Authorization, error)
	GetManyWorkDependency(ctx context.Context) ([]models.WorkDependency, error)
}

func (*authorizationService) Create(ctx context.Context, authorization models.Authorization) (models.Authorization, error) {
	authorization.UUIDAuthorization = uuid.New().String()
	return AuthorizationStorage.Create(ctx, authorization)
}

func (*authorizationService) GetManyWorkDependency(ctx context.Context) ([]models.WorkDependency, error) {
	return AuthorizationStorage.GetManyWorkDependency(ctx)
}
