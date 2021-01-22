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
	GetManyAuthorizations(ctx context.Context) ([]models.Authorization, error)
	GetOnlyAuthorization(ctx context.Context, uuid string) (models.Authorization, error)
	UpdateAuthorization(ctx context.Context, authorization models.Authorization, uuid string) (models.Authorization, error)

	GetManyWorkDependency(ctx context.Context) ([]models.WorkDependency, error)
	CreateWorkDependency(ctx context.Context, dependency models.WorkDependency) (string, error)

	ManyJobs(ctx context.Context) ([]models.Job, error)
	CreateJob(ctx context.Context, job models.Job) (string, error)
}

func (*repoAuthorization) GetManyAuthorizations(ctx context.Context) ([]models.Authorization, error) {
	autorization := models.Authorization{}
	autorizations := []models.Authorization{}

	query := `SELECT a.uuid, a.register, a.dateemmited, p.fullname, p.cui FROM autorization a 
			  INNER JOIN person p ON a.person_idperson = p.uuid;`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return autorizations, err
	}

	for rows.Next() {
		if err := rows.Scan(&autorization.UUIDAuthorization, &autorization.Register, &autorization.Dateemmited, &autorization.Fullname, &autorization.CUI); err != nil {
			return autorizations, err
		}

		autorizations = append(autorizations, autorization)
	}

	return autorizations, nil
}

func (*repoAuthorization) GetOnlyAuthorization(ctx context.Context, uuid string) (models.Authorization, error) {
	autorization := models.Authorization{}

	query := `
	SELECT 
		a.uuid,
		a.person_idperson,
		a.workdependency_uuid,
		a.register,
		p.cui,
		p.fullname,
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
		j.name as job
	FROM autorization a
	INNER JOIN person p ON a.person_idperson = p.uuid
	INNER JOIN workdependency w ON a.workdependency_uuid = w.uuid
	INNER JOIN job j ON p.job_uuid = j.uuid
	WHERE a.uuid = ?;`

	err := db.QueryRowContext(ctx, query, uuid).Scan(
		&autorization.UUIDAuthorization,
		&autorization.UUID,
		&autorization.WorkdependencyUUID,
		&autorization.Register,
		&autorization.CUI,
		&autorization.Fullname,
		&autorization.Dateemmited,
		&autorization.Startdate,
		&autorization.Enddate,
		&autorization.Resumework,
		&autorization.Holidays,
		&autorization.TotalDays,
		&autorization.Pendingdays,
		&autorization.Observation,
		&autorization.Authorizationyear,
		&autorization.Partida,
		&autorization.Workdependency,
		&autorization.Job,
	)

	if err != nil {
		return autorization, err
	}

	return autorization, nil
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

	authoriza, err = DataPDFAuthorization(ctx, authorization.UUIDAuthorization, db)
	if err != nil {
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

func (*repoAuthorization) CreateWorkDependency(ctx context.Context, dependency models.WorkDependency) (string, error) {
	query := "INSERT INTO workdependency VALUES (?, ?)"

	_, err := db.QueryContext(ctx, query, dependency.Uuidwork, dependency.Name)
	if err != nil {
		return "", err
	}

	return dependency.Uuidwork, nil
}

func (*repoAuthorization) CreateJob(ctx context.Context, job models.Job) (string, error) {
	query := "INSERT INTO job (uuid, name) VALUES (?, ?)"

	_, err := db.QueryContext(ctx, query, job.UUIDJob, job.Job)
	if err != nil {
		return "", err
	}

	return job.UUIDJob, nil
}

func (*repoAuthorization) UpdateAuthorization(ctx context.Context, authorization models.Authorization, uuid string) (models.Authorization, error) {
	authoriza := models.Authorization{}
	trans, err := db.BeginTx(ctx, nil)

	if err != nil {
		return authoriza, err
	}
	defer trans.Rollback()

	query := `
		UPDATE autorization SET
			dateemmited = ?,
			startdate = ?,
			enddate = ?,
			resumework = ?,
			holidays = ?,
			totaldays = ?,
			pendingdays = ?,
			observation = ?,
			authorizationyear = ?,
			partida = ?,
			workdependency_uuid = ?
		WHERE uuid = ?;`

	_, err = db.QueryContext(ctx, query,
		authorization.Dateemmited,
		authorization.Startdate,
		authorization.Enddate,
		authorization.Resumework,
		authorization.Holidays,
		authorization.TotalDays,
		authorization.Pendingdays,
		authorization.Observation,
		authorization.Authorizationyear,
		authorization.Partida,
		authorization.WorkdependencyUUID,
		uuid,
	)
	if err != nil {
		return authoriza, err
	}

	authoriza, err = DataPDFAuthorization(ctx, uuid, db)

	if err != nil {
		return authoriza, err
	}

	if errtrans := trans.Commit(); errtrans != nil {
		return authoriza, errtrans
	}

	return authoriza, nil
}

func DataPDFAuthorization(ctx context.Context, UUIDAuthorization string, trans *sql.DB) (models.Authorization, error) {
	authoriza := models.Authorization{}
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
	err := db.QueryRowContext(ctx, querySelect, UUIDAuthorization).Scan(
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

	if err != nil {
		return authoriza, err
	}

	return authoriza, nil
}
