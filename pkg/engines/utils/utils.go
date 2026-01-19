package utils

import (
	"time"

	"github.com/self-sasi/mnemosyne/pkg/api"
)

func BadBackupResponse(engineName api.EngineName, startTime time.Time, finishTime time.Time) api.BackupResponse {
	return api.BackupResponse{
		Status:     api.Failure,
		Engine:     engineName,
		StartedAt:  startTime,
		FinishedAt: finishTime,
	}
}
