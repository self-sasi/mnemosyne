package postgresql

import (
	"context"

	"github.com/self-sasi/mnemosyne/pkg/api"
)

const postgresqlEngineName api.EngineName = "postgresql"

type PostgreSQLEngine struct {
}

func NewEngine() api.Engine {
	return &PostgreSQLEngine{}
}

func (postgreSQLEngine *PostgreSQLEngine) Name() api.EngineName {
	return postgresqlEngineName
}

func (postgreSQLEngine *PostgreSQLEngine) Backup(ctx context.Context, request api.BackupRequest) (api.BackupResponse, error) {
	return api.BackupResponse{}, nil
}

func (postgreSQLEngine *PostgreSQLEngine) Restore(ctx context.Context, request api.RestoreRequest) (api.RestoreResponse, error) {
	return api.RestoreResponse{}, nil
}
