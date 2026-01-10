package api

import (
	"errors"
)

type Client struct {
	registry *Registry
}

// NewClient constructs and initializes a new mnemo [Client].
//
// The returned client is initialized with a default, empty [Registry].
// No engines are registered by default; engine implementations must be
// explicitly registered using [Client.RegisterEngine] before they can
// be used for backup or restore operations.
func NewClient() *Client {
	return &Client{
		registry: NewRegistry(),
	}
}

// RegisterEngine registers an [Engine] implementation to the client's
// [Registry].
//
// Registered engines enable the client to perform backup and restore
// operations for the database type implemented by the engine (for example,
// a MySQL engine enables MySQL backups and restores).
//
// Each engine is identified by its [Engine.Name] and must be registered
// before it can be used.
func (client *Client) RegisterEngine(engine Engine) error {
	if engine == nil {
		return errors.New(`nil engine cannot be registed`)
	}
	client.registry.AddEngine(engine)
	return nil
}

// UseRegistry replaces the client's current [Registry] with the provided one.
//
// This allows multiple clients to share the same registry instance and,
// by extension, the same set of registered [Engine] implementations.
// Changes to the registry (such as registering additional engines) are
// immediately visible to all clients using it.
func (client *Client) UseRegistry(registry *Registry) {
	client.registry = registry
}
