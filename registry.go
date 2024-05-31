package goevent

import (
	"sync"
)

// DispatcherRegistry holds all event dispatchers and tracks async operations.
type DispatcherRegistry struct {
	dispatchers map[string]interface{}
	AsyncWG     sync.WaitGroup
	mu          sync.RWMutex
}

// NewDispatcherRegistry creates a new DispatcherRegistry.
func NewDispatcherRegistry() *DispatcherRegistry {
	return &DispatcherRegistry{
		dispatchers: make(map[string]interface{}),
	}
}

// GetDispatcher returns the dispatcher for a specific event type.
func (r *DispatcherRegistry) GetDispatcher(eventName string) interface{} {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.dispatchers[eventName]
}

// SetDispatcher sets a dispatcher for a specific event type.
func (r *DispatcherRegistry) SetDispatcher(eventName string, dispatcher interface{}) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.dispatchers[eventName]; !ok {
		r.dispatchers[eventName] = dispatcher
	}
}

// WaitForAsyncCompletion waits for all async operations to complete.
func (r *DispatcherRegistry) WaitForAsyncCompletion() {
	r.AsyncWG.Wait()
}

// Global registry instance
var globalRegistry = NewDispatcherRegistry()

// GetGlobalDispatcherRegistry returns the global dispatcher registry.
func GetGlobalDispatcherRegistry() *DispatcherRegistry {
	return globalRegistry
}
