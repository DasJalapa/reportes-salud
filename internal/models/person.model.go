package models

import "github.com/gobuffalo/nulls"

// Person implenta el modelo de la base de datos
type Person struct {
	UUID     string       `json:"uuid,omitempty"`
	Fullname string       `json:"fullname,omitempty"`
	CUI      nulls.String `json:"cui,omitempty"`
	Job      string       `json:"job,omitempty"`
	JobUUUID string       `json:"job_uuid,omitempty"`
}
