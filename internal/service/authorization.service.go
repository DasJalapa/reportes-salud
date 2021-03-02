package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/DasJalapa/reportes-salud/internal/models"
	"github.com/DasJalapa/reportes-salud/internal/mysql"
	"github.com/DasJalapa/reportes-salud/internal/storage"
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
	GetManyAuthorizations(ctx context.Context) ([]models.Authorization, error)
	GetOnlyAuthorization(ctx context.Context, uuid string) (models.Authorization, error)
	UpdateAuthorization(ctx context.Context, authorization models.Authorization, uuid string) (models.Authorization, error)
	GetOnlyAuthorizationPDF(ctx context.Context, UUIDAuthorization string) (models.Authorization, error)
}

func (*authorizationService) Create(ctx context.Context, authorization models.Authorization) (models.Authorization, error) {
	authorization.UUIDAuthorization = uuid.New().String()
	return AuthorizationStorage.Create(ctx, authorization)
}

func (*authorizationService) GetManyAuthorizations(ctx context.Context) ([]models.Authorization, error) {
	return AuthorizationStorage.GetManyAuthorizations(ctx)
}

func (*authorizationService) GetOnlyAuthorization(ctx context.Context, uuid string) (models.Authorization, error) {
	return AuthorizationStorage.GetOnlyAuthorization(ctx, uuid)
}

func (*authorizationService) UpdateAuthorization(ctx context.Context, authorization models.Authorization, uuid string) (models.Authorization, error) {
	return AuthorizationStorage.UpdateAuthorization(ctx, authorization, uuid)
}

func (*authorizationService) GetOnlyAuthorizationPDF(ctx context.Context, UUIDAuthorization string) (models.Authorization, error) {
	return storage.DataPDFAuthorization(ctx, UUIDAuthorization, mysql.Connect())
}
