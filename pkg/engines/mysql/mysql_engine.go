package mysql

import (
	"context"

	"github.com/self-sasi/mnemosyne/pkg/api"
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
	return api.BackupResponse{}, nil
}

func (mySqlEngine *MySQLEngine) Restore(ctx context.Context, request api.RestoreRequest) (api.RestoreResponse, error) {
	return api.RestoreResponse{}, nil
}
