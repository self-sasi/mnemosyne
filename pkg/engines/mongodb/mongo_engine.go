package mongodb

import (
	"context"

	"github.com/self-sasi/mnemosyne/pkg/api"
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
	return api.BackupResponse{}, nil
}

func (mongoEngine *MongoEngine) Restore(ctx context.Context, request api.RestoreRequest) (api.RestoreResponse, error) {
	return api.RestoreResponse{}, nil
}
