package storage

import (
	"context"
	"database/sql"

	"github.com/DasJalapa/reportes-salud/internal/lib"
	"github.com/DasJalapa/reportes-salud/internal/models"
)

func NewWorkDependencyStorage() IWorkDependencyStorage {
	return &repoWorkDependency{}
}

type repoWorkDependency struct {
}

type IWorkDependencyStorage interface {
	GetManyWorkDependency(ctx context.Context) ([]models.WorkDependency, error)
	CreateWorkDependency(ctx context.Context, dependency models.WorkDependency) (string, error)
	OneWorkDependency(ctx context.Context, uuid string) (models.WorkDependency, error)
	Delete(ctx context.Context, uuid string) error
	Update(ctx context.Context, workdependency models.WorkDependency, uuid string) (models.WorkDependency, error)
}

func (*repoWorkDependency) GetManyWorkDependency(ctx context.Context) ([]models.WorkDependency, error) {
	work := models.WorkDependency{}
	works := []models.WorkDependency{}

	query := "SELECT uuid, name FROM workdependency ORDER BY name ASC;"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return works, err
	}

	for rows.Next() {
		err := rows.Scan(&work.Uuidwork, &work.Name)
		if err != nil {
			return works, err
		}

		works = append(works, work)
	}
	return works, nil
}

func (*repoWorkDependency) CreateWorkDependency(ctx context.Context, dependency models.WorkDependency) (string, error) {
	query := "INSERT INTO workdependency VALUES (?, ?)"

	_, err := db.QueryContext(ctx, query, dependency.Uuidwork, dependency.Name)
	if err != nil {
		return "", err
	}

	return dependency.Uuidwork, nil
}

func (*repoWorkDependency) OneWorkDependency(ctx context.Context, uuid string) (models.WorkDependency, error) {
	query := "SELECT uuid, name FROM workdependency WHERE uuid = ?;"
	wd := models.WorkDependency{}

	err := db.QueryRowContext(ctx, query, uuid).Scan(&wd.Uuidwork, &wd.Name)

	if err == sql.ErrNoRows {
		return wd, lib.ErrNotFound
	}

	if err != nil {
		return wd, err
	}

	return wd, nil
}

func (*repoWorkDependency) Delete(ctx context.Context, uuid string) error {

	queryUpdate := "DELETE FROM workdependency WHERE uuid = ?;"

	rows, err := db.ExecContext(ctx, queryUpdate, uuid)
	if err != nil {
		return lib.ExtractMysqlError(err)
	}

	resultDelete, _ := rows.RowsAffected()
	if resultDelete == 0 {
		return lib.ErrNotFound
	}

	return nil
}

func (*repoWorkDependency) Update(ctx context.Context, workdependency models.WorkDependency, uuid string) (models.WorkDependency, error) {
	queryVerify := "SELECT COUNT(name) FROM workdependency WHERE uuid = ?;"
	var isInDB int
	err := db.QueryRowContext(ctx, queryVerify, uuid).Scan(&isInDB)
	if err != nil {
		return workdependency, err
	}

	if isInDB == 0 {
		return workdependency, lib.ErrNotFound
	}

	queryUpdate := "UPDATE workdependency SET name = ? WHERE uuid = ?;"
	_, err = db.ExecContext(ctx, queryUpdate, workdependency.Name, uuid)
	if err != nil {
		return workdependency, err
	}

	workdependency.Uuidwork = uuid
	return workdependency, nil
}
