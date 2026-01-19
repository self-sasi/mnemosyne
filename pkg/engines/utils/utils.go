package utils

import (
	"os/exec"
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

func ResolveBinary(name string) (string, error) {
	return exec.LookPath(name)
}
