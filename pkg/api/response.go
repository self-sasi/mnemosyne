package api

import "time"

type ResponseStatus string

type BackupResponse struct {
	Status       ResponseStatus
	Engine       EngineName
	ArtifactPath string
	StartedAt    time.Time
	FinishedAt   time.Time
}

type RestoreResponse struct {
	Status     ResponseStatus
	Engine     EngineName
	StartedAt  time.Time
	FinishedAt time.Time
}

const (
	Success ResponseStatus = "success"
	Failure ResponseStatus = "failure"
)
