package api

type Registry struct {
	engines map[EngineName]Engine
}

func NewRegistry() *Registry {
	return &Registry{
		engines: make(map[EngineName]Engine),
	}
}

func (registry *Registry) AddEngine(engineName EngineName, engine Engine) {
	registry.engines[engineName] = engine
}

func (registry *Registry) GetEngine(engineName EngineName) (Engine, bool) {
	engine, ok := registry.engines[engineName]
	return engine, ok
}
