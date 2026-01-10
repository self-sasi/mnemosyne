package api

// Registry represents a store of [Engine] implementations. Every mnemo [Client]
// instance holds a registry through which it interacts with the available engines.
type Registry struct {
	engines map[EngineName]Engine
}

// NewRegistry constructs and initializes a new [Registry].
func NewRegistry() *Registry {
	return &Registry{
		engines: make(map[EngineName]Engine),
	}
}

// AddEngine registers an [Engine] implementation under the given [EngineName]
// in the registry.
func (registry *Registry) AddEngine(engineName EngineName, engine Engine) {
	registry.engines[engineName] = engine
}

// GetEngine retrieves the registered [Engine] interface value for the given
// [EngineName] in the registry.
func (registry *Registry) GetEngine(engineName EngineName) (Engine, bool) {
	engine, ok := registry.engines[engineName]
	return engine, ok
}
