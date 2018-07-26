package registry

import (
	"reflect"
)

type Registry struct {
	registry map[string]reflect.Type
}

func NewRegistry() *Registry {
	registry := &Registry{
		registry: make(map[string]reflect.Type),
	}
	return registry
}

func (r *Registry) addRegistry(key string, value interface{}) {

	r.registry[key] = reflect.TypeOf(value)
}

func (r *Registry) getRegistry(key string) reflect.Type {

	return r.registry[key]
}

//----------------------------------------//
type ControllerRegistry struct {
	Registry
}

func NewControllerRegistry() *ControllerRegistry {
	registry := &ControllerRegistry{
		Registry: *NewRegistry(),
	}
	return registry
}

//-----------------------------------------//

type RegistryManager struct {
	controllerRegistry ControllerRegistry
}

func NewRegistryManager() *RegistryManager {
	manager := &RegistryManager{
		controllerRegistry: *NewControllerRegistry(),
	}
	return manager
}

/**
	receive controllers from user
 */
func (manager *RegistryManager) RegisterControllerRegistry(controllers map[string]interface{}) {
	for key, value := range controllers {
		manager.addControllerRegistry(key, value)
	}
}

func (manager *RegistryManager) addControllerRegistry(key string, value interface{}) {

	manager.controllerRegistry.addRegistry(key, value)
}

func (manager *RegistryManager) GetControllerRegistry(key string) reflect.Type {
	return manager.controllerRegistry.getRegistry(key)
}
