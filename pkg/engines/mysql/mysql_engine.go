package mysql

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/self-sasi/mnemosyne/pkg/api"
	"github.com/self-sasi/mnemosyne/pkg/engines/utils"
)

const mySqlEngineName api.EngineName = "mysql"

type MySQLEngine struct {
}

func NewEngine() api.Engine {
	return &MySQLEngine{}
}

func (mySqlEngine *MySQLEngine) Name() api.EngineName {
	return mySqlEngineName
}

func (mySqlEngine *MySQLEngine) Backup(ctx context.Context, request api.BackupRequest) (api.BackupResponse, error) {
	startTime := time.Now()

	mysqldump, err := utils.ResolveBinary("mysqldump")
	if err != nil {
		return utils.BadBackupResponse(mySqlEngineName, startTime, time.Now()),
			errors.New(`mysqldump not found in PATH. Install MySQL client tools or configure MySQLDumpPath`)
	}

	if err := os.MkdirAll(request.TargetPath, 0o755); err != nil {
		return utils.BadBackupResponse(mySqlEngineName, startTime, time.Now()), fmt.Errorf("create target path: %w", err)
	}

	args := []string{
		"--databases", request.DBCredentials.DBName,
		"--host=" + request.DBCredentials.Host,
		"--port=" + strconv.Itoa(request.DBCredentials.Port),
		"--user=" + request.DBCredentials.UserName,
		"--password=" + request.DBCredentials.Password,
		"--result-file=" + utils.ArtifactName(mySqlEngineName, request.DBCredentials.DBName, "sql"),
	}

	cmd := exec.CommandContext(ctx, mysqldump, args...)
	if err := cmd.Run(); err != nil {
		return utils.BadBackupResponse(mySqlEngineName, startTime, time.Now()), err
	}

	return api.BackupResponse{
		Status:       api.Success,
		Engine:       mySqlEngineName,
		ArtifactPath: request.TargetPath,
		StartedAt:    startTime,
		FinishedAt:   time.Now(),
	}, nil
}

func (mySqlEngine *MySQLEngine) Restore(ctx context.Context, request api.RestoreRequest) (api.RestoreResponse, error) {
	return api.RestoreResponse{}, nil
}
