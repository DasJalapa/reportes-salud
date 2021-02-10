package service

import (
	"context"

	"github.com/DasJalapa/reportes-salud/internal/models"
	"github.com/DasJalapa/reportes-salud/internal/storage"
)

type personService struct {
}

var PersonStorage storage.PersonStorage

// NewPersonService retorna un nuevo servicio para los usuarios
func NewPersonService(personStorage storage.PersonStorage) PersonService {
	PersonStorage = personStorage
	return &personService{}
}

// PersonService implementa el conjunto de metodos de servicio para usuario
type PersonService interface {
	GetOne(ctx context.Context, uuid string) (models.Person, error)
	GetMany(ctx context.Context, filter string, limit, page int) ([]models.Person, error)

	Update(ctx context.Context, uuid string, person models.Person) (string, error)
}

func (*personService) GetOne(ctx context.Context, uuid string) (models.Person, error) {
	return PersonStorage.GetOne(ctx, uuid)
}
func (*personService) GetMany(ctx context.Context, filter string, limit, page int) ([]models.Person, error) {
	return PersonStorage.PaginationQuery(limit, page).GetMany(ctx, filter)
}

func (*personService) Update(ctx context.Context, uuid string, person models.Person) (string, error) {
	return PersonStorage.Update(ctx, uuid, person)
}
