package storage

import (
	"context"
	"database/sql"

	"github.com/DasJalapa/reportes-salud/internal/lib"
	"github.com/DasJalapa/reportes-salud/internal/models"
)

// NewPersonStorage  constructor para userStorage
func NewPersonStorage() PersonStorage {
	return &repoPerson{
		limit:  10,
		offset: 1,
	}
}

type repoPerson struct {
	limit  int
	offset int
	// p PersonStorage
}

type PersonStorage interface {
	GetOne(ctx context.Context, uuid string) (models.Person, error)
	GetMany(ctx context.Context, filter string) ([]models.Person, error)

	Update(ctx context.Context, uuid string, person models.Person) (string, error)
	PaginationQuery(page, limit int) *repoPerson
}

func (*repoPerson) GetOne(ctx context.Context, uuid string) (models.Person, error) {
	person := models.Person{}

	query := `SELECT p.uuid, p.fullname, p.cui, p.job_uuid, j.name as job FROM person p
			  INNER JOIN job j ON p.job_uuid = j.uuid
			  WHERE p.uuid = ?;`

	rows := db.QueryRowContext(ctx, query, uuid).Scan(&person.UUID, &person.Fullname, &person.CUI, &person.JobUUUID, &person.Job)
	if rows == sql.ErrNoRows {
		return person, lib.ErrNotFound
	}

	return person, nil
}

func (p *repoPerson) GetMany(ctx context.Context, filter string) ([]models.Person, error) {
	person := models.Person{}
	persons := []models.Person{}

	var query string
	args := []interface{}{}

	if filter != "" {
		filter = "%" + filter + "%"
		query = `SELECT p.uuid, p.fullname, p.cui, j.name as job, j.uuid FROM person p
		INNER JOIN job j ON p.job_uuid = j.uuid
		WHERE fullname LIKE ? OR cui LIKE ?
		ORDER BY p.fullname ASC
		LIMIT ? OFFSET ?;`

		args = append(args, filter, filter, p.limit, p.offset)

	} else {
		query = `SELECT p.uuid, p.fullname, p.cui, j.name as job, j.uuid FROM person p
		INNER JOIN job j ON p.job_uuid = j.uuid
		ORDER BY p.fullname ASC
		LIMIT ? OFFSET ?;`
		args = append(args, p.limit, p.offset)
	}

	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return persons, err
	}

	for rows.Next() {
		err := rows.Scan(&person.UUID, &person.Fullname, &person.CUI, &person.Job, &person.JobUUUID)
		if err != nil {
			return persons, err
		}

		persons = append(persons, person)
	}

	return persons, nil
}

func (*repoPerson) Update(ctx context.Context, uuid string, person models.Person) (string, error) {

	query := "UPDATE person SET fullname = ?, cui = ?, job_uuid = ? "
	query += " WHERE uuid = ?;"

	_, err := db.QueryContext(ctx, query, person.Fullname, person.CUI, person.JobUUUID, uuid)
	if err != nil {
		return "", err
	}

	return person.UUID, nil
}

func (p *repoPerson) PaginationQuery(limit, page int) *repoPerson {
	if limit != 0 {
		p.limit = limit
	}

	if page >= 1 {
		p.offset = page - 1
	} else {
		p.offset = 0
	}

	return p
}
