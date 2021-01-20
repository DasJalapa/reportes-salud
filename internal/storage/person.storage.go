package storage

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Mike2397/reportes-salud/internal/lib"
	"github.com/Mike2397/reportes-salud/internal/models"
)

// NewPersonStorage  constructor para userStorage
func NewPersonStorage() PersonStorage {
	return &repoPerson{}
}

type repoPerson struct{}

type PersonStorage interface {
	GetOne(ctx context.Context, uuid string) (models.Person, error)
	GetMany(ctx context.Context, filter string, limit int) ([]models.Person, error)
}

func (*repoPerson) GetOne(ctx context.Context, uuid string) (models.Person, error) {
	person := models.Person{}
	fmt.Println(uuid)

	query := `SELECT p.uuid, p.fullname, p.cui, j.name as job FROM person p
			  INNER JOIN job j ON p.job_uuid = j.uuid
			  WHERE p.uuid = ?;`

	rows := db.QueryRowContext(ctx, query, uuid).Scan(&person.UUID, &person.Fullname, &person.CUI, &person.Job)
	if rows == sql.ErrNoRows {
		return person, lib.ErrNotFound
	}

	return person, nil
}

func (*repoPerson) GetMany(ctx context.Context, filter string, limit int) ([]models.Person, error) {
	person := models.Person{}
	persons := []models.Person{}

	if limit == 0 {
		limit = 10
	}

	var query string
	args := []interface{}{}

	if filter != "" {
		query = `SELECT p.uuid, p.fullname, p.cui, j.name as job FROM person p
		INNER JOIN job j ON p.job_uuid = j.uuid
		WHERE fullname LIKE ? OR cui LIKE ?
		LIMIT ?;`

		args = append(args, filter, filter, limit)

	} else {
		query = `SELECT p.uuid, p.fullname, p.cui, j.name as job FROM person p
		INNER JOIN job j ON p.job_uuid = j.uuid
		limit ?`
		args = append(args, limit)
	}

	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return persons, err
	}

	for rows.Next() {
		err := rows.Scan(&person.UUID, &person.Fullname, &person.CUI, &person.Job)
		if err != nil {
			return persons, err
		}

		persons = append(persons, person)
	}

	return persons, nil
}
