package storage

import (
	"context"
	"database/sql"

	"github.com/DasJalapa/reportes-salud/internal/models"
)

type repoAuthorization struct{}

// NewPersonStorage  constructor para userStorage
func NewAuthorizationStorage() AuthorizationStorage {
	return &repoAuthorization{}
}

type AuthorizationStorage interface {
	Create(ctx context.Context, authorization models.Authorization) (models.Authorization, error)

	GetManyWorkDependency(ctx context.Context) ([]models.WorkDependency, error)
	ManyJobs(ctx context.Context) ([]models.Job, error)
}

func (*repoAuthorization) Create(ctx context.Context, authorization models.Authorization) (models.Authorization, error) {
	authoriza := models.Authorization{}

	query := `
	INSERT INTO
    autorization (
        uuid,
        register,
        dateemmited,
        startdate,
        enddate,
		resumework,
		holidays,
        totaldays,
        pendingdays,
        observation,
        authorizationyear,
        person_idperson,
        partida,
        workdependency_uuid,
        user_uuid
    )
    VALUES(?,
		(
            SELECT Count(*) + 1 AS count FROM autorization a
        ),?,?,?,?,?,?,?,?,?,?,?,?,?);`
	trans, err := db.BeginTx(ctx, nil)

	if err != nil {
		return authoriza, err
	}
	defer trans.Rollback()

	_, err = db.QueryContext(
		ctx,
		query,
		authorization.UUIDAuthorization,
		authorization.Dateemmited,
		authorization.Startdate,
		authorization.Enddate,
		authorization.Resumework,
		authorization.Holidays,
		authorization.TotalDays,
		authorization.Pendingdays,
		authorization.Observation,
		authorization.Authorizationyear,
		authorization.UUID,
		authorization.Partida,
		authorization.Workdependency,
		authorization.User,
	)
	if err != nil {
		return authoriza, err
	}
	querySelect := `
	SELECT
    	a.register,
    	a.dateemmited,
    	a.startdate,
    	a.enddate,
		a.resumework,
		a.holidays,
    	a.totaldays,
    	a.pendingdays,
    	a.observation,
    	a.authorizationyear,
    	a.partida,
    	w.name as workdependency,
    	p.fullname,
    	p.cui,
    	j.name as job
	FROM
    	autorization a
    	INNER JOIN person p ON a.person_idperson = p.uuid
		INNER JOIN job j ON p.job_uuid = j.uuid
		INNER JOIN workdependency w ON a.workdependency_uuid = w.uuid
    	WHERE a.uuid = ?
	`
	err = trans.QueryRowContext(ctx, querySelect, authorization.UUIDAuthorization).Scan(
		&authoriza.Register,
		&authoriza.Dateemmited,
		&authoriza.Startdate,
		&authoriza.Enddate,
		&authoriza.Resumework,
		&authoriza.Holidays,
		&authoriza.TotalDays,
		&authoriza.Pendingdays,
		&authoriza.Observation,
		&authoriza.Authorizationyear,
		&authoriza.Partida,
		&authoriza.Workdependency,
		&authoriza.Fullname,
		&authoriza.CUI,
		&authoriza.Job,
	)

	if err != sql.ErrNoRows {
		return authoriza, err
	}

	if errtrans := trans.Commit(); errtrans != nil {
		return authoriza, errtrans
	}

	return authoriza, nil
}

func (*repoAuthorization) GetManyWorkDependency(ctx context.Context) ([]models.WorkDependency, error) {
	work := models.WorkDependency{}
	works := []models.WorkDependency{}

	query := "SELECT * FROM workdependency;"

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

func (*repoAuthorization) ManyJobs(ctx context.Context) ([]models.Job, error) {
	query := "SELECT uuid, name FROM job;"
	job := models.Job{}
	jobs := []models.Job{}

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return jobs, err
	}

	for rows.Next() {
		if err := rows.Scan(&job.UUIDJob, &job.Job); err != nil {
			return jobs, err
		}

		jobs = append(jobs, job)
	}

	return jobs, nil
}
