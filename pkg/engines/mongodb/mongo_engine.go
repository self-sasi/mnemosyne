package mongodb

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/self-sasi/mnemosyne/pkg/api"
	"github.com/self-sasi/mnemosyne/pkg/engines/utils"
)

const mongoEngineName api.EngineName = "mongodb"

type MongoEngine struct {
}

func NewEngine() api.Engine {
	return &MongoEngine{}
}

func (mongoEngine *MongoEngine) Name() api.EngineName {
	return mongoEngineName
}

func (mongoEngine *MongoEngine) Backup(ctx context.Context, request api.BackupRequest) (api.BackupResponse, error) {
	startTime := time.Now()

	mongodump, err := mongoEngine.resolveMongoDump()
	if err != nil {
		return utils.BadBackupResponse(mongoEngineName, startTime, time.Now()), err
	}

	if err := os.MkdirAll(request.TargetPath, 0o755); err != nil {
		return utils.BadBackupResponse(mongoEngineName, startTime, time.Now()), fmt.Errorf("create target path: %w", err)
	}

	extension := ".archive"
	if request.Compress {
		extension = ".archive.gz"
	}

	archivePath := filepath.Join(
		request.TargetPath,
		fmt.Sprintf("mnemo_mongodb_%s_%s%s",
			request.DBCredentials.DBName,
			startTime.UTC().Format("20060102T150405Z"),
			extension,
		),
	)

	uri := mongoURI(request.DBCredentials)
	args := []string{
		"--uri=" + uri,
		"--archive=" + archivePath,
	}

	if request.Compress {
		args = append(args, "--gzip")
	}

	cmd := exec.CommandContext(ctx, mongodump, args...)
	if err := cmd.Run(); err != nil {
		return utils.BadBackupResponse(mongoEngineName, startTime, time.Now()), err
	}

	return api.BackupResponse{
		Status:       api.Success,
		Engine:       mongoEngineName,
		ArtifactPath: archivePath,
		StartedAt:    startTime,
		FinishedAt:   time.Now(),
	}, nil
}

func (mongoEngine *MongoEngine) Restore(ctx context.Context, request api.RestoreRequest) (api.RestoreResponse, error) {
	return api.RestoreResponse{}, nil
}

func mongoURI(credentials api.DBCredentials) string {
	hostPort := credentials.Host + ":" + strconv.Itoa(credentials.Port)
	if strings.TrimSpace(credentials.UserName) == "" {
		return fmt.Sprintf("mongodb://%s/%s", hostPort, credentials.DBName)
	}
	return fmt.Sprintf("mongodb://%s:%s@%s/%s",
		credentials.UserName, credentials.Password, hostPort, credentials.DBName,
	)
}

// why is this an instance method?
// because the tool, later, should have the functionality to accept db tool path per backup request
// different engines could talk to different versions of the same tool
func (engine *MongoEngine) resolveMongoDump() (string, error) {
	p, err := exec.LookPath("mongodump")
	if err != nil {
		return "", errors.New(`mongodump not found in PATH. Install "MongoDB Database Tools" or configure MongoDumpPath`)
	}
	return p, nil
}
