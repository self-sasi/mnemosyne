package api

// EngineName uniquely identifies a database engine implementation
// (for example, "mysql" or "postgres").
//
// EngineName values are used as [Registry] keys and must be stable
// across versions.
type EngineName string

// Engine represents a database-specific implementation capable of
// performing backup and restore operations.
type Engine interface {
}
