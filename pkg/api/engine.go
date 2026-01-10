package api

import "context"

// EngineName uniquely identifies a database engine implementation
// (for example, "mysql" or "postgres").
//
// EngineName values are used as [Registry] keys and must be stable
// across versions.
type EngineName string

// Engine represents a database-specific implementation capable of
// performing backup and restore operations.
//
// Each Engine must be self-describing via Name(), which is used to
// uniquely identify and register the engine.
type Engine interface {
	Name() EngineName
	Backup(ctx context.Context, request BackupRequest) (BackupResponse, error)
	Restore(ctx context.Context, request RestoreRequest) (RestoreResponse, error)
}
