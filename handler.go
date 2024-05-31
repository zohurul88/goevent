package goevent

// EventHandler is a function type that handles an event.
type EventHandler[T any] func(T)

// PriorityHandler holds an event handler and its priority.
type PriorityHandler[T any] struct {
	Handler  EventHandler[T]
	Priority int
}
