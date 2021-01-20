package models

type Job struct {
	UUIDJob     string `json:"uuid_job,omitempty"`
	Job         string `json:"job,omitempty"`
	Description string `json:"description,omitempty"`
}
