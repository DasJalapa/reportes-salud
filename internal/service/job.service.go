package service

import (
	"context"

	"github.com/DasJalapa/reportes-salud/internal/models"
	"github.com/DasJalapa/reportes-salud/internal/storage"
	"github.com/google/uuid"
)

type jobService struct{}

var JobStorage storage.IJobStorage

// NewAuthorizationService retorna un nuevo servicio para los usuarios
func NewJobService(jobStorage storage.IJobStorage) IJobService {
	JobStorage = jobStorage
	return &jobService{}
}

type IJobService interface {
	ManyJobs(ctx context.Context) ([]models.Job, error)
	OneJob(ctx context.Context, uuid string) (models.Job, error)
	CreateJob(ctx context.Context, job models.Job) (string, error)
	Delete(ctx context.Context, uuid string) error
	Update(ctx context.Context, job models.Job, uuid string) (models.Job, error)
}

func (*jobService) CreateJob(ctx context.Context, job models.Job) (string, error) {
	job.UUIDJob = uuid.New().String()
	return JobStorage.CreateJob(ctx, job)
}

func (*jobService) ManyJobs(ctx context.Context) ([]models.Job, error) {
	return JobStorage.ManyJobs(ctx)
}

func (*jobService) Delete(ctx context.Context, uuid string) error {
	return JobStorage.Delete(ctx, uuid)
}

func (*jobService) Update(ctx context.Context, job models.Job, uuid string) (models.Job, error) {
	return JobStorage.Update(ctx, job, uuid)
}

func (*jobService) OneJob(ctx context.Context, uuid string) (models.Job, error) {
	return JobStorage.OneJob(ctx, uuid)
}
