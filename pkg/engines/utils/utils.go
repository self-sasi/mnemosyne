package utils

import (
	"fmt"
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

func ArtifactName(engineName api.EngineName, dbName string, extension string) string {
	return fmt.Sprintf(`%v_%v_%v.%v`, engineName, dbName, time.Now(), extension)
}
