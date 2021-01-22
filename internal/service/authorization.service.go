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

	GetManyWorkDependency(ctx context.Context) ([]models.WorkDependency, error)
	CreateWorkDependency(ctx context.Context, dependency models.WorkDependency) (string, error)

	ManyJobs(ctx context.Context) ([]models.Job, error)
	CreateJob(ctx context.Context, job models.Job) (string, error)
}

func (*authorizationService) Create(ctx context.Context, authorization models.Authorization) (models.Authorization, error) {
	authorization.UUIDAuthorization = uuid.New().String()
	return AuthorizationStorage.Create(ctx, authorization)
}

func (*authorizationService) GetManyWorkDependency(ctx context.Context) ([]models.WorkDependency, error) {
	return AuthorizationStorage.GetManyWorkDependency(ctx)
}

func (*authorizationService) ManyJobs(ctx context.Context) ([]models.Job, error) {
	return AuthorizationStorage.ManyJobs(ctx)
}

func (*authorizationService) CreateWorkDependency(ctx context.Context, dependency models.WorkDependency) (string, error) {
	dependency.Uuidwork = uuid.New().String()
	return AuthorizationStorage.CreateWorkDependency(ctx, dependency)
}

func (*authorizationService) CreateJob(ctx context.Context, job models.Job) (string, error) {
	job.UUIDJob = uuid.New().String()
	return AuthorizationStorage.CreateJob(ctx, job)
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
