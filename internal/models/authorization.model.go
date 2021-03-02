package models

import "github.com/gobuffalo/nulls"

// Authorization implementa el modelo de autorización de la base de datos
type Authorization struct {
	UUIDAuthorization  string `json:"uuid_authorization,omitempty"`
	Register           int    `json:"register,omitempty"`
	Dateemmited        string `json:"dateemmited,omitempty"`
	Startdate          string `json:"startdate,omitempty"`
	Enddate            string `json:"enddate,omitempty"`
	Resumework         string `json:"resumework,omitempty"`
	Holidays           int    `json:"holidays,omitempty"`
	TotalDays          int    `json:"total_days,omitempty"`
	Pendingdays        int    `json:"pendingdays,omitempty"`
	Observation        string `json:"observation,omitempty"`
	Authorizationyear  string `json:"authorizationyear,omitempty"`
	Partida            string `json:"partida,omitempty"`
	Workdependency     string `json:"workdependency,omitempty"`
	WorkdependencyUUID string `json:"workdependency_uuid,omitempty"`
	User               string `json:"user,omitempty"`

	Person `json:"person,omitempty"`

	PersonnelOfficer          nulls.String `json:"personnelOfficer,omitempty"`
	PersonnelOfficerPosition  nulls.String `json:"personnelOfficerPosition,omitempty"`
	PersonnelOfficerArea      nulls.String `json:"personnelOfficerArea,omitempty"`
	ExecutiveDirector         nulls.String `json:"executiveDirector,omitempty"`
	ExecutiveDirectorPosition nulls.String `json:"executiveDirectorPosition,omitempty"`
	ExecutiveDirectorArea     nulls.String `json:"executiveDirectorArea,omitempty"`
}
